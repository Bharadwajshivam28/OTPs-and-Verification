package main

import (
	"Bharadwajshivam28/sms/api"
	"github.com/gin-gonic/gin"
)

func main () {
	router := gin.Default()
	app := api.Config{Router: router}
	app.Routes()
	router.Run(":8000")
}