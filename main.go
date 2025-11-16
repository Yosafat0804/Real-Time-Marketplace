package main

import (
	"log"
	"net/http"
	"stock/config"
	"stock/controllers"
	"stock/routes"
)

func main() {

	// =====================================================
	// 1. MENGHUBUNGKAN APLIKASI DENGAN DATABASE
	// =====================================================
	// Bagian ini menjalankan fungsi ConnectDB() agar aplikasi terhubung
	// ke database. Dengan begitu, semua proses seperti tambah data,
	// edit data, dan hapus data bisa dilakukan.
	config.ConnectDB()

	// =====================================================
	// 2. MEMBUAT ROUTER UTAMA UNTUK MENANGANI REQUEST
	// =====================================================
	// Router adalah "pengatur jalan" yang menentukan:
	// - jika user membuka URL tertentu → fungsi apa yang dijalankan?
	// SetupRoutes() berisi daftar semua endpoint API (GET, POST, dsb).
	r := routes.SetupRoutes()

	// =====================================================
	// 3. MENGIZINKAN AKSES STATIC FILES (HTML, CSS, JS)
	// =====================================================
	// Folder /public berisi file tampilan seperti admin.html, viewer.html,
	// style.css, dan script.js.
	// Bagian ini memastikan file-file tersebut bisa diakses melalui URL:
	// contoh: http://localhost:3000/public/style.css
	r.PathPrefix("/public/").Handler(
		http.StripPrefix(
			"/public/",
			http.FileServer(http.Dir("./public")), // mengambil file dari folder public
		),
	)

	// =====================================================
	// 4. MENAMPILKAN HALAMAN HTML UNTUK ADMIN DAN VIEWER
	// =====================================================

	// Halaman untuk admin (bisa tambah/edit/hapus data)
	// URL: http://localhost:3000/admin
	r.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/admin.html")
	})

	// Halaman viewer hanya bisa melihat data, tanpa bisa mengedit.
	// URL: http://localhost:3000/viewer
	r.HandleFunc("/viewer", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/viewer.html")
	})

	// =====================================================
	// 5. WEBSOCKET – UNTUK UPDATE DATA SECARA REALTIME
	// =====================================================
	// Endpoint ini digunakan supaya admin dan viewer bisa langsung
	// melihat perubahan data secara realtime tanpa harus refresh manual.
	// Contohnya: admin menambah barang → viewer otomatis ikut berubah.
	r.HandleFunc("/ws", controllers.WebsocketHandler)

	// =====================================================
	// 6. MENJALANKAN SERVER DI PORT 3000
	// =====================================================
	log.Println("Server running at http://localhost:3000")

	// Menjalankan server dan "mengikatnya" dengan router yang sudah dibuat di atas.
	// Jika server mengalami error fatal, aplikasi akan dihentikan.
	log.Fatal(http.ListenAndServe(":3000", r))
}
