package controllers

import(
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "../models"
)
type Response struct {
  Type string `json:"type"`
  Data []models.Domain `json:"data"`
}

func DomainsIndex(w http.ResponseWriter, r *http.Request, db gorm.DB) {
  domains := []models.Domain{}
  db.Find(&domains)
  response := Response{"domains", domains}
  data, _ := json.Marshal(response)
  w.Header().Set("Content-Type", "application/json")
  w.Write(data)
}

func DomainsCreate(w http.ResponseWriter, r *http.Request, db gorm.DB) {
  params := mux.Vars(r)
  domain := models.Domain{params["url"], params["active"], params["priority"]}
  db.Create(domain)

  
}
