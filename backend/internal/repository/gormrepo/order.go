package gormrepo

import (
	"glasscutting/internal/domain/model"
	"glasscutting/internal/repository"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) FindByID(id uint) (*model.Order, error) {
	var order model.Order
	err := r.db.First(&order, id).Error
	return &order, err
}

func (r *orderRepository) ListAll() ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *orderRepository) Update(order *model.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) ListByAssigned(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Where("assigned_to = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *orderRepository) ListByUser(userID uint) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}
