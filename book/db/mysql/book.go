package mysql

import "book/model"

func GetBookBYId(id int) (book model.Book, err error) {
	err = db.Limit(1).Find(&book, id).Error
	return
}
