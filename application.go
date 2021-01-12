// Shopalystlication.go

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Shopalyst struct {
	Router *mux.Router
	DB     *sql.DB
}

func (shop *Shopalyst) Initialize(Category, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", Category, password, dbname)
	fmt.Println(connectionString)
	var err error
	shop.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	shop.Router = mux.NewRouter()
	shop.initializeRoutes()
}

func (shop *Shopalyst) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, shop.Router))
}

func (shop *Shopalyst) initializeRoutes() {
	shop.Router.HandleFunc("/Categorys", shop.getCategorys).Methods("GET")
	shop.Router.HandleFunc("/Category", shop.createCategory).Methods("POST")
	shop.Router.HandleFunc("/Category/{name}", shop.getCategory).Methods("GET")
	shop.Router.HandleFunc("/Category/{id:[0-9]+}", shop.updateCategory).Methods("PUT")
	shop.Router.HandleFunc("/Category/{id:[0-9]+}", shop.deleteCategory).Methods("DELETE")

	shop.Router.HandleFunc("/Products", shop.getProducts).Methods("GET")
	shop.Router.HandleFunc("/Product", shop.createProduct).Methods("POST")
	shop.Router.HandleFunc("/Products/{ProductName}", shop.getProduct).Methods("GET")
	shop.Router.HandleFunc("/Products/{id}", shop.updateProduct).Methods("PUT")
	shop.Router.HandleFunc("/Products/{id}", shop.deleteProduct).Methods("DELETE")
	
	shop.Router.HandleFunc("/varients", shop.getVarients).Methods("GET")
	shop.Router.HandleFunc("/varient", shop.createVarient).Methods("POST")
	shop.Router.HandleFunc("/varient/{VarientName}", shop.getVarient).Methods("GET")
	shop.Router.HandleFunc("/varient/{id:[0-9]+}", shop.updateVarient).Methods("PUT")
	shop.Router.HandleFunc("/varient/{id}", shop.deleteVarient).Methods("DELETE")
}

func (shop *Shopalyst) getCategorys(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getCategorys(shop.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}
func (shop *Shopalyst) getProducts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getProducts(shop.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func (shop *Shopalyst) createCategory(w http.ResponseWriter, r *http.Request) {
	var category Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := category.createCategory(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, category)
}

func (shop *Shopalyst) getCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	category := Category{Name: name}
	if err := category.getCategory(shop.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Category not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, category)
}

func (shop *Shopalyst) updateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	var category Category
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	category.CategoryID = id

	if err := category.updateCategory(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, category)
}

func (shop *Shopalyst) deleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Category ID")
		return
	}

	category := Category{CategoryID: id}
	if err := category.deleteCategory(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "Shopalystlication/json")
	w.WriteHeader(code)
	w.Write(response)
}
func (shop *Shopalyst) createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := product.createProduct(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, product)
}

func (shop *Shopalyst) getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["ProductName"]
	fmt.Println(name)
	product := Product{ProductName: name}
	if err := product.getProduct(shop.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, product)
}

func (shop *Shopalyst) updateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var product Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	product.ProductID = id

	if err := product.updateProduct(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, product)
}

func (shop *Shopalyst) deleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product := Product{ProductID: id}
	if err := product.deleteProduct(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (shop *Shopalyst) createVarient(w http.ResponseWriter, r *http.Request) {
	var varient Varient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&varient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := varient.createVarient(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, varient)
}

func (shop *Shopalyst) getVarient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["VarientName"]
	fmt.Println(name)
	varient := Varient{VarientName: name}
	if err := varient.getVarient(shop.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "var not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, varient)
}

func (shop *Shopalyst) updateVarient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	fmt.Println(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid var ID")
		return
	}

	var varient Varient
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&varient); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	varient.VarID = id

	if err := varient.updateVarient(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, varient)
}

func (shop *Shopalyst) deleteVarient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	varient := Varient{VarID: id}
	if err := varient.deleteVarient(shop.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (shop *Shopalyst) getVarients(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	varients, err := getVarients(shop.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, varients)
}
