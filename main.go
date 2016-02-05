package main

import(
  "github.com/gorilla/mux"
  "net/http"
  "log"
  "./controllers"
  dbDriver "./db"
)

func main() {
  db := dbDriver.InitConn()
  router := mux.NewRouter()
   router.HandleFunc("/domains", func(w http.ResponseWriter, r *http.Request) { controllers.DomainsIndex(w, r, db) }).Methods("GET")
   router.HandleFunc("/domains", func(w http.ResponseWriter, r *http.Request) { controllers.DomainsIndex(w, r, db) }).Methods("POST")
  log.Fatal(http.ListenAndServe(":8080", router))
}
