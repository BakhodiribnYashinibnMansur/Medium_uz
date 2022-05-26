package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/BakhodiribnYashinibnMansur/Medium_uz/model"
	"github.com/BakhodiribnYashinibnMansur/Medium_uz/util/logrus"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserDB struct {
	db    *sqlx.DB
	redis *redis.Client
}

func NewUserDB(db *sqlx.DB, redis *redis.Client) *UserDB {
	return &UserDB{db: db, redis: redis}
}

func (repo *UserDB) GetUserData(id string, logrus *logrus.Logger) (model.UserFull, error) {
	var user model.UserFull
	query := fmt.Sprintf("SELECT  	id,	email,	firstname,secondname,nickname,		city,	is_verified,	bio,interests,account_image_path,	phone,	rating,	post_views_count,	follower_count, following_count,like_count,is_super_user	FROM %s WHERE id=$1 ", usersTable)
	err := repo.db.Get(&user, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("ID not found")
		}
		logrus.Errorf("ERROR: don't get users %s", err)
		return user, err
	}
	logrus.Info("DONE:get user data from psql")
	return user, nil
}

func (repo *UserDB) UpdateUserVerified(id string, logrus *logrus.Logger) (effectedRowsNum int64, err error) {
	query := fmt.Sprintf("	UPDATE %s SET is_verified = true,verification_date=NOW(),updated_at=NOW() WHERE id = $1  RETURNING id ", usersTable)
	rows, err := repo.db.Exec(query, id)

	if err != nil {
		logrus.Errorf("ERROR: Update verificationCode : %v", err)
		return 0, err
	}
	effectedRowsNum, err = rows.RowsAffected()
	if err != nil {
		logrus.Errorf("ERROR: Update verificationCode effectedRowsNum : %v", err)
		return 0, err
	}
	logrus.Info("DONE:Update verify email")
	return effectedRowsNum, nil
}

func (repo *UserDB) UpdateAccountImage(id int, filePath string, logrus *logrus.Logger) (int64, error) {
	query := fmt.Sprintf("	UPDATE %s SET account_image_path=$1,updated_at=NOW()	WHERE id = $1  RETURNING id ", usersTable)
	rows, err := repo.db.Exec(query, filePath, id)

	if err != nil {
		logrus.Errorf("ERROR: Update Update account image : %v", err)
		return 0, err
	}
	effectedRowsNum, err := rows.RowsAffected()
	if err != nil {
		logrus.Errorf("ERROR: Update Update account image effectedRowsNum : %v", err)
		return 0, err
	}
	logrus.Info("DONE:Update account image")
	return effectedRowsNum, nil
}

func (repo *UserDB) UpdateAccount(id int, user model.UpdateUser, logrus *logrus.Logger) (int64, error) {
	query := fmt.Sprintf("	UPDATE %s SET 	firstname = COALESCE($1,firstname), 	secondname  = COALESCE($2,secondname), 	email = COALESCE($3,email), 	nickname = COALESCE( $4,nickname), 	password_hash = COALESCE($5,password_hash),  	interests = COALESCE($6, interests), 	bio = COALESCE($7,bio), 	city = COALESCE($8,city), 	phone = COALESCE($9,phone),  	updated_at=NOW()		WHERE id = $10 	 RETURNING id ", usersTable)
	rows, err := repo.db.Exec(query, user.FirstName, user.SecondName, user.Email, user.NickName, user.Password, pq.Array(user.Interests), user.Bio, user.City, user.Phone, id)

	if err != nil {
		logrus.Errorf("ERROR: Update verificationCode : %v", err)
		return 0, err
	}
	effectedRowsNum, err := rows.RowsAffected()
	if err != nil {
		logrus.Errorf("ERROR: Update verificationCode effectedRowsNum : %v", err)
		return 0, err
	}
	logrus.Info("DONE:Update verify email")
	return effectedRowsNum, nil
}

func (repo *UserDB) FollowingAccount(userID, followingID int, logrus *logrus.Logger) (int64, error) {

	followingQuery := fmt.Sprintln("SELECT toggle_following_user($1,$2)")

	row, err := repo.db.Exec(followingQuery, userID, followingID)

	if err != nil {
		logrus.Info("DONE: ERROR  Following Data PSQL %s ", err)
		return -1, err
	}

	numRowEffected, err := row.RowsAffected()
	if err != nil {
		logrus.Info("DONE: ERROR  Following Data PSQL %s ", err)
		return -1, err
	}

	logrus.Info("DONE: INSERTED  Following Data PSQL")
	return numRowEffected, nil

}

func (repo *UserDB) FollowerAccount(userID, followerID int, logrus *logrus.Logger) (int64, error) {

	followerQuery := fmt.Sprintln("SELECT toggle_follower_user($1,$2)")

	row, err := repo.db.Exec(followerQuery, userID, followerID)

	if err != nil {
		logrus.Info("DONE: ERROR  Following Data PSQL %s ", err)
		return -1, err
	}

	numRowEffected, err := row.RowsAffected()
	if err != nil {
		logrus.Info("DONE: ERROR  Following Data PSQL %s ", err)
		return -1, err
	}

	logrus.Info("DONE: INSERTED  Following Data PSQL")
	return numRowEffected, nil
}

func (repo *UserDB) GetFollowers(userID int, logrus *logrus.Logger) (user []model.UserFull, err error) {

	query := fmt.Sprintf("SELECT 	u.id,	u.email,	u.firstname,u.secondname,u.nickname,		city,	u.is_verified,	u.bio,u.interests,u.account_image_path,	u.phone,	u.rating,	u.post_views_count,	u.follower_count, u.following_count,u.like_count,u.is_super_user FROM %s u INNER JOIN %s uf on u.id = uf.follower_id WHERE uf.account_id = $1 AND uf.deleted_at IS NULL",
		usersTable, userFollowerTable)
	err = repo.db.Select(&user, query, userID)

	return user, err
}
func (repo *UserDB) GetFollowings(userID int, logrus *logrus.Logger) (user []model.UserFull, err error) {
	query := fmt.Sprintf("SELECT 	u.id,	u.email,	u.firstname,u.secondname,u.nickname,		city,	u.is_verified,	u.bio,u.interests,u.account_image_path,	u.phone,	u.rating,	u.post_views_count,	u.follower_count, u.following_count,u.like_count,u.is_super_user FROM %s u INNER JOIN %s uf on u.id = uf.following_id WHERE uf.account_id = $1 AND uf.deleted_at IS NULL",
		usersTable, userFollowingTable)
	err = repo.db.Select(&user, query, userID)

	return user, err
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// REDIS

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (repo *UserDB) CheckCode(email, code string, logrus *logrus.Logger) error {
	saveCode, err := repo.redis.Get(email).Result()
	if err != nil {
		logrus.Errorf("ERROR:don't save code %s", err)
		return err
	}
	if saveCode != code {
		return errors.New("code not found ")
	}
	logrus.Info("DONE: verify code")
	return nil
}
