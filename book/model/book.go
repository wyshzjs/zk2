package model

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	gorm.Model
	BookTitle string    `json:"book_title" gorm:"type:varchar(255);comment:书本名称"`
	Nickname  string    `json:"nickname"   gorm:"type:varchar(255);comment:作者名称"`
	Isbn      string    `json:"isbn"       gorm:"type:varchar(255);comment:ISBN码"`
	Price     float64   `json:"price"      gorm:"type:decimal(11,2);comment:价格"`
	CbTime    time.Time `json:"cb_time"    gorm:"type:datetime;comment:出版日期"`
}
