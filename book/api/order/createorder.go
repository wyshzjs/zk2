package order

import (
	"book/db/mysql"
	"book/model"
	"book/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var userId = 1

type JyInfo struct {
	BookId int    `json:"book_id"`
	JyTime string `json:"jy_time"`
}

func CreateOrder(c *gin.Context) {
	var jyInfo JyInfo
	err := c.ShouldBind(&jyInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}

	if jyInfo.JyTime == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "借阅时间不能为空",
		})
		return
	}

	book, err := mysql.GetBookBYId(jyInfo.BookId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}

	if book.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "该书不存在",
		})
		return
	}
	amount := book.Price
	sn := uuid.NewString()

	order := model.Order{
		UserId: userId,
		Sn:     sn,
		Amount: amount,
	}

	err = mysql.CreateOrder(order, book, jyInfo.JyTime)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}
	url, err := pkg.Pay(sn, sn, fmt.Sprintf("%.2f", amount))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    url,
	})
}
