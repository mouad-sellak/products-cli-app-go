document.addEventListener("DOMContentLoaded", function() {
    const form = document.getElementById("addProductForm");
    const productList = document.getElementById("productList");

    form.addEventListener("submit", function(event) {
        event.preventDefault();

        const name = document.getElementById("name").value;
        const description = document.getElementById("description").value;
        const price = document.getElementById("price").value;

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
                loadProducts();
            }
        });
    });

    function loadProducts() {
        fetch("/products")
        .then(response => response.json())
        .then(data => {
            productList.innerHTML = "";
            data.forEach(product => {
                const li = document.createElement("li");
                li.textContent = `${product.name} - ${product.description} - $${product.price.toFixed(2)}`;
                productList.appendChild(li);
            });
        });
    }

    loadProducts();
});
