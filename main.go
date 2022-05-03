package main

import (
  "github.com/gin-gonic/gin"

  "me.honghao/helloworld/todo/server"
  "me.honghao/helloworld/todo/utils"
)


func main() {
  config := utils.GetConfig()
  engine := gin.New()
  server.Setup(engine)
  engine.Run(config.Address)
}
