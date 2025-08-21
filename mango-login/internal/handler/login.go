package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
    var req LoginRequest

    // Bind and validate JSON input
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Missing or invalid username/password",
        })
        return
    }

    // Mock authentication logic
    if req.Username == "kornvipa" && req.Password == "backend123" {
        c.JSON(http.StatusOK, gin.H{
            "message": "Login successful 🎉",
            "user":    req.Username,
        })
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Invalid credentials ❌",
        })
    }
}