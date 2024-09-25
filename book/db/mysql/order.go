package mysql

import (
	"book/model"
	"gorm.io/gorm"
	"time"
)

func CreateOrder(order model.Order, book model.Book, jyTime string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&order).Error
		if err != nil {
			return err
		}
		parse, err := time.Parse(time.DateTime, jyTime)
		if err != nil {
			return err
		}
		orderBook := model.OrderBook{
			OrderId:   int(order.ID),
			BookId:    int(book.ID),
			JyTime:    parse,
			BookTitle: book.BookTitle,
		}

		return tx.Create(&orderBook).Error
	})
}

func GetOrderBySn(sn string) (order model.Order, err error) {
	err = db.Where("sn = ? and status = 1", sn).Limit(1).Find(&order).Error
	return
}

func GetOrderBookByOrderId(orderId uint) (orderBook model.OrderBook, err error) {
	err = db.Where("order_id = ?", orderId).Limit(1).Find(&orderBook).Error
	return
}
