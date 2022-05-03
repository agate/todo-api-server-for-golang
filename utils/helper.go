package utils

import (
  "fmt"
  "os"

  "me.honghao/helloworld/todo/models"
)

const DEFAULT_HOST = "0.0.0.0"
const DEFAULT_PORT = "8080"

type Config = models.Config

func GetEnv(key string, fallback string) string {
  if value, ok := os.LookupEnv(key); ok {
    return value
  }
  return fallback
}

func GetConfig() Config {
  host := GetEnv("TODO_HOST", DEFAULT_HOST)
  port := GetEnv("TODO_PORT", DEFAULT_PORT)
  return Config {
    Host: host,
    Port: port,
    Address: fmt.Sprintf("%s:%s", host, port),
  }
}
