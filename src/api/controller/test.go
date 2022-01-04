package controller

import (
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Print("Request recibido")

}
