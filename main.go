package main

import (
	"log"
	"net/http"

	"github.com/dougmendes/snoopy/controller"
)

func main() {
	http.HandleFunc("/", controller.Home)
	http.HandleFunc("/echo", controller.Echo)
	http.HandleFunc("/readjson", controller.ReadJSON)
	log.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
