package routes

import (
	"github.com/gorilla/mux"
	"stock/controllers"
)

func SetupRoutes() *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/items", controllers.GetItems).Methods("GET")

	r.HandleFunc("/items", controllers.CreateItem).Methods("POST")

	r.HandleFunc("/items/{id}", controllers.UpdateItem).Methods("PUT")

	r.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")

	return r
}
