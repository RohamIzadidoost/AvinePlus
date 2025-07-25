package model

import "gorm.io/gorm"

const (
	StatusRegistered = "registered"
	StatusVerified   = "verified"
	StatusAssigned   = "assigned"
	StatusApproved   = "approved"
	StatusDone       = "done"
)

type ShapeType string

const (
	ShapeRectangular ShapeType = "rectangular"
	ShapeCircular    ShapeType = "circular"
	ShapePolygon     ShapeType = "polygon"
)

type Rectangular struct {
	Width  float64
	Height float64
}

type Circular struct {
	Diameter float64
}

type Polygon struct {
	Description string
}

type OnTableGlassDetails struct {
	Thickness   float64
	Color       string
	RoundTrim   bool
	Pocket      string
	ShapeType   ShapeType
	Rectangular *Rectangular `gorm:"embedded;embeddedPrefix:rect_"`
	Circular    *Circular    `gorm:"embedded;embeddedPrefix:circ_"`
	Polygon     *Polygon     `gorm:"embedded;embeddedPrefix:poly_"`
}

type Order struct {
	gorm.Model
	UserID     uint
	Service    string
	Details    OnTableGlassDetails `gorm:"embedded;embeddedPrefix:details_"`
	Status     string
	AssignedTo *uint
}
