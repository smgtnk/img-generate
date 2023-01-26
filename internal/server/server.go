package server

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", imgHandler)
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/robots.txt", robotsHandler)

	log.Println("Server starting ...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func rend(w http.ResponseWriter, msg string) {
	if _, err := w.Write([]byte(msg)); err != nil {
		fmt.Println(err)
	}
}

func imgHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "img")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "favicon")
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "robots")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	rend(w, "PONG")
}
