package db

import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

func cursor() *gorm.DB {
  dsn := "mysql:123456@tcp(10.0.90.151:3306)/ysman?charset=utf8mb4&parseTime=True&loc=Local"
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  return db
}
