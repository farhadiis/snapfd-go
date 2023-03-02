package router

import (
	"farhadiis/order/interface/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(app *gin.Engine, c controller.AppController) *gin.Engine {

	//app.GET("/api/v1/order", c.GetOrders)
	//app.GET("/api/v1/order/:id", c.GetOrder)
	app.POST("/api/v1/order", c.CreateOrder)
	//app.DELETE("/api/v1/order/:id", c.DeleteOrder)

	return app
}
