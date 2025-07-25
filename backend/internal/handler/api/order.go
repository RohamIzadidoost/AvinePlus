package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"glasscutting/internal/domain/model"
	"glasscutting/internal/service"
)

type OrderHandler struct {
	svc *service.OrderService
}

func NewOrderHandler(s *service.OrderService) *OrderHandler {
	return &OrderHandler{svc: s}
}

// create order for OnTableGlass
func (h *OrderHandler) Create(c *gin.Context) {
	var req struct {
		Thickness   float64         `json:"thickness"`
		Color       string          `json:"color"`
		RoundTrim   bool            `json:"round_trim"`
		Pocket      string          `json:"pocket"`
		ShapeType   model.ShapeType `json:"shape_type"`
		Width       float64         `json:"width"`
		Height      float64         `json:"height"`
		Diameter    float64         `json:"diameter"`
		Description string          `json:"description"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uidVal, _ := c.Get("userID")
	uid := uidVal.(uint)
	details := model.OnTableGlassDetails{
		Thickness: req.Thickness,
		Color:     req.Color,
		RoundTrim: req.RoundTrim,
		Pocket:    req.Pocket,
		ShapeType: req.ShapeType,
	}
	switch req.ShapeType {
	case model.ShapeRectangular:
		details.Rectangular = &model.Rectangular{Width: req.Width, Height: req.Height}
	case model.ShapeCircular:
		details.Circular = &model.Circular{Diameter: req.Diameter}
	case model.ShapePolygon:
		details.Polygon = &model.Polygon{Description: req.Description}
	}
	order := &model.Order{UserID: uid, Service: model.ServiceOnTableGlass, Details: details}
	if err := h.svc.Create(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) ListAll(c *gin.Context) {
	orders, _ := h.svc.ListAll()
	c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.svc.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.svc.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.svc.Update(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}
