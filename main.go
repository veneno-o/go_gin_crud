package main

import (
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	t.ParseFiles("./templates/index.html")
	t.Execute(w, "Hello Go Template")
}

func main() {
	//engine := gin.Default()
	//engine.GET("/", index)
	//engine.Run()
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	server.ListenAndServe()
}
