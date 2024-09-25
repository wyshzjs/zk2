package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserId  int        `gorm:"type:int(11);comment:借阅者id"`
	Sn      string     `gorm:"type:char(42);comment:订单编号"`
	Amount  float64    `gorm:"type:decimal(11,2);comment:总价格"`
	PayType int8       `gorm:"type:tinyint(1);comment:0支付宝 1微信"`
	PayTime *time.Time `gorm:"type:datetime;comment:支付时间"`
	Status  int8       `gorm:"type:tinyint(1);default:0;comment:0待支付 1已支付"`
}

type OrderBook struct {
	gorm.Model
	OrderId   int       `gorm:"type:int(11);comment:订单id"`
	BookId    int       `gorm:"type:int(11);comment:书本id"`
	JyTime    time.Time `gorm:"type:datetime;comment:借书时间"`
	BookTitle string    `gorm:"type:varchar(255);comment:书本名称"`
}
