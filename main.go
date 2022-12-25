package main

import (
    "html/template"
    "log"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("index.html")
  if err != nil {
		log.Println(err)
  }
  t.Execute(w, "トップページです！")
}

func hello(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("hello.html")
  if err != nil {
		log.Println(err)
  }
  t.Execute(w, "こんにちは！！")
}

func main() {
	http.HandleFunc("/", index)
  http.HandleFunc("/hello", hello)

  http.ListenAndServe(":8080", nil)
}