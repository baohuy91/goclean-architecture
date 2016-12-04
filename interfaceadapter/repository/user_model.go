package repository

import (
	"time"
)

// Data model that will be stored to database
// It could have some extra fields like CreatedTime, UpdatedTime...
type User struct {
	Id string
	Name string
	Email string
	CreatedTime time.Time
	UpdatedTime time.Time
}
