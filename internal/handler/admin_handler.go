package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thesayedirfan/ecommerce/internal/errors"
	"github.com/thesayedirfan/ecommerce/internal/service"
	"github.com/thesayedirfan/ecommerce/internal/types"
)

type AdminHandler struct {
	discountService *service.AdminService
}

func NewAdminHandler(discountService *service.AdminService) *AdminHandler {
	return &AdminHandler{discountService: discountService}
}

func (h *AdminHandler) GenerateDiscountCode(c *gin.Context) {
	var req types.GenerateDiscountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrUserIDCannotBeEmpty.Error()})
		return
	}

	code, err := h.discountService.GenerateDiscountCode(req.UserID, 2,10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"discount_code": code})
}

func (h *AdminHandler) GetAdminStats(c *gin.Context) {
	stats := h.discountService.GetAdminStats()
	c.JSON(http.StatusOK, stats)
}
