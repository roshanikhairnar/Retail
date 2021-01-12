package main

import (
	"database/sql"
	"fmt"
)

type Product struct {
	ProductID   int    `json:"ProductID"`
	ProductName string `json:"ProductName"`
	SubcategoryID int `json:"SubcategoryID"`
}

func (product *Product) getProduct(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * FROM product where ProductName='%s'", product.ProductName)
	fmt.Println(statement)
	return db.QueryRow(statement).Scan( &product.ProductID,&product.ProductName,&product.SubcategoryID)
}

func (product *Product) updateProduct(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE product SET ProductName='%s'WHERE ProductID=%d", product.ProductName, product.ProductID)
	_, err := db.Exec(statement)
	return err
}

func (product *Product) deleteProduct(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM product WHERE ProductID=%d", product.ProductID)
	_, err := db.Exec(statement)
	return err
}

func (product *Product) createProduct(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO Product(Productname,ProductID,SubcategoryID) VALUES('%s','%d','%d')", product.ProductName, product.ProductID,product.SubcategoryID)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	//err = db.QueryRow("SELECT LAST_INSERT_ProductID()").Scan(&Product.ProductID)

	if err != nil {
		return err
	}

	return nil
}
func getProducts(db *sql.DB, start, count int) ([]Product, error) {
	statement := fmt.Sprintf("SELECT ProductID, ProductName, SubcategoryID FROM product LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Products := []Product{}

	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ProductID, &product.ProductName,&product.SubcategoryID); err != nil {
			return nil, err
		}
		Products = append(Products, product)
	}

	return Products, nil
}
