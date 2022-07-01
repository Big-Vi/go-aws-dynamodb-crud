package models

import (
	"context"
	"fmt"

	"github.com/big-vi/go-aws-dynamodb-crud/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
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

func GetOrderById(Id string) Order {
	result, err := dynamodbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("my-table"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: Id},
		},
	})

	if err != nil {
		panic(err)
	}

	var order Order

	err = attributevalue.UnmarshalMap(result.Item, &order)
	if err != nil {
		err = fmt.Errorf("failed to marshal Record, %w", err)
		fmt.Println("An error happened:", err)
	}

	return order
}

func DeleteOrder(Id string) *dynamodb.DeleteItemOutput {
	result, err := dynamodbClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String("my-table"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: Id},
		},
	})
	if err != nil {
		panic(err)
	}

	return result
}

func CreateOrder() *dynamodb.PutItemOutput {
	id := uuid.New()
	result, err := dynamodbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("my-table"),
		Item: map[string]types.AttributeValue{
			"id":    &types.AttributeValueMemberS{Value: id.String()},
			"title": &types.AttributeValueMemberS{Value: "John Doe"},
		},
	})

	if err != nil {
		panic(err)
	}

	return result
}

func UpdateOrder(Id string) *dynamodb.PutItemOutput {
	result, err := dynamodbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("my-table"),
		Item: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: Id},
			"title": &types.AttributeValueMemberS{Value: "Peter"},
		},
	})

	if err != nil {
		panic(err)
	}

	return result
}
