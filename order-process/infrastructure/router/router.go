package router

import (
	"farhadiis/order-process/interface/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter(app *gin.Engine, c controller.AppController) *gin.Engine {

	app.GET("/api/v1", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{"service": "running"})
	})
	//app.GET("/api/v1/order", c.GetOrders)
	//app.GET("/api/v1/order/:id", c.GetOrder)
	//app.DELETE("/api/v1/order/:id", c.DeleteOrder)

	return app
}
