package gormrepo

import (
	"glasscutting/internal/domain/model"
	"glasscutting/internal/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(u *model.User) error {
	return r.db.Create(u).Error
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var u model.User
	err := r.db.Where("email = ?", email).First(&u).Error
	return &u, err
}

func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var u model.User
	err := r.db.First(&u, id).Error
	return &u, err
}
