package es

import (
	"book/global"
	"book/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func GetBook(query map[string]interface{}) (*[]model.Book, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(global.INDEX_NAME),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	r := map[string]interface{}{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	var books []model.Book
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		v := hit.(map[string]interface{})["_source"]

		strTime := fmt.Sprintf("%v %v",
			v.(map[string]interface{})["cb_time"].(string)[0:10],
			v.(map[string]interface{})["cb_time"].(string)[11:19],
		)

		cbTime, err_ := time.Parse(time.DateTime, strTime)
		if err_ != nil {
			return nil, err_
		}

		books = append(books, model.Book{
			Model:     gorm.Model{ID: uint(v.(map[string]interface{})["ID"].(float64))},
			BookTitle: v.(map[string]interface{})["book_title"].(string),
			Nickname:  v.(map[string]interface{})["nickname"].(string),
			Isbn:      v.(map[string]interface{})["isbn"].(string),
			Price:     v.(map[string]interface{})["price"].(float64),
			CbTime:    cbTime,
		})
	}

	return &books, err
}
