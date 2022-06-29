package controllers

import(
	"net/http"
	"context"
	"fmt"

	"github.com/big-vi/go-aws-dynamodb-crud/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	svc := config.ConnectDynamoDB()

	items, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
        TableName: aws.String("my-table"),
    })
    if err != nil {
        panic(err)
    }

    fmt.Println(items.Items)

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