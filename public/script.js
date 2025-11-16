// WebSocket connection object
let ws;

// ID item yang sedang diedit (null = mode tambah)
let editId = null;

// ============================
//  CONNECT WEBSOCKET
// ============================
function connectWS() {

    // Membuat koneksi WebSocket otomatis mengikuti protokol (ws / wss)
    ws = new WebSocket(
        (location.protocol === "https:" ? "wss://" : "ws://") 
        + location.host 
        + "/ws" // endpoint websocket server
    );

    // Ketika WebSocket berhasil terkoneksi
    ws.onopen = () => console.log("WS Connected");

    // Ketika server mengirimkan pesan (NotifyAll)
    ws.onmessage = (msg) => {
        const data = JSON.parse(msg.data);

        // Jika server kirim event update, lakukan reload data
        if (data.event === "update") loadData();
    };

    // Jika koneksi terputus, reconnect otomatis setiap 1 detik
    ws.onclose = () => {
        console.warn("WS Closed. Reconnecting...");
        setTimeout(connectWS, 1000);
    };
}

// Mulai koneksi websocket saat halaman terbuka
connectWS();

// Deteksi apakah halaman ini adalah admin.html (admin memiliki tombol edit/delete)
const IS_ADMIN = window.location.pathname.includes("admin");

// ============================
//  LOAD DATA dari server
// ============================
function loadData() {

    // Request GET /items ke backend
    fetch("/items")
        .then(res => res.json())
        .then(items => {
            const tbody = document.getElementById("data");
            tbody.innerHTML = ""; // Kosongkan tabel sebelum menambah data baru

            // Loop semua item
            items.forEach(item => {
                const row = document.createElement("tr");

                // Jika halaman admin → tampilkan tombol edit + hapus
                // Jika viewer → hanya tampilkan data
                row.innerHTML = `
                    <td>${item.name}</td>
                    <td>${item.price}</td>
                    <td>${item.qty}</td>

                    ${
                        IS_ADMIN 
                        ?
                        /* Tombol untuk ADMIN saja */
                        `<td>
                            <button class="edit-btn"
                                data-id="${item.id}"                    /* ID item */
                                data-name='${JSON.stringify(item.name)}' /* Nama dibungkus JSON agar aman */
                                data-price="${item.price}"              /* Harga tetap STRING */
                                data-qty="${item.qty}">
                                Edit
                            </button>

                            <button class="delete-btn" onclick="deleteItem(${item.id})">Hapus</button>
                        </td>`
                        :
                        "" // viewer tidak perlu tombol
                    }
                `;

                tbody.appendChild(row);
            });

            // Setelah tabel selesai terisi, pasang event tombol edit
            attachEditEvents();
        })
        .catch(err => console.error("Load error:", err));
}

// Panggil load pertama kali saat halaman dibuka
loadData();

// ============================
//  EVENT HANDLER TOMBOL EDIT
// ============================
function attachEditEvents() {

    // Ambil semua tombol edit
    document.querySelectorAll(".edit-btn").forEach(btn => {

        btn.addEventListener("click", () => {

            const id = btn.dataset.id;

            // JSON.parse untuk nama agar tetap aman meskipun ada karakter khusus
            const name = JSON.parse(btn.dataset.name);

            // price disimpan sebagai STRING agar format seperti "20.000" tidak dirusak JavaScript
            const price = btn.dataset.price;

            // qty tetap angka
            const qty = Number(btn.dataset.qty);

            // Isi form dengan data item yang dipilih
            editItem(id, name, price, qty);
        });
    });
}

// ============================
//  ADD / UPDATE ITEM
// ============================
function addItem() {
    if (!IS_ADMIN) return; // Hanya admin yang boleh menambahkan / mengubah item

    // Ambil nilai form
    const name = document.getElementById("name").value;
    const price = document.getElementById("price").value; // STRING (contoh: "20.000")
    const qty = Number(document.getElementById("qty").value);

    // Validasi wajib isi
    if (!name || !price || !qty) {
        alert("Semua field wajib diisi!");
        return;
    }

    // =====================
    // MODE EDIT (PUT)
    // =====================
    if (editId !== null) {
        fetch(`/items/${editId}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ name, price, qty }) // Kirim data
        }).then(resetForm);

        return;
    }

    // =====================
    // MODE TAMBAH BARU (POST)
    // =====================
    fetch("/items", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, price, qty }) // Kirim data baru
    }).then(resetForm);
}

// ============================
//  DELETE ITEM
// ============================
function deleteItem(id) {
    if (!IS_ADMIN) return; // viewer tidak boleh hapus

    fetch(`/items/${id}`, { method: "DELETE" });
}

// ============================
//  LOAD DATA INTO FORM (EDIT MODE)
// ============================
function editItem(id, name, price, qty) {

    // Simpan ID item yang sedang diedit
    editId = id;

    // Tampilkan data ke input form
    document.getElementById("name").value = name;
    document.getElementById("price").value = price; // STRING tetap utuh
    document.getElementById("qty").value = qty;

    // Ubah tombol dari "Tambah Barang" menjadi "Simpan Perubahan"
    document.querySelector(".btn").innerText = "Simpan Perubahan";
}

// ============================
//  RESET FORM setelah tambah / edit
// ============================
function resetForm() {

    // Kembali ke mode tambah
    editId = null;

    // Bersihkan isi form
    document.getElementById("name").value = "";
    document.getElementById("price").value = "";
    document.getElementById("qty").value = "";

    // Ubah kembali tombol menjadi default
    document.querySelector(".btn").innerText = "Tambah Barang";
}
