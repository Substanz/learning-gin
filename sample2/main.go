package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()

  r.GET("/user/:name", func(c *gin.Context) {
    name := c.Params.ByName("name")
    message := "Hello " + name
    c.String(200, message)
  })

  r.GET("/user/:name/*action", func(c *gin.Context) {
    name := c.Params.ByName("name")
    action := c.Params.ByName("action")
    message := name + " is " + action
    c.String(200, message)
  })

  r.Run(":8080")
}
