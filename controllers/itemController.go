package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"stock/config"
	"stock/models"
)

// ==========================================
// GET /items → Mengambil seluruh daftar barang
// ==========================================
func GetItems(w http.ResponseWriter, r *http.Request) {

	// Menyiapkan wadah untuk menampung semua data barang
	var items []models.Item

	// Mengambil seluruh data dari tabel items di database
	config.DB.Find(&items)

	// Memberi tahu browser bahwa data yang dikirim berbentuk JSON
	w.Header().Set("Content-Type", "application/json")

	// Mengubah data Go menjadi JSON dan mengirimkannya ke frontend
	json.NewEncoder(w).Encode(items)
}

// ==========================================
// POST /items → Menambah barang baru
// ==========================================
func CreateItem(w http.ResponseWriter, r *http.Request) {

	// Tempat menyimpan data barang yang dikirim dari frontend (format JSON)
	var item models.Item

	// Mengubah data JSON dari request menjadi struct item
	json.NewDecoder(r.Body).Decode(&item)

	// Menyimpan data item baru ke dalam database
	config.DB.Create(&item)

	// Mengirim notifikasi agar halaman admin/viewer melakukan update data
	NotifyAll()

	// Mengembalikan data item yang baru disimpan ke frontend
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// ==========================================
// PUT /items/{id} → Mengubah data barang berdasarkan ID
// ==========================================
func UpdateItem(w http.ResponseWriter, r *http.Request) {

	// Mengambil nilai ID dari URL, misalnya /items/5 → id = 5
	id := mux.Vars(r)["id"]

	// Tempat menyimpan data item yang akan diupdate
	var item models.Item

	// Mengambil data item dari database berdasarkan ID
	config.DB.First(&item, id)

	// Mengubah data JSON yang dikirim frontend menjadi update untuk item
	json.NewDecoder(r.Body).Decode(&item)

	// Menyimpan perubahan data ke database
	config.DB.Save(&item)

	// Memberi tahu semua client agar memperbarui tampilan
	NotifyAll()

	// Mengirim data terbaru kembali ke frontend
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// ==========================================
// DELETE /items/{id} → Menghapus barang berdasarkan ID
// ==========================================
func DeleteItem(w http.ResponseWriter, r *http.Request) {

	// Mengambil ID item dari URL
	id := mux.Vars(r)["id"]

	// Menghapus data item dari database berdasarkan ID
	config.DB.Delete(&models.Item{}, id)

	// Mengirim notifikasi update ke semua halaman viewer/admin
	NotifyAll()

	// Mengirim pesan sederhana bahwa item berhasil dihapus
	w.Write([]byte("Item deleted"))
}
