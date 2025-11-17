let ws;

let editId = null;

function connectWS() {

    ws = new WebSocket(
        (location.protocol === "https:" ? "wss://" : "ws://") 
        + location.host 
        + "/ws"
    );

    ws.onopen = () => console.log("WS Connected");

    ws.onmessage = (msg) => {
        const data = JSON.parse(msg.data);

        if (data.event === "update") loadData();
    };

    ws.onclose = () => {
        console.warn("WS Closed. Reconnecting...");
        setTimeout(connectWS, 1000);
    };
}

connectWS();

const IS_ADMIN = window.location.pathname.includes("admin");

function loadData() {

    fetch("/items")
        .then(res => res.json())
        .then(items => {
            const tbody = document.getElementById("data");
            tbody.innerHTML = "";

            items.forEach(item => {
                const row = document.createElement("tr");

                row.innerHTML = `
                    <td>${item.name}</td>
                    <td>${item.price}</td>
                    <td>${item.qty}</td>

                    ${
                        IS_ADMIN 
                        ?

                        `<td>
                            <button class="edit-btn"
                                data-id="${item.id}"                    
                                data-name='${JSON.stringify(item.name)}' 
                                data-price="${item.price}"              
                                data-qty="${item.qty}">
                                Edit
                            </button>

                            <button class="delete-btn" onclick="deleteItem(${item.id})">Hapus</button>
                        </td>`
                        :
                        "" 
                    }
                `;

                tbody.appendChild(row);
            });

            attachEditEvents();
        })
        .catch(err => console.error("Load error:", err));
}

loadData();

function attachEditEvents() {

    document.querySelectorAll(".edit-btn").forEach(btn => {

        btn.addEventListener("click", () => {

            const id = btn.dataset.id;

            const name = JSON.parse(btn.dataset.name);

            const price = btn.dataset.price;

            const qty = Number(btn.dataset.qty);

            editItem(id, name, price, qty);
        });
    });
}

function addItem() {
    if (!IS_ADMIN) return;

    
    const name = document.getElementById("name").value;
    const price = document.getElementById("price").value;
    const qty = Number(document.getElementById("qty").value);

    if (!name || !price || !qty) {
        alert("Semua field wajib diisi!");
        return;
    }

    if (editId !== null) {
        fetch(`/items/${editId}`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ name, price, qty })
        }).then(resetForm);

        return;
    }


    fetch("/items", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, price, qty })
    }).then(resetForm);
}

function deleteItem(id) {
    if (!IS_ADMIN) return;

    fetch(`/items/${id}`, { method: "DELETE" });
}

function editItem(id, name, price, qty) {

    editId = id;

    document.getElementById("name").value = name;
    document.getElementById("price").value = price;
    document.getElementById("qty").value = qty;

    document.querySelector(".btn").innerText = "Simpan Perubahan";
}

function resetForm() {

    editId = null;

    document.getElementById("name").value = "";
    document.getElementById("price").value = "";
    document.getElementById("qty").value = "";

    document.querySelector(".btn").innerText = "Tambah Barang";
}
