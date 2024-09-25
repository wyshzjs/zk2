package routers

import (
	"book/routers/book"
	"book/routers/order"
	"github.com/gin-gonic/gin"
)

func Routers(g *gin.Engine) {
	book.Router(g)
	order.Router(g)
}
