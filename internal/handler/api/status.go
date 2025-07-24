package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"glasscutting/internal/domain/model"
	"glasscutting/internal/service"
)

type StatusHandler struct {
	svc *service.OrderService
}

func NewStatusHandler(s *service.OrderService) *StatusHandler {
	return &StatusHandler{svc: s}
}

// Admin verifies order
func (h *StatusHandler) Verify(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.svc.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	order.Status = model.StatusVerified
	h.svc.Update(order)
	c.JSON(http.StatusOK, order)
}

// Admin assign to employee
func (h *StatusHandler) Assign(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req struct {
		EmployeeID uint `json:"employee_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, err := h.svc.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	order.Status = model.StatusAssigned
	order.AssignedTo = &req.EmployeeID
	h.svc.Update(order)
	c.JSON(http.StatusOK, order)
}

// Employee approves completion
func (h *StatusHandler) EmployeeApprove(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.svc.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	order.Status = model.StatusApproved
	h.svc.Update(order)
	c.JSON(http.StatusOK, order)
}

// User approve completion and mark done
func (h *StatusHandler) UserApprove(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	order, err := h.svc.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	order.Status = model.StatusDone
	h.svc.Update(order)
	c.JSON(http.StatusOK, order)
}
