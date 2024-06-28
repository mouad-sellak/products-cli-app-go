package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jlaffaye/ftp"
	"github.com/spf13/cobra"
	"github.com/tealeg/xlsx"
	"golang.org/x/crypto/ssh"
)

var db *sql.DB

var rootCmd = &cobra.Command{
	Use:   "products-cli-app-go",
	Short: "Product Manager CLI Application",
	Long:  `A CLI application to manage products with various functionalities.`,
}

var name, description string
var price float64

func main() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/products_manager_go"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	menu()
}

func menu() {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Ajouter un produit")
		fmt.Println("2. Afficher la liste des produits")
		fmt.Println("3. Modifier un produit")
		fmt.Println("4. Supprimer un produit")
		fmt.Println("5. Exporter les informations produits dans un fichier Excel (en .xlsx)")
		fmt.Println("6. Lancer un serveur Http avec une page web")
		fmt.Println("7. Se connecter à une VM en ssh")
		fmt.Println("8. Se connecter à un serveur FTP")
		fmt.Println("9. Quitter")
		fmt.Println("10. Bonus) Lancer l'interface web sur le port 9000")
		fmt.Print("Choisissez une option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addProduct()
		case 2:
			listProducts()
		case 3:
			updateProduct()
		case 4:
			deleteProduct()
		case 5:
			exportProducts()
		case 6:
			startHTTPServer()
		case 7:
			connectToVM()
		case 8:
			connectToFTP()
		case 9:
			fmt.Println("Quitting...")
			return
		case 10:
			launchWebServer()
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addProduct() {
	fmt.Print("Enter product name: ")
	fmt.Scan(&name)
	fmt.Print("Enter product description: ")
	fmt.Scan(&description)
	fmt.Print("Enter product price: ")
	fmt.Scan(&price)

	_, err := db.Exec("INSERT INTO product (name, description, price) VALUES (?, ?, ?)", name, description, price)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product added successfully.")
}

func listProducts() {
	rows, err := db.Query("SELECT id, name, description, price FROM product")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, description string
		var price float64
		err := rows.Scan(&id, &name, &description, &price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Description: %s, Price: %.2f\n", id, name, description, price)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func updateProduct() {
	var id int
	fmt.Print("Enter product ID to update: ")
	fmt.Scan(&id)
	fmt.Print("Enter new product name: ")
	fmt.Scan(&name)
	fmt.Print("Enter new product description: ")
	fmt.Scan(&description)
	fmt.Print("Enter new product price: ")
	fmt.Scan(&price)

	_, err := db.Exec("UPDATE product SET name = ?, description = ?, price = ? WHERE id = ?", name, description, price, id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product updated successfully.")
}

func deleteProduct() {
	var id int
	fmt.Print("Enter product ID to delete: ")
	fmt.Scan(&id)

	_, err := db.Exec("DELETE FROM product WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product deleted successfully.")
}

func exportProducts() {
	fmt.Println("Exporting products to Excel...")
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Products")
	if err != nil {
		log.Printf(err.Error())
		return
	}
	header := sheet.AddRow()
	header.AddCell().Value = "ID"
	header.AddCell().Value = "Name"
	header.AddCell().Value = "Description"
	header.AddCell().Value = "Price"

	rows, err := db.Query("SELECT id, name, description, price FROM product")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, description string
		var price float64
		err := rows.Scan(&id, &name, &description, &price)
		if err != nil {
			log.Fatal(err)
		}

		row := sheet.AddRow()
		row.AddCell().Value = fmt.Sprintf("%d", id)
		row.AddCell().Value = name
		row.AddCell().Value = description
		row.AddCell().Value = fmt.Sprintf("%.2f", price)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	err = file.Save("products.xlsx")
	if err != nil {
		log.Printf(err.Error())
	} else {
		fmt.Println("Products exported successfully to products.xlsx")
	}
}

func startHTTPServer() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Product CLI APP GO",
		})
	})
	fmt.Println("Starting HTTP server on port 8080...")
	router.Run(":8080")
}

