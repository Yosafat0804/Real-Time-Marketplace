package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// Import folder models supaya GORM bisa membaca struct Item
	"stock/models"
)

// DB akan menampung koneksi database yang bisa dipakai di seluruh project
var DB *gorm.DB

func ConnectDB() {

	// DSN (Data Source Name) berisi informasi untuk login ke database.
	// Format umumnya:
	// username:password@tcp(host:port)/nama_database?opsi_tambahan
	//
	// Pada kasus ini:
	// - username = root
	// - password = kosong
	// - host = 127.0.0.1 (lokal)
	// - port = 3306 (port default MySQL)
	// - database = stock_db
	//
	// charset=utf8mb4 digunakan agar semua karakter bisa disimpan tanpa masalah.
	dsn := "root:@tcp(127.0.0.1:3306)/stock_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Membuka koneksi ke MySQL menggunakan GORM.
	// Jika koneksi gagal, program akan berhenti agar tidak memproses data tanpa database.
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal konek ke database:", err)
	}

	// AutoMigrate berfungsi membuat tabel atau menyesuaikan struktur tabel
	// berdasarkan struct yang ada di folder models.
	//
	// Contoh:
	// - Jika struct Item punya field baru, tabel akan otomatis ditambah kolomnya.
	// - Tidak akan menghapus data yang sudah ada.
	//
	// Fitur ini sangat membantu saat masih tahap pengembangan.
	db.AutoMigrate(&models.Item{})

	// Jika berhasil, tampilkan pesan berhasil konek.
	fmt.Println("✅ Database connected")

	// Simpan koneksi database ke variabel global DB
	// supaya bisa diakses controller lainnya tanpa perlu reconnect.
	DB = db
}
