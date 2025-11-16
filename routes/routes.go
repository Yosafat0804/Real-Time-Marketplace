package routes

import (
	"github.com/gorilla/mux"
	"stock/controllers"
)

// SetupRoutes berfungsi untuk membuat dan mengatur semua alamat URL (route)
// yang dapat diakses oleh frontend. Intinya, bagian ini menentukan:
// "Jika user membuka URL tertentu, fungsi backend mana yang harus dijalankan?"
func SetupRoutes() *mux.Router {

	// Membuat router baru menggunakan Gorilla Mux.
	// Router ini bertugas menampung semua endpoint API yang kita miliki.
	r := mux.NewRouter()

	// ==========================================
	//              DAFTAR ROUTE API
	// Setiap route berhubungan dengan proses
	// CRUD (Create, Read, Update, Delete) data.
	// ==========================================

	// GET /items
	// Route ini digunakan untuk mengambil semua data barang
	// dari database. Biasanya dipakai untuk menampilkan tabel barang.
	r.HandleFunc("/items", controllers.GetItems).Methods("GET")

	// POST /items
	// Route untuk menambah barang baru.
	// Data dikirim dari frontend dalam bentuk JSON,
	// lalu disimpan ke database oleh controller CreateItem.
	r.HandleFunc("/items", controllers.CreateItem).Methods("POST")

	// PUT /items/{id}
	// Route untuk mengubah (mengupdate) data barang tertentu berdasarkan ID.
	// Route ini dipakai ketika admin menekan tombol "Edit".
	r.HandleFunc("/items/{id}", controllers.UpdateItem).Methods("PUT")

	// DELETE /items/{id}
	// Route untuk menghapus data barang berdasarkan ID.
	// Digunakan ketika admin menekan tombol "Hapus".
	r.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")

	// Mengembalikan router yang sudah berisi semua route di atas.
	// Router ini nantinya dipakai oleh main.go untuk menjalankan server.
	return r
}
