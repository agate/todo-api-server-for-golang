package server

import(
  "github.com/gin-gonic/gin"

  "me.honghao/helloworld/todo/models"
  "me.honghao/helloworld/todo/utils"
)

type Todo = models.Todo

func setupApiRouter(router *gin.Engine) {
  db := utils.GetDb()

  router.GET("/api/todos", func(c *gin.Context) {
    var records []Todo
    db.Find(&records)
    c.JSON(200, gin.H{
      "todos": records,
    })
  })

  router.POST("/api/todos", func(c *gin.Context) {
    Title := c.PostForm("Title")
    db.Create(&Todo{
      Title: Title,
      Completed: false,
    })
    c.JSON(200, gin.H{
      "success": true,
    })
  })

  router.POST("/api/todos/:id/complete", func(c *gin.Context) {
    var todo Todo
    id := c.Param("id")
    db.First(&todo, id)
    db.Model(&todo).Update("Completed", true)
    c.JSON(200, gin.H{
      "success": true,
    })
  })

  router.POST("/api/todos/:id/incomplete", func(c *gin.Context) {
    var record Todo
    id := c.Param("id")
    db.First(&record, id)
    db.Model(&record).Update("Completed", false)
    c.JSON(200, gin.H{
      "success": true,
    })
  })
}
