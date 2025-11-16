package controllers

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// ====================================================================
// Clients
// --------------------------------------------------------------------
// Variabel ini digunakan untuk menyimpan semua koneksi WebSocket
// yang sedang aktif.
// 
// Jadi setiap kali ada user yang membuka halaman viewer/admin,
// browser mereka akan membuat koneksi WebSocket ke server,
// dan koneksi tersebut kita simpan di sini.
// 
// Map berisi:
//   key   = koneksi WebSocket milik client
//   value = boolean (hanya sebagai penanda, tidak dipakai)
// ====================================================================
var Clients = make(map[*websocket.Conn]bool)

// ====================================================================
// Upgrader
// --------------------------------------------------------------------
// Bagian ini bertugas mengubah koneksi HTTP biasa menjadi WebSocket.
// 
// CheckOrigin di-set menjadi true supaya semua domain (localhost,
// 127.0.0.1, dll.) diperbolehkan melakukan koneksi WebSocket.
// ====================================================================
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ====================================================================
// WebsocketHandler
// --------------------------------------------------------------------
// Fungsi ini akan dijalankan jika ada client membuka koneksi ke
// endpoint  /ws
// 
// Di sini:
// 1. HTTP di-upgrade menjadi WebSocket
// 2. Koneksi disimpan ke map Clients
// 3. Server terus menunggu pesan dari client
// 4. Jika client menutup tab atau disconnect, server akan menghapus
//    koneksi tersebut dari daftar Clients.
// ====================================================================
func WebsocketHandler(w http.ResponseWriter, r *http.Request) {

	// Mengubah koneksi HTTP → WebSocket
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return // jika terjadi error, fungsi langsung berhenti
	}

	// Menyimpan koneksi client ke daftar
	Clients[conn] = true

	// Server akan terus menunggu pesan dari client
	for {
		// ReadMessage hanya untuk mendeteksi apakah client masih terhubung.
		// Kita tidak memakai isi pesannya, jadi dibiarkan kosong.
		_, _, err := conn.ReadMessage()

		// Jika terjadi error, berarti client sudah disconnect
		if err != nil {
			// Hapus dari daftar client
			delete(Clients, conn)

			// Tutup koneksi WebSocket
			conn.Close()
			break
		}
	}
}

// ====================================================================
// NotifyAll
// --------------------------------------------------------------------
// Fungsi ini dipanggil setiap kali ada perubahan data barang:
//  - menambah item (Create)
//  - mengupdate item (Update)
//  - menghapus item (Delete)
// 
// Tugas fungsi ini adalah mengirimkan pesan ke semua client
// WebSocket yang sedang terhubung agar mereka memperbarui data
// di halaman masing-masing secara real-time.
// ====================================================================
func NotifyAll() {
	for client := range Clients {
		// Mengirim pesan JSON sederhana ke setiap client
		client.WriteJSON(map[string]string{
			"event": "update", // client nanti menangkap sinyal ini
		})
	}
}
