package main

import (
	"database/sql"
	"fmt"
)

type Varient struct {
	VarID       int    `json:"VarID"`
	VarientName string `json:"VarientName"`
	ProductID   int    `json:"ProductID"`
	SIZE        string `json:"SIZE"`
	MRP         int    `json:"MRP"`
}

func (varient *Varient) getVarient(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * FROM variety where VarientName='%s'", varient.VarientName)
	fmt.Println(statement)
	return db.QueryRow(statement).Scan(&varient.VarID, &varient.VarientName, &varient.ProductID, &varient.SIZE, &varient.MRP)
}

func (varient *Varient) updateVarient(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE variety SET VarientName='%s'WHERE VarID=%d", varient.VarientName, varient.VarID)
	_, err := db.Exec(statement)
	return err
}

func (varient *Varient) deleteVarient(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM variety WHERE VarID=%d", varient.VarID)
	_, err := db.Exec(statement)
	return err
}

func (varient *Varient) createVarient(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO variety(VarID,VarientName,ProductID,SIZE,MRP) VALUES('%d','%s','%d','%s','%d')",
		varient.VarID, varient.VarientName, varient.ProductID,
		varient.SIZE, varient.MRP)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	//err = db.QueryRow("SELECT LAST_INSERT_VarID()").Scan(&varient.VarID)

	if err != nil {
		return err
	}

	return nil
}
func getVarients(db *sql.DB, start, count int) ([]Varient, error) {
	statement := fmt.Sprintf("SELECT VarID, VarientName, ProductID ,MRP FROM variety LIMIT %d OFFSET %d", count, start)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	Varients := []Varient{}

	for rows.Next() {
		var varient Varient
		if err := rows.Scan(&varient.VarID, &varient.VarientName, &varient.ProductID, &varient.MRP); err != nil {
			return nil, err
		}
		Varients = append(Varients, varient)
	}

	return Varients, nil
}
