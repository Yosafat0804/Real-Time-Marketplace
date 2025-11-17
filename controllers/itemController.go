package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"stock/config"
	"stock/models"
)

func GetItems(w http.ResponseWriter, r *http.Request) {

	var items []models.Item

	config.DB.Find(&items)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {

	var item models.Item

	json.NewDecoder(r.Body).Decode(&item)

	config.DB.Create(&item)

	NotifyAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	var item models.Item

	config.DB.First(&item, id)

	json.NewDecoder(r.Body).Decode(&item)

	config.DB.Save(&item)

	NotifyAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]

	config.DB.Delete(&models.Item{}, id)

	NotifyAll()

	w.Write([]byte("Item deleted"))
}
