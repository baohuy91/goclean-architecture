package usecase

import "goclean/entity"

// Interface for the Database service
// This is what outside layer should do to full fill the business logic
type UserRepo interface {
	Get(id string) (*entity.User, error)
}
