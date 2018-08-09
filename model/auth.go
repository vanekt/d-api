package model

import (
	"github.com/jmoiron/sqlx"
	"github.com/kataras/iris/core/errors"
	"github.com/op/go-logging"
)

type UserModel struct {
	db     *sqlx.DB
	logger *logging.Logger
}

func NewUserModel(db *sqlx.DB, logger *logging.Logger) *UserModel {
	return &UserModel{db, logger}
}

func (m *UserModel) GetUserByLogin() (err error) {
	err = errors.New("Not implemented yet")
	return
}
