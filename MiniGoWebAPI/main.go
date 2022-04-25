package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter,r *http.Request){
  fmt.Fprintf(w,"Hello %s",r.URL.Path[1:])
}

func main(){
  http.HandlerFunc("/",handler)
  http.ListenAndServe(":8083",nil)

  fmt.Println("web server")
}
