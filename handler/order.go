package handler

import (
	"assignment2/dto"
	"assignment2/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	OrderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) orderHandler {
	return orderHandler{OrderService: orderService}
}

// CreateOrder godoc
// @Tags Orders
// @Description Create an Order with Items
// @ID create-order
// @Accept json
// @Produce json
// @Param Order body dto.NewOrderRequest true "Create order request body"
// @Success 201 {object} models.Order
// @Router /orders [post]
func (oh *orderHandler) CreateOrder(ctx *gin.Context) {
	var newOrderRequest dto.NewOrderRequest

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid json request",
		})
		return
	}
	err := oh.OrderService.CreateOrder(&newOrderRequest)
	dto.WriteJsonResponse(ctx, err)
}

// GetOrders godoc
// @Tags Orders
// @Description Get all order data
// @ID get-all-order
// @Accept json
// @Produce json
// @Success 200 {object} models.Order
// @Router /orders [get]
func (oh *orderHandler) GetOrders(ctx *gin.Context) {
	err := oh.OrderService.GetOrders()
	dto.WriteJsonResponse(ctx, err)
}

// UpdateOrder godoc
// @Tags Orders
// @Description Update order by given orderId
// @ID update-order
// @Accept json
// @Produce json
// @Params orderId path int true "Order Id want to be updated"
// @Params Order body dto.NewOrderRequest true "Update order request body"
// @Success 200 {object} models.Order
// @Router /orders/{orderId} [put]
func (oh *orderHandler) UpdateOrder(ctx *gin.Context) {
	var newOrderRequest dto.NewOrderRequest
	var orderId, _ = strconv.Atoi(ctx.Param("orderId"))

	if err := ctx.ShouldBindJSON(&newOrderRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "invalid json request",
		})
		return
	}

	err := oh.OrderService.UpdateOrder(orderId, &newOrderRequest)
	dto.WriteJsonResponse(ctx, err)
}

// DeleteOrder godoc
// @Tags Orders
// @Description Delete order by given orderId
// @ID delete-order
// @Accept json
// @Produce json
// @Params orderId path int true "Order Id want to be deleted"
// @Success 204 "No Content"
// @Router /orders/{orderId} [delete]
func (oh *orderHandler) DeleteOrder(ctx *gin.Context) {
	var orderId, _ = strconv.Atoi(ctx.Param("orderId"))

	err := oh.OrderService.DeleteOrder(orderId)
	dto.WriteJsonResponse(ctx, err)
}