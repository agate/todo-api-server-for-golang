package utils

import (
  "sync"

  "gorm.io/gorm"
  "gorm.io/driver/sqlite"

  "me.honghao/helloworld/todo/models"
)

type Todo = models.Todo

var db *gorm.DB
var locker = &sync.Mutex{}

func GetDb() *gorm.DB {
  if db == nil {
    locker.Lock()
    defer locker.Unlock()

    ins, err := gorm.Open(sqlite.Open("db/todo.db"), &gorm.Config{})

    if err != nil {
      panic("failed to connect database")
    } else {
      db = ins
    }

    // Migrate the schema
    db.AutoMigrate(&Todo{})
  }

  return db
}
