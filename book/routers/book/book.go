package book

import (
	"book/api/book"
	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine) {
	r := g.Group("book")

	r.POST("list", book.GetBookList)
	r.POST("detail", book.GetBookDetail)
}
