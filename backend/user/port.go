package user

import (
	"ecommerce/domain"
	userHandler "ecommerce/rest/handlers/user"
)

type Service interface {
	userHandler.Service // embading er mane hocche signature import kora
}

type UserRepo interface {
	Create(p domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
	// List(p User) ([]*User, error)
	// Delete(userID int) error
	// Update(p User) (*User, error)
}