func connectToVM() {
	var address, username, keyPath string
	fmt.Print("Enter VM address : ")
	fmt.Scan(&address)
	fmt.Print("Enter VM port : ")
	var port string
	fmt.Scan(&port)
	fmt.Print("Enter VM username : ")
	fmt.Scan(&username)
	fmt.Print("Enter path to your private key : ")
	fmt.Scan(&keyPath)

	addressWithPort := address + ":" + port

	key, err := os.ReadFile(keyPath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", addressWithPort, config)
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("ls -l"); err != nil {
		log.Fatalf("Failed to run: %v", err)
	}
	fmt.Println(b.String())
	fmt.Println("Successfully connected to the VM and executed the command.")
}

func connectToFTP() {
	var address, username, password, port string
	fmt.Print("Enter FTP server address: ")
	fmt.Scan(&address)
	fmt.Print("Enter FTP port : ")
	fmt.Scan(&port)
	fmt.Print("Enter FTP username: ")
	fmt.Scan(&username)
	fmt.Print("Enter FTP password: ")
	fmt.Scan(&password)

	addressWithPort := address + ":" + port

	conn, err := ftp.Dial(addressWithPort)
	if err != nil {
		log.Fatalf("Failed to connect to FTP server: %v", err)
	}
	defer conn.Quit()

	err = conn.Login(username, password)
	if err != nil {
		log.Fatalf("Failed to login to FTP server: %v", err)
	}
	defer conn.Logout()

	entries, err := conn.List("/")
	if err != nil {
		log.Fatalf("Failed to list directory: %v", err)
	}

	fmt.Println("Directory listing:")
	for _, entry := range entries {
		fmt.Println(entry.Name)
	}
	fmt.Println("Successfully connected to the FTP server and listed the directory.")
}

func launchWebServer() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/products", func(c *gin.Context) {
		var products []map[string]interface{}

		rows, err := db.Query("SELECT id, name, description, price FROM product")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name, description string
			var price float64
			err := rows.Scan(&id, &name, &description, &price)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			product := map[string]interface{}{
				"id":          id,
				"name":        name,
				"description": description,
				"price":       price,
			}
			products = append(products, product)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, products)
	})

	router.POST("/products", func(c *gin.Context) {
		var product struct {
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Price       float64 `json:"price"`
		}

		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("INSERT INTO product (name, description, price) VALUES (?, ?, ?)",
			product.Name, product.Description, product.Price)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Product added successfully"})
	})

	router.PUT("/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		var product struct {
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Price       float64 `json:"price"`
		}

		if err := c.ShouldBindJSON(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.Exec("UPDATE product SET name = ?, description = ?, price = ? WHERE id = ?",
			product.Name, product.Description, product.Price, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Product updated successfully"})
	})

	router.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec("DELETE FROM product WHERE id = ?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Product deleted successfully"})
	})

	router.GET("/export", func(c *gin.Context) {
		file := xlsx.NewFile()
		sheet, err := file.AddSheet("Products")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		header := sheet.AddRow()
		header.AddCell().Value = "ID"
		header.AddCell().Value = "Name"
		header.AddCell().Value = "Description"
		header.AddCell().Value = "Price"

		rows, err := db.Query("SELECT id, name, description, price FROM product")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			var name, description string
			var price float64
			err := rows.Scan(&id, &name, &description, &price)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			row := sheet.AddRow()
			row.AddCell().Value = fmt.Sprintf("%d", id)
			row.AddCell().Value = name
			row.AddCell().Value = description
			row.AddCell().Value = fmt.Sprintf("%.2f", price)
		}

		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = file.Save("products.xlsx")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "Products exported successfully to products.xlsx"})
		}
	})

	fmt.Println("Starting web server on port 9000...")
	router.Run(":9000")
}
