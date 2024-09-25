package order

import (
	"book/api/order"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine) {
	r := g.Group("order")

	r.POST("create", order.CreateOrder)
	r.POST("get", order.GetOrder)
}
