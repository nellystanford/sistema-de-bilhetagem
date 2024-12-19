## Miro Board  
[Access the Miro board here](https://miro.com/app/board/uXjVL2kCRYI=/)

## Steps to Run Locally  

### Start Docker Service  
Run the following command to start the Docker service:  
```bash
docker compose up
```

### Set Default AWS Credentials  
Run the AWS configuration command:  
```bash
aws configure
```

Use the following credentials:  
```plaintext
AWS_ACCESS_KEY_ID="test"
AWS_SECRET_ACCESS_KEY="test"
AWS_DEFAULT_REGION="us-east-1"
```

Alternatively, you can export the credentials directly:  
```bash
export AWS_ACCESS_KEY_ID="test"
export AWS_SECRET_ACCESS_KEY="test"
export AWS_DEFAULT_REGION="us-east-1"
```

### Create DynamoDB Table  
Run the following command to create a DynamoDB table:  
```bash
aws dynamodb create-table --table-name processor.totalCost \
    --attribute-definitions AttributeName=id,AttributeType=S \
    --key-schema AttributeName=id,KeyType=HASH \
    --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5 \
    --endpoint-url http://localhost:8000
```

To check if the table was created, use:  
```bash
aws dynamodb list-tables --endpoint-url http://localhost:8000
```

To scan the table content, use:  
```bash
aws dynamodb scan --table-name processor.totalCost --endpoint-url http://localhost:8000
```

### Initialize Mock Service  
Inside the `mock` folder, run:  
```bash
go run main.go
```

### Create SQS Queue  
Run the following command to create an SQS queue:  
```bash
aws sqs create-queue --endpoint-url http://localhost:4566 --queue-name teste --profile localstack
```

### Send Message to SQS Queue  
To send a message to the SQS queue, run:  
```bash
aws sqs send-message --endpoint-url http://localhost:4566 \
  --queue-url http://localhost:4566/000000000000/teste \
  --message-body file://./message.json --profile localstack
