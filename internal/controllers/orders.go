package controllers

import(
	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
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