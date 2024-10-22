package repository

import (
	"github.com/Jurv1/userService/internal/models"
	"github.com/jackc/pgx"
)

type UserQueryRepository interface {
	GetAll() []*models.User
	GetById(id int64) (*models.User, error)
}

type UserQRepo struct {
	db *pgx.Conn
}

func (u UserQRepo) GetAll() []*models.User {
	//TODO implement me
	panic("implement me")
}

func (u UserQRepo) GetById(id int64) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserQRepo(db *pgx.Conn) UserQRepo {
	return UserQRepo{db}
}
