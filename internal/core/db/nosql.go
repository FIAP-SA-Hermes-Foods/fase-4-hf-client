package db

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type NoSQLDatabase interface {
	Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}
