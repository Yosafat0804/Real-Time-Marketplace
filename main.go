package main

import (
	"log"
	"net/http"
	"stock/config"
	"stock/controllers"
	"stock/routes"
)

func main() {

	config.ConnectDB()

	r := routes.SetupRoutes()

	r.PathPrefix("/public/").Handler(
		http.StripPrefix(
			"/public/",
			http.FileServer(http.Dir("./public")),
		),
	)

	r.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/admin.html")
	})

	r.HandleFunc("/viewer", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/viewer.html")
	})

	r.HandleFunc("/ws", controllers.WebsocketHandler)

	log.Println("Server running at http://localhost:3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}
