
# Products CLI App Go

Products CLI App Go est une application en ligne de commande (CLI) pour la gestion des produits, écrite en Go. Cette application permet d'ajouter, afficher, modifier, supprimer et exporter des informations sur les produits. Elle offre également une interface web pour ces fonctionnalités.

## Fonctionnalités

- Ajouter un produit
- Afficher la liste des produits
- Modifier un produit
- Supprimer un produit
- Exporter les informations des produits dans un fichier Excel (.xlsx)
- Lancer un serveur HTTP avec une page web
- Se connecter à une VM en SSH
- Se connecter à un serveur FTP
- Lancer une interface web sur le port 9000

## Prérequis

- Go 1.16 ou supérieur
- MySQL
- Git

## Installation

1. Clonez le dépôt GitHub :

   ```sh
   git clone https://github.com/mouad-sellak/products-cli-app-go.git
   cd products-cli-app-go
   ```

2. Installez les dépendances nécessaires :

   ```sh
   go get -u github.com/gin-gonic/gin
   go get -u github.com/go-sql-driver/mysql
   go get -u github.com/jlaffaye/ftp
   go get -u github.com/tealeg/xlsx
   go get -u golang.org/x/crypto/ssh
   ```

3. Configurez votre base de données MySQL. Créez une base de données appelée `products_manager_go` et une table `product` :

   ```sql
   CREATE DATABASE products_manager_go;
   USE products_manager_go;
   CREATE TABLE product (
       id INT AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(100),
       description TEXT,
       price FLOAT
   );
   ```

## Utilisation

### Exécuter l'application CLI

1. Compilez et exécutez l'application :

   ```sh
   go run main.go
   ```

2. Suivez les instructions du menu pour utiliser les différentes fonctionnalités.

### Lancer l'interface web

1. Dans le menu de l'application CLI, sélectionnez l'option 10 pour lancer l'interface web.

2. Ouvrez votre navigateur et accédez à `http://localhost:9000` pour utiliser l'interface web.

## Structure du projet

- `main.go`: Fichier principal contenant la logique de l'application.
- `templates/`: Répertoire contenant les fichiers HTML pour l'interface web.
- `static/`: Répertoire contenant les fichiers CSS et JavaScript pour l'interface web.

## Auteur

Mouad Sellak - [mouad-sellak](https://github.com/mouad-sellak)
```
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

```markdown
# Products CLI App Go

Products CLI App Go is a command-line interface (CLI) application for managing products, written in Go. This application allows you to add, display, update, delete, and export product information. It also offers a web interface for these functionalities.

## Features

- Add a product
- Display the list of products
- Update a product
- Delete a product
- Export product information to an Excel file (.xlsx)
- Launch an HTTP server with a web page
- Connect to a VM via SSH
- Connect to an FTP server
- Launch a web interface on port 9000

## Prerequisites

- Go 1.16 or higher
- MySQL
- Git

## Installation

1. Clone the GitHub repository:

   ```sh
   git clone https://github.com/mouad-sellak/products-cli-app-go.git
   cd products-cli-app-go
   ```

2. Install the necessary dependencies:

   ```sh
   go get -u github.com/gin-gonic/gin
   go get -u github.com/go-sql-driver/mysql
   go get -u github.com/jlaffaye/ftp
   go get -u github.com/tealeg/xlsx
   go get -u golang.org/x/crypto/ssh
   ```

3. Set up your MySQL database. Create a database called `products_manager_go` and a `product` table:

   ```sql
   CREATE DATABASE products_manager_go;
   USE products_manager_go;
   CREATE TABLE product (
       id INT AUTO_INCREMENT PRIMARY KEY,
       name VARCHAR(100),
       description TEXT,
       price FLOAT
   );
   ```

## Usage

### Run the CLI Application

1. Compile and run the application:

   ```sh
   go run main.go
   ```

2. Follow the menu instructions to use the different functionalities.

### Launch the Web Interface

1. In the CLI application menu, select option 10 to launch the web interface.

2. Open your browser and go to `http://localhost:9000` to use the web interface.

## Project Structure

- `main.go`: Main file containing the application logic.
- `templates/`: Directory containing HTML files for the web interface.
- `static/`: Directory containing CSS and JavaScript files for the web interface.

## Author

Mouad Sellak - [mouad-sellak](https://github.com/mouad-sellak)
```

You can copy and paste this content into a file named `README.md` in the root directory of your project, then add, commit, and push this file to GitHub:

```sh
git add README.md
git commit -m "Add project README"
git push origin main
```