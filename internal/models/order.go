package models

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/big-vi/go-aws-dynamodb-crud/config"
)

var dynamodbClient *dynamodb.Client

func init() {
	dynamodbClient = config.ConnectDynamoDB()
}

type Order struct {
	ID    string
	Title string
}

func GetOrders() []Order {
	result, err := dynamodbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("my-table"),
	})
	if err != nil {
		panic(err)
	}

	var orders []Order

	//Using Go SDK helper function to parse dynamodb attributes to struct
	err = attributevalue.UnmarshalListOfMaps(result.Items, &orders)
	if err != nil {
		err = fmt.Errorf("failed to marshal Record, %w", err)
		fmt.Println("An error happened:", err)
	}

	return orders
}
