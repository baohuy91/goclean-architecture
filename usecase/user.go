package usecase

import "goclean/entity"

type UserUseCase interface {
	GetUser(id string) (*entity.User, error)
}

func NewUserUseCase(userRepo UserRepo) UserUseCase {
	return &userUseCaseImpl{
		userRepo:userRepo,
	}
}

type userUseCaseImpl struct {
	userRepo UserRepo
}

// Business logic for getting user will be implemented here
func (u *userUseCaseImpl) GetUser(id string) (*entity.User, error) {
	// Get user from repository & handle error if necessary
	user, err := u.userRepo.Get(id)
	if err!= nil{
		return nil, err
	}

	return user, err
}