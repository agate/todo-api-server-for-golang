package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

const PORT = 8080

type Todo struct {
  gorm.Model
  Title string
  Completed bool
}

func main() {
  r := gin.Default()

  db, err := gorm.Open(sqlite.Open("db/todo.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&Todo{})

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

  r.GET("/api/todos", func(c *gin.Context) {
    var todos []Todo
    db.Find(&todos)
    c.JSON(200, gin.H{
      "todos": todos,
    })
  })

  r.POST("/api/todos", func(c *gin.Context) {
    Title := c.PostForm("Title")
    db.Create(&Todo{
      Title: Title,
      Completed: false,
    })
    c.JSON(200, gin.H{
      "success": true,
    })
  })

  r.POST("/api/todos/:id/complete", func(c *gin.Context) {
    var todo Todo
    id := c.Param("id")
    db.First(&todo, id)
    db.Model(&todo).Update("Completed", true)
    c.JSON(200, gin.H{
      "success": true,
    })
  })

  r.POST("/api/todos/:id/incomplete", func(c *gin.Context) {
    var todo Todo
    id := c.Param("id")
    db.First(&todo, id)
    db.Model(&todo).Update("Completed", false)
    c.JSON(200, gin.H{
      "success": true,
    })
  })

  r.Run(fmt.Sprintf(":%d", PORT))
}
