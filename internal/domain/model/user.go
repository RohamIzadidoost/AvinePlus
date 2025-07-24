package model

import "gorm.io/gorm"

// Role type constants
const (
	RoleCustomer = "customer"
	RoleAdmin    = "admin"
	RoleEmployee = "employee"
)

// User represents application user
// gorm.Model includes ID, CreatedAt, UpdatedAt, DeletedAt

// User entity

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Role     string
}
