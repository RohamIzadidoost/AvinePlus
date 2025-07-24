package service

import (
	"glasscutting/internal/domain/model"
	"glasscutting/internal/repository"
)

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(r repository.OrderRepository) *OrderService {
	return &OrderService{repo: r}
}

func (s *OrderService) Create(order *model.Order) error {
	order.Status = model.StatusRegistered
	return s.repo.Create(order)
}

func (s *OrderService) ListAll() ([]model.Order, error) {
	return s.repo.ListAll()
}

func (s *OrderService) Get(id uint) (*model.Order, error) {
	return s.repo.FindByID(id)
}

func (s *OrderService) Update(order *model.Order) error {
	return s.repo.Update(order)
}

func (s *OrderService) ListByAssigned(userID uint) ([]model.Order, error) {
	return s.repo.ListByAssigned(userID)
}

func (s *OrderService) ListByUser(userID uint) ([]model.Order, error) {
	return s.repo.ListByUser(userID)
}
