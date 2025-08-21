package main

import (
    "github.com/gin-gonic/gin"
    "mango-login/internal/handler"
)

func main() {
    r := gin.Default()

    r.POST("/login", handler.LoginHandler)

    r.Run(":8084") // Start server
}