package models

import(
  "time"
  _ "github.com/lib/pq"
)

type Model struct {
    id        uint `gorm:"primary_key" form:"id", json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    DeletedAt *time.Time
}
