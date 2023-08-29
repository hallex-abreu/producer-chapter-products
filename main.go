package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	queueURL := os.Getenv("QUEUE_URL")
	region := os.Getenv("AWS_REGION")

	// Configuração da sessão AWS
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		fmt.Println("Erro ao criar sessão:", err)
		return
	}

	// Crie um cliente SQS
	svc := sqs.New(sess)

	uuid := uuid.New()

	// Mensagem que você deseja enviar
	messageBody := fmt.Sprintf(`{
		"id": "%s",
		"user_id": 1,
		"name": "hallex",
		"target": "product-ms",
		"action": "update",
		"old_value": "{id:4848,name:Coca cola,amount:2.56}",
		"new_value": "{id:4848,name:Coca cola,amount:3.15}",
		"created_at": "2012-01-02T15:04:05"
	}`, uuid.String())

	// Envie a mensagem para a fila SQS
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: &messageBody,
		QueueUrl:    &queueURL,
	})
	if err != nil {
		fmt.Println("Erro ao enviar mensagem:", err)
		return
	}

	fmt.Println("Mensagem enviada com sucesso. ID:", *result.MessageId)
}
