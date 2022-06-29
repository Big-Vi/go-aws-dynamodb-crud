package main

import (
	// "context"
	// "fmt"
	"log"
	"net/http"

	"github.com/big-vi/go-aws-dynamodb-crud/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	// out, err := svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
    //     AttributeDefinitions: []types.AttributeDefinition{
    //         {
    //             AttributeName: aws.String("id"),
    //             AttributeType: types.ScalarAttributeTypeS,
    //         },
    //     },
    //     KeySchema: []types.KeySchemaElement{
    //         {
    //             AttributeName: aws.String("id"),
    //             KeyType:       types.KeyTypeHash,
    //         },
    //     },
    //     TableName:   aws.String("my-table"),
    //     BillingMode: types.BillingModePayPerRequest,
    // })
    // if err != nil {
    //     panic(err)
    // }
	// fmt.Print(out)

	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
