package server

import(
  "github.com/gin-gonic/gin"
)

func Setup(engine *gin.Engine) {
  setupMiddleware(engine)
  setupMiscRouter(engine)
  setupApiRouter(engine)
}
