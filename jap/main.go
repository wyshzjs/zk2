package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type Book struct {
	gorm.Model
	BookTitle string    `json:"book_title" gorm:"type:varchar(255);comment:书本名称"`
	Nickname  string    `json:"nickname" gorm:"type:varchar(255);comment:作者名称"`
	Isbn      string    `json:"isbn" gorm:"type:varchar(255);comment:ISBN码"`
	Price     float64   `json:"price" gorm:"type:decimal(11,2);comment:价格"`
	CbTime    time.Time `json:"cb_time" gorm:"type:datetime;comment:出版日期"`
}

func main() {
	db, err := InitMySQL()
	if err != nil {
		panic(err)
	}

	var book []Book
	err = db.Find(&book).Error
	if err != nil {
		panic(err)
	}

	for _, v := range book {
		marshal, _ := json.Marshal(v)

		err = syncBook(marshal, v.ID)
		if err != nil {
			panic(err)
		}
	}

}

func syncBook(data []byte, id uint) error {
	es, err := InitEs()
	if err != nil {
		return err
	}

	req := esapi.IndexRequest{
		Index:      "book",
		DocumentID: strconv.Itoa(int(id)),
		Body:       bytes.NewReader(data),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	fmt.Println(res)
	return nil
}

func InitMySQL() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/zk2?charset=utf8mb4&parseTime=True&loc=Local"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func InitEs() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://111.231.76.156:9200",
		},
	}

	return elasticsearch.NewClient(cfg)
}
