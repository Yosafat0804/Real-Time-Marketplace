package models

// ============================================================================
// Struct Item
// -----------------------------------------------------------------------------
// Struct ini adalah bentuk representasi dari tabel "items" yang ada di database.
// GORM akan membaca struktur ini dan otomatis membuat tabel tersebut jika belum
// ada, atau menyesuaikan kolom jika ada perubahan.
//
// Setiap field di struct ini akan menjadi kolom di tabel database.
// Selain itu, tag json:"..." dipakai agar data bisa dikirim/diterima dalam
// format JSON ke frontend (browser).
// ============================================================================
type Item struct {

	// ID → primary key (kunci utama) dari tabel.
	// Tipe datanya uint karena GORM akan menggunakannya sebagai auto-increment.
	// Tag json:"id" memastikan bahwa ketika data dikirim ke frontend,
	// field ini ditampilkan sebagai "id".
	ID uint `json:"id" gorm:"primaryKey"`

	// Name → menyimpan nama barang.
	// Misalnya: "Laptop", "Meja", "Botol Minum".
	// Tag json:"name" menentukan nama field saat data diubah ke JSON.
	Name string `json:"name"`

	// Price → menyimpan harga barang.
	// Tipe datanya string, bukan integer, supaya bisa menampung format harga
	// seperti "20.000" atau "1.500.000". Jika integer, titiknya akan hilang.
	// Contoh nilai: "50.000".
	Price string `json:"price"`

	// Qty → jumlah atau stok barang.
	// Tipe integer karena jumlah selalu berupa angka.
	Qty int `json:"qty"`
}
