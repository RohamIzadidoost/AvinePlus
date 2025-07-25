package repository

import "glasscutting/internal/domain/model"

type UserRepository interface {
	Create(u *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
}
