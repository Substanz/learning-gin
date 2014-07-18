package main

import "github.com/gin-gonic/gin"

type LoginJSON struct {
  User     string `json:"user" binding:"required"`
  Password string `json:"password" binding:"required"`
}

func main() {
  r := gin.Default()

  r.POST("/login", func(c *gin.Context) {
    var json LoginJSON

    if c.EnsureBody(&json) {
      if json.User == "manu" && json.Password == "123" {
        c.JSON(200, gin.H{"status": "you are logged in"})
      } else {
        c.JSON(401, gin.H{"status": "unauthorized"})
      }
    }
  })

  r.Run(":8080")
}

