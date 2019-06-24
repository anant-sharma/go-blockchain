package router

import (
	"log"
	"strconv"

	"github.com/anant-sharma/go-boilerplate/config"
	v1Router "github.com/anant-sharma/go-boilerplate/routes/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter Definition
func InitRouter() {

	config := config.GetConfig()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1Router.InitRouter(v1)

	/*
		Start Server
	*/
	err := r.Run(":" + strconv.Itoa(config.PORT))
	if err != nil {
		log.Fatal(err)
	}

}
