package repository

import (
	"github.com/Jurv1/userService/internal/models"
	"github.com/jackc/pgx"
)

type UserRepository interface {
	Store(user *models.User) (int64, error)
	Update(user *models.User) (bool, error)
	Delete(id int64) (bool, error)
}

type UserRepo struct {
	db *pgx.Conn
}

func (u UserRepo) Store(user *models.User) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) Update(user *models.User) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) Delete(id int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserRepo(db *pgx.Conn) UserRepo {
	return UserRepo{db}
}
