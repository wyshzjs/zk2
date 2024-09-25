package order

import (
	"book/db/mysql"
	"book/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetOrderInfo struct {
	Sn string `json:"sn"`
}

type res struct {
	Order     model.Order     `json:"order"`
	OrderBook model.OrderBook `json:"order_book"`
}

func GetOrder(c *gin.Context) {
	var getOrderInfo GetOrderInfo
	err := c.ShouldBind(&getOrderInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}

	order, err := mysql.GetOrderBySn(getOrderInfo.Sn)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}

	orderBook, err := mysql.GetOrderBookByOrderId(order.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "感谢借阅",
		"data": res{
			Order:     order,
			OrderBook: orderBook,
		},
	})

}
