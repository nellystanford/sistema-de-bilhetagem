package db

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/nellystanford/sistema-de-bilhetagem/internal/entity"
)

var tableName string = "processor.totalCost"

type TotalCostDBItem struct {
	ID          string    `dynamodbav:"id"`
	Tenant      string    `dynamodbav:"tenant"`
	SpentAmount float64   `dynamodbav:"spent_amount"`
	Product     string    `dynamodbav:"product"`
	Date        time.Time `dynamodbav:"date"`
}

func InsertItem(ctx context.Context, dbClient *dynamodb.Client, input entity.TotalCost) error {
	inputItem := TotalCostDBItem{
		ID:          uuid.New().String(),
		Tenant:      input.Tenant,
		SpentAmount: input.SpentAmount,
		Product:     input.Product,
		Date:        input.Date,
	}

	item, err := attributevalue.MarshalMap(inputItem)
	if err != nil {
		panic(err)
	}

	_, err = dbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName), Item: item,
	})
	if err != nil {
		return err
	}

	fmt.Println("item added successfully to dynamodb table")
	return nil
}
