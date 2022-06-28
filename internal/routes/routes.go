package routes

import (
	"github.com/big-vi/go-aws-dynamodb-crud/internal/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/orders", controllers.GetOrders).Methods("GET")
	router.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", controllers.GetOrderById).Methods("GET")
	router.HandleFunc("/orders/{id}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/orders/{id}", controllers.DeleteOrder).Methods("DELETE")
}
