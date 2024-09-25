package book

import (
	"book/db/es"
	"github.com/gin-gonic/gin"
	"net/http"
)

type QueryBook struct {
	BookTitle string `json:"book_title"`
	Nickname  string `json:"nickname"`
	Isbn      string `json:"isbn"`
}

func GetBookList(c *gin.Context) {
	var queryBook QueryBook
	err := c.ShouldBind(&queryBook)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
		})
		return
	}

	query := GetQuery(queryBook)

	book, err := es.GetBook(query)
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
		"data":    book,
	})

}

func GetQuery(queryBook QueryBook) map[string]interface{} {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
		"sort": map[string]interface{}{
			"cb_time": "desc",
		},
	}

	if queryBook.BookTitle != "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"book_title": queryBook.BookTitle,
				},
			},
			"sort": map[string]interface{}{
				"cb_time": "desc",
			},
		}

	} else if queryBook.Nickname != "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"nickname": queryBook.Nickname,
				},
			},
			"sort": map[string]interface{}{
				"cb_time": "desc",
			},
		}
	} else if queryBook.Isbn != "" {
		query = map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"isbn": queryBook.Isbn,
				},
			},
			"sort": map[string]interface{}{
				"cb_time": "desc",
			},
		}
	}

	return query
}
