package repository

import (
	"goclean/entity"
	"goclean/usecase"
	"goclean-architecture/infrastructure"
)

func NewUserRepo() usecase.UserRepo {
	return &userRepoImpl{}
}

// Implement the database service interface in usecase layer
type userRepoImpl struct{}

func (r *userRepoImpl) Get(id string) (*entity.User, error) {
	// TODO: call database, this is just draft call
	row, err := infrastructure.Get("user", id)
	if err != nil{
		return nil, err
	}
	rowData, _ := row.(map[string]string)

	user := &entity.User{
		Id: rowData["id"],
	}
	return user, nil
}