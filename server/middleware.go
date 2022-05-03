package server

import(
  "github.com/gin-gonic/gin"
)

func setupMiddleware(engine *gin.Engine) {
  engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
}
