package main

import "github.com/gin-gonic/gin"

func main() {
  r := gin.Default()
  r.LoadHTMLTemplates("templates/*")
  r.GET("/index", func(c *gin.Context) {
    obj := gin.H{"title": "Main website"}
    c.HTML(200, "index.tmpl", obj)
  })

  r.Run(":8080")
}
