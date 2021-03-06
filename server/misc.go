package server

import(
  "github.com/gin-gonic/gin"
)

func setupMiscRouter(router *gin.Engine) {
  router.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })
}
