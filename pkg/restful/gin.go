package restful

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	errService "github.com/nguyenhoai890/wager_service/pkg/error"
	"github.com/nguyenhoai890/wager_service/pkg/service"
	"net/http"
)

type Handler struct {
	engine *gin.Engine
	service service.IWager
}

func InitGin(service service.IWager) http.Handler {
	h := &Handler{engine: gin.Default(), service: service}
	h.SetupRouter()
	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.engine.ServeHTTP(w, req)
}

func (h *Handler) SetupRouter() {
	h.engine.POST("/wagers", h.CreateWager)
	h.engine.POST("/buy/:wager_id", h.BuyWager)
	h.engine.GET("/wagers", h.ListWager)
}

func (h *Handler) BuyWager(c *gin.Context) {
	uriData := struct{
		WagerId int64 `uri:"wager_id" binding:"required,gt=0"`
	}{}
	err := c.ShouldBindUri(&uriData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errService.ErrInvalidData.Error()})
		return
	}
	var request service.BuyWagerRequest
	if err = c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errService.ErrInvalidData.Error()})
		return
	}
	transaction, err := h.service.Buy(uriData.WagerId, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if transaction == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("there is some thing wrong with wager Id %v and buying price %v",uriData.WagerId, request.BuyingPrice)})
		return
	}
	c.JSON(http.StatusCreated, transaction)
	return
}

func (h *Handler) CreateWager(c *gin.Context) {
	var request service.CreateWagerRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errService.ErrInvalidData.Error()})
		return
	}

	wager, err := h.service.Create(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, wager)
	return
}

func (h *Handler) ListWager(c *gin.Context) {
	var request service.ListWagerRequestQuery
	err := c.ShouldBindWith(&request, binding.Query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errService.ErrInvalidData.Error()})
		return
	}
	wagers, err := h.service.List(request.Page, request.Limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wagers)
}

func (h *Handler) Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}