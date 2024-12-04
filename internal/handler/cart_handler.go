package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thesayedirfan/ecommerce/internal/errors"
	"github.com/thesayedirfan/ecommerce/internal/service"
	"github.com/thesayedirfan/ecommerce/internal/types"
)

type CartHandler struct {
	cartService     *service.CartService
	discountService *service.AdminService
}

func NewCartHandler(cartService *service.CartService, discountService *service.AdminService) *CartHandler {
	return &CartHandler{cartService: cartService,
		discountService: discountService,
	}
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	var req types.AddToCartRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.ProductID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrItemProductIDCannotBeEmpty.Error()})
		return
	}

	if req.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrUserIDCannotBeEmpty.Error()})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrItemNameCannotBeEmpty.Error()})
		return
	}

	if req.Quantity == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrItemQuantityCannotBeZero.Error()})
		return
	}
	if req.Price == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrItemPriceCannotBeZero.Error()})
		return
	}

	err := h.cartService.AddToCart(req.UserID, req.ProductID, req.Name, req.Price, req.Quantity)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	code, _ := h.discountService.GenerateDiscountCode(req.UserID, 2, 10)

	if code != "" {
		c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully", "discount_code": code})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart successfully"})
}

func (h *CartHandler) Checkout(c *gin.Context) {
	var req types.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.ErrUserIDCannotBeEmpty.Error()})
		return
	}

	order, err := h.cartService.Checkout(req.UserID, req.DiscountCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}
