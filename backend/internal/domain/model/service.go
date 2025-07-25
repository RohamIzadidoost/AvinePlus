package model

type Service struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"unique"`
	Description string
	ShipmentFee float64
}

const (
	ServiceOnTableGlass = "ontable_glass"
)
