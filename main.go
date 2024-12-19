package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/nellystanford/sistema-de-bilhetagem/internal/contract"
	db "github.com/nellystanford/sistema-de-bilhetagem/internal/db/costs"
	"github.com/nellystanford/sistema-de-bilhetagem/internal/usecase/process"
)

const (
	queueUrl string = "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/teste"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config: %v", err)
	}

	// point to localstack
	cfg.BaseEndpoint = aws.String("http://localhost:4566")
	sqsClient := sqs.NewFromConfig(cfg)

	// point to local dynamodb
	cfg.BaseEndpoint = aws.String("http://localhost:8000")
	dbClient := dynamodb.NewFromConfig(cfg)

	receiveMessages(ctx, sqsClient, queueUrl, dbClient)
}

func receiveMessages(ctx context.Context, client *sqs.Client, queueURL string, dbClient *dynamodb.Client) {
	for {
		msgInput := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL),
			MaxNumberOfMessages: 5, // can read up to 10 messages
			VisibilityTimeout:   30,
			WaitTimeSeconds:     10, // long polling to save resources
			MessageAttributeNames: []string{
				"All", // captures all message attributes
			},
		}

		resp, err := client.ReceiveMessage(ctx, msgInput)
		if err != nil {
			log.Printf("error while reading message: %v", err)
			continue
		}

		if len(resp.Messages) == 0 {
			log.Println("no messages found")
			continue
		}

		for _, msg := range resp.Messages {
			consumptionMessage, err := parseMessageBody(*msg.Body)
			if err != nil {
				log.Printf("error parsing message: %v", err)
				continue
			}

			fmt.Println("message parsed successfully:")
			fmt.Printf("tenant_id: %s\n", consumptionMessage.TenantID)
			fmt.Printf("product: %s\n", consumptionMessage.Product)
			fmt.Printf("used_amount: %s\n", consumptionMessage.UsedAmount)
			fmt.Printf("use_unity: %s\n", consumptionMessage.UseUnity)

			result, err := process.ProcessMessage(process.Input(*consumptionMessage))
			if err != nil {
				log.Printf("unable to process message: %v", err)
				continue
			}

			err = db.InsertItem(ctx, dbClient, result)
			if err != nil {
				log.Printf("error publishing results: %v", err)
				continue
			}

			deleteMessage(ctx, client, queueURL, msg.ReceiptHandle)
		}
	}
}

func parseMessageBody(body string) (*contract.ConsumptionMessage, error) {
	var msg contract.ConsumptionMessage
	if err := json.Unmarshal([]byte(body), &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

func deleteMessage(ctx context.Context, client *sqs.Client, queueURL string, receiptHandle *string) {
	_, err := client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: receiptHandle,
	})
	if err != nil {
		log.Printf("error while deleting message: %v", err)
	} else {
		log.Println("message successfully deleted")
	}
}
