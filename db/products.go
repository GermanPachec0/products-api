package db

import (
	"database/sql"
	"fmt"
	"os"
	"products-api/models"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func ConnectDB() error {

	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "localhost",
		DBName:               "products",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}
	fmt.Println("Connected!")
	return nil
}

func GetProducts() ([]models.Product, error) {
	var products []models.Product
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		return nil, fmt.Errorf("Error getting products %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var prod models.Product
		if err := rows.Scan(&prod.ID, &prod.Title, &prod.Description, &prod.Price, &prod.DiscountPercentage, &prod.Rating,
			&prod.Stock, &prod.Brand, &prod.Category, &prod.Thumbnail); err != nil {
			return nil, fmt.Errorf("Error Parsing products %v", err)
		}

		images, err := getImagesByProduct(prod.ID)
		if err != nil {
			fmt.Printf("Error getting images %v", err)
		}
		prod.Images = append(prod.Images, images...)
		products = append(products, prod)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error in rows")
	}
	return products, nil
}

func getImagesByProduct(id int) ([]string, error) {
	var img []string
	rows, err := db.Query("SELECT * FROM image WHERE product_id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error getting image %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var i models.Image
		if err := rows.Scan(&i.ID, &i.ProductID, &i.Img); err != nil {
			return nil, fmt.Errorf("error parsing image %v", err)
		}
		img = append(img, i.Img)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error in rows")
	}
	return img, nil
}
