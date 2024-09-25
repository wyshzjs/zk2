package main

import (
	"book/db/es"
	"book/db/mysql"
	"book/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	err := mysql.InitMySQL()
	if err != nil {
		panic(err)
	}

	err = es.InitEs()
	if err != nil {
		panic(err)
	}

	g := gin.Default()

	routers.Routers(g)

	err = g.Run(":8080")
	if err != nil {
		panic(err)
	}
}
