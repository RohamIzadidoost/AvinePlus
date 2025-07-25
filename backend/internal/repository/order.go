package repository

import "glasscutting/internal/domain/model"

// OrderRepository defines interface for orders

type OrderRepository interface {
	Create(order *model.Order) error
	FindByID(id uint) (*model.Order, error)
	ListAll() ([]model.Order, error)
	Update(order *model.Order) error
	ListByAssigned(userID uint) ([]model.Order, error)
	ListByUser(userID uint) ([]model.Order, error)
}
