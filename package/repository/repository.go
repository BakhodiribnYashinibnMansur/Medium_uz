package repository

import (
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User, logrus *logrus.Logger) (int, error)
	GetUserID(email string, logrus *logrus.Logger) (int, error)
	CheckDataExistsEmailPassword(email, password string, logrus *logrus.Logger) (int, error)
	CheckDataExistsEmailNickName(email, nickname string, logrus *logrus.Logger) (int, int, error)
	SaveVerificationCode(email, code string, logrus *logrus.Logger) error
}

type User interface {
	GetUserData(id string, logrus *logrus.Logger) (model.UserFull, error)
	CheckCode(email, code string, logrus *logrus.Logger) error
	UpdateUserVerified(id string, logrus *logrus.Logger) (int64, error)
	UpdateAccountImage(id int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdateAccount(id int, user model.UpdateUser, logrus *logrus.Logger) (int64, error)
}

type Post interface {
	CreatePost(post model.Post, logrus *logrus.Logger) (int, error)
	CreatePostUser(userId, postId int, logrus *logrus.Logger) (int, error)
	GetPostById(id int, logrus *logrus.Logger) (model.PostFull, error)
	UpdatePostImage(userID, postID int, filePath string, logrus *logrus.Logger) (int64, error)
	UpdatePost(userID, postID int, post model.UpdatePost, logrus *logrus.Logger) (int64, error)
	DeletePost(userID, postID int, logrus *logrus.Logger) (int64, int64, error)
}
type Search interface {
	SearchPost(search string, logrus *logrus.Logger) ([]model.PostFull, error)
	SearchUser(search string, logrus *logrus.Logger) ([]model.UserFull, error)
}

type Repository struct {
	Authorization
	User
	Post
	Search
}

func NewRepository(db *sqlx.DB, redis *redis.Client) *Repository {
	return &Repository{Authorization: NewAuthDB(db, redis), User: NewUserDB(db, redis), Post: NewPostDB(db, redis), Search: NewSearchDB(db)}
}
