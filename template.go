package main

import (
  "html/template"
  "log"
  "net/http"
  "time"
)

type PageVariables struct {
  Date string
  Time string
}

func main() {
  http.HandleFunc("/", HomePage)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request){
  now := time.Now()
  HomePageVars := PageVariables{
    Date: now.Format("06-06-2020"),
    Time: now.Format("12:04:04"),
  }
  t, err := template.ParseFiles("homepage.html")
  if err != nil {
    log.Print("Template parsing error: ", err)
  }
  err = t.Execute(w, HomePageVars)
  if err != nil {
    log.Print("Template executing error: ", err)
  }
}
