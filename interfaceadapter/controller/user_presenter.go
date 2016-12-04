package controller

import "goclean/entity"

// There API response data (i.e. json) presentations
type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Present user entity to json response
func NewUser(ue *entity.User) *User {
	return &User{
		Id: ue.Id,
		Name: ue.Name,
		Email: ue.Email,
	}
}