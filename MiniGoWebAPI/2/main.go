package main

import (
  "fmt"
  "net/http"
)

type Human struct {
  FName string
  LName string
  Age int
}

func (hum Human) ServerHTTP(w http.ResponseWriter,r *http.Request) {
  hum.FName = "Mustafa"
  hum.LName = "Güler"
  hum.Age = 18

  r.ParseForm()

  fmt.Println(r.Form)

  fmt.Println("path",r.URL.Path)  
}

func main(){
	var hum Human
	err := http.ListenAndServe("localhost:8083",hum)
	checkError(err)
}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}
