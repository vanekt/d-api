package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/op/go-logging"
	"vanekt/dental-api/entity"
)

type UserModel struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func NewUserModel(db *sqlx.DB, logger *logging.Logger) *UserModel {
	return &UserModel{db, logger}
}

func (m *UserModel) GetUserByLogin(login string) (user entity.User, err error) {
	err = m.db.Get(&user, "select * from users where login = ?", login)
	return
}
