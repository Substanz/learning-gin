package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()

  r.GET("/someJSON", func(c *gin.Context) {
    c.JSON(200, gin.H{"message": "hey", "status": 200})
  })
  
  r.GET("/moreJSON", func(c *gin.Context) {
    var msg struct {
      Name    string `json:"user"`
      Message string
      Number  int
    }
    msg.Name = "Lena"
    msg.Message = "hey"
    msg.Number = 123

    c.JSON(200, msg)
  })

  r.GET("/someXML", func(c *gin.Context) {
    c.XML(200, gin.H{"message": "hey", "status": 200})
  })

  r.Run(":8080")
}
