package models

import(
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
)

type Domain struct {
  gorm.Model
  Url string `form:"url" binding:"required" json:"url"`
  Active bool `json:"active"`
  Priority uint64 `json:"priority"`
}
