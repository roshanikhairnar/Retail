package main

import (
	"database/sql"
	"fmt"
)

type Category struct {
	CategoryID int    `json:"CategoryID"`
	Name       string `json:"name"`
}

func (category *Category) getCategory(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT name,CategoryID FROM Category where name='%s'", category.Name)
	fmt.Println(statement)
	return db.QueryRow(statement).Scan(&category.Name, &category.CategoryID)
}

func (category *Category) updateCategory(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE Category SET name='%s'WHERE CategoryID=%d", category.Name, category.CategoryID)
	_, err := db.Exec(statement)
	return err
}

func (category *Category) deleteCategory(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM Category WHERE CategoryID=%d", category.CategoryID)
	_, err := db.Exec(statement)
	return err
}

func (category *Category) createCategory(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO Category(name,CategoryID) VALUES('%s','%d')", category.Name, category.CategoryID)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	//err = db.QueryRow("SELECT LAST_INSERT_CategoryID()").Scan(&category.CategoryID)

	if err != nil {
		return err
	}

	return nil
}

func getCategorys(db *sql.DB, start, count int) ([]Category, error) {
	statement := fmt.Sprintf("SELECT CategoryID, name FROM Category LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Categorys := []Category{}

	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.CategoryID, &category.Name); err != nil {
			return nil, err
		}
		Categorys = append(Categorys, category)
	}

	return Categorys, nil
}
