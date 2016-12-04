package repository

import (
	"goclean/entity"
	"goclean/usecase"
)

func NewUserRepo() usecase.UserRepo {
	return &userRepoImpl{}
}

// Implement the database service interface in usecase layer
type userRepoImpl struct{}

func (r *userRepoImpl) Get(id string) (*entity.User, error) {
	// TODO: call database
	user := &entity.User{
		Id: id,
	}
	return user, nil
}