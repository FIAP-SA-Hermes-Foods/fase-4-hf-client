package db

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type NoSQLDatabase interface {
	Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
	Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error)
}
