package service

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/configs"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/package/repository"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/convert"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (service *UserService) GetUserData(id string, logrus *logrus.Logger) (user model.UserFull, err error) {
	user, err = service.repo.GetUserData(id, logrus)
	if err != nil {
		logrus.Error("ERROR: get user Data failed: %v", err)
		return user, err
	}
	return user, nil
}

func (service *UserService) VerifyCode(id, email, code string, logrus *logrus.Logger) (int64, error) {
	err := service.repo.CheckCode(email, code, logrus)
	if err != nil {
		logrus.Errorf("ERROR: check error code : %v", err)
		return 0, err
	}
	effectedRowsNum, err := service.repo.UpdateUserVerified(id, logrus)
	if err != nil {
		logrus.Errorf("ERROR: check error code update : %v", err)
		return 0, err
	}
	return effectedRowsNum, nil
}

func (service *UserService) UploadImage(file multipart.File, header *multipart.FileHeader, logrus *logrus.Logger) (string, error) {

	filename := header.Filename
	folderPath := "public/"
	err := os.MkdirAll(folderPath, 0777)
	if err != nil {
		logrus.Errorf("ERROR: Failed to create folder %s: %v", folderPath, err)
		return "", err
	}
	err = os.Chmod(folderPath, 0777)
	if err != nil {
		logrus.Errorf("ERROR: Failed Access Permission denied %s", err)
		return "", err
	}
	filePath := folderPath + filename
	out, err := os.Create(filePath)
	if err != nil {
		logrus.Errorf("ERROR: Failed CreateFile: %v", err)
		return "", err
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		logrus.Errorf("ERROR: Failed copy %s", err)
		return "", err
	}
	configs, err := configs.InitConfig()
	logrus.Infof("configs %v", configs)
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	imageURL := configs.ServiceHost + "/" + filePath
	return imageURL, nil
}

func (service *UserService) UpdateAccountImage(id int, filePath string, logrus *logrus.Logger) (int64, error) {
	return service.repo.UpdateAccountImage(id, filePath, logrus)
}

func (service *UserService) UpdateAccount(id int, user model.User, logrus *logrus.Logger) (int64, error) {
	var updateUser model.UpdateUser
	updateUser.Bio = convert.EmptyStringToNull(user.Bio)
	updateUser.Email = convert.EmptyStringToNull(user.Email)
	user.Password = generatePasswordHash(user.Password)
	updateUser.Password = convert.EmptyStringToNull(user.Password)
	updateUser.City = convert.EmptyStringToNull(user.City)
	updateUser.FirstName = convert.EmptyStringToNull(user.FirstName)
	updateUser.SecondName = convert.EmptyStringToNull(user.SecondName)
	updateUser.NickName = convert.EmptyStringToNull(user.NickName)
	updateUser.Phone = convert.EmptyStringToNull(user.Phone)

	updateUser.Interests = convert.EmptyArrayStringToNullArray(user.Interests)

	logrus.Info("successfully password_hash")
	return service.repo.UpdateAccount(id, updateUser, logrus)
}

func (service *UserService) FollowingAccount(userID, followingID int, logrus *logrus.Logger) (int64, error) {
	return service.repo.FollowingAccount(userID, followingID, logrus)
}

func (service *UserService) FollowerAccount(userID, followerID int, logrus *logrus.Logger) (int64, error) {
	return service.repo.FollowerAccount(userID, followerID, logrus)
}

func (service *UserService) GetFollowers(userID int, logrus *logrus.Logger) (user []model.UserFull, err error) {
	return service.repo.GetFollowers(userID, logrus)
}
func (service *UserService) GetFollowings(userID int, logrus *logrus.Logger) (user []model.UserFull, err error) {
	return service.repo.GetFollowings(userID, logrus)
}
