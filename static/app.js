document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("addProductForm");
    const productList = document.getElementById("productList");
    const exportButton = document.getElementById("exportButton");

    form.addEventListener("submit", function(event) {
        event.preventDefault();

        const name = document.getElementById("name").value;
        const description = document.getElementById("description").value;
        const price = parseFloat(document.getElementById("price").value);

        if (isNaN(price)) {
            alert("Please enter a valid price.");
            return;
        }

        fetch("/products", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ name, description, price })
        })
        .then(response => response.json())
        .then(data => {
            if (data.status === "Product added successfully") {
                form.reset();
                loadProducts();
            } else {
                alert("Failed to add product: " + data.error);
            }
        })
        .catch(error => console.error("Error:", error));
    });

    exportButton.addEventListener("click", function() {
        window.location.href = "/export";
    });

    function loadProducts() {
        fetch("/products")
        .then(response => response.json())
        .then(data => {
            productList.innerHTML = "";
            data.forEach(product => {
                const row = document.createElement("tr");
                row.innerHTML = `
                    <td>${product.name}</td>
                    <td>${product.description}</td>
                    <td>${product.price.toFixed(2)}</td>
                    <td>
                        <button class="btn btn-warning btn-sm" onclick="editProduct(${product.id})">Edit</button>
                        <button class="btn btn-danger btn-sm" onclick="deleteProduct(${product.id})">Delete</button>
                    </td>
                `;
                productList.appendChild(row);
            });
        })
        .catch(error => console.error("Error:", error));
    }

    window.editProduct = function(id) {
        const name = prompt("Enter new name:");
        const description = prompt("Enter new description:");
        const price = parseFloat(prompt("Enter new price:"));

        if (isNaN(price)) {
            alert("Please enter a valid price.");
            return;
        }

        fetch(`/products/${id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ name, description, price })
        })
        .then(response => response.json())
        .then(data => {
            if (data.status === "Product updated successfully") {
                loadProducts();
            } else {
                alert("Failed to update product: " + data.error);
            }
        })
        .catch(error => console.error("Error:", error));
    };

    window.deleteProduct = function(id) {
        fetch(`/products/${id}`, {
            method: "DELETE"
        })
        .then(response => response.json())
        .then(data => {
            if (data.status === "Product deleted successfully") {
                loadProducts();
            } else {
                alert("Failed to delete product: " + data.error);
            }
        })
        .catch(error => console.error("Error:", error));
    };

    loadProducts();
});
