package controllers

import (
	"net/http"
	"encoding/json"

	"github.com/big-vi/go-aws-dynamodb-crud/internal/models"
)

var Orders models.Order

func GetOrders(w http.ResponseWriter, r *http.Request) {
	orders := models.GetOrders()
	res, _ := json.Marshal(orders)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
