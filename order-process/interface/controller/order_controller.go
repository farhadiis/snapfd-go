package controller

import (
	"farhadiis/order-process/application/interactor"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderController interface {
	GetOrders(*gin.Context)
	GetOrder(*gin.Context)
	DeleteOrder(*gin.Context)
}

type orderController struct {
	OrderInteractor interactor.OrderInteractor
}

func (t *orderController) GetOrders(c *gin.Context) {
	orders, err := t.OrderInteractor.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if orders == nil {
		c.IndentedJSON(http.StatusOK, []int{})
		return
	}
	c.IndentedJSON(http.StatusOK, orders)
}

func (t *orderController) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := t.OrderInteractor.Get(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if order != nil {
		c.IndentedJSON(http.StatusOK, order)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"_id": id, "founded": false})
	}
}

func (t *orderController) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	ok, err := t.OrderInteractor.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if ok {
		c.IndentedJSON(http.StatusOK, gin.H{"_id": id, "deleted": ok})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"_id": id, "deleted": ok})
	}
}

func NewOrderController(ti interactor.OrderInteractor) OrderController {
	return &orderController{ti}
}
