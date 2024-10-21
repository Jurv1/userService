package usecase

import (
	repository2 "github.com/Jurv1/userService/internal/repository"
	"github.com/Jurv1/userService/models"
)

type UserUseCase interface {
	GetById(id int64) (*models.User, error)
	GetAll() ([]*models.User, error)
	Store(user *models.User) (int64, error)
	Update(user *models.User) (bool, error)
	Delete(id int64) (bool, error)
}

type userUseCase struct {
	uRepo      repository2.UserRepository
	uQueryRepo repository2.UserQueryRepository
}

func (u userUseCase) GetById(id int64) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) GetAll() ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) Store(user *models.User) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) Update(user *models.User) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (u userUseCase) Delete(id int64) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewUserUseCase(uRepo repository2.UserRepository, uQueryRepo repository2.UserQueryRepository) UserUseCase {
	return &userUseCase{uRepo, uQueryRepo}
}
