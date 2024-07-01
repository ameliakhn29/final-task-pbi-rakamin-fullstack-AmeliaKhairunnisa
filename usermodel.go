package models

import (
	"database/sql"

	"github.com/Asus/final-task/app"
	"github.com/Asus/final-task/database"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := database.DBConn()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Where(user *app.User, fieldName, fieldValue string) error {

	row, err := u.db.Query("select id, username, email, password from users where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	}

	return nil

}

func (u UserModel) Create(user app.User) (int64, error) {

	result, err := u.db.Exec("insert into users (username, email, password) values(?,?,?)",
		user.Username, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil

}
