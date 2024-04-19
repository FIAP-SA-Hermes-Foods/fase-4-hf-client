package reponosql

import (
	"fase-4-hf-client/internal/core/db"
	"fase-4-hf-client/internal/core/domain/entity/dto"
	"fase-4-hf-client/internal/core/domain/repository"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var _ repository.ClientRepository = (*userDB)(nil)

type userDB struct {
	Database  db.NoSQLDatabase
	tableName string
}

func NewClientRepository(database db.NoSQLDatabase, tableName string) *userDB {
	return &userDB{Database: database, tableName: tableName}
}

func (c *userDB) GetClientByCPF(cpf string) (*dto.ClientDB, error) {
	filter := "cpf = :value"
	attrSearch := map[string]*dynamodb.AttributeValue{
		":value": {
			S: aws.String(cpf),
		},
	}

	input := &dynamodb.ScanInput{
		TableName:                 aws.String(c.tableName),
		FilterExpression:          aws.String(filter),
		ExpressionAttributeValues: attrSearch,
	}

	result, err := c.Database.Scan(input)
	if err != nil {
		return nil, err
	}

	var userList = make([]dto.ClientDB, 0)
	for _, item := range result.Items {
		var c dto.ClientDB
		if err := dynamodbattribute.UnmarshalMap(item, &c); err != nil {
			return nil, err
		}
		userList = append(userList, c)
	}

	if len(userList) > 0 {
		return &userList[0], nil
	}

	return nil, nil
}

func (c *userDB) SaveClient(client dto.ClientDB) (*dto.ClientDB, error) {

	putItem := map[string]*dynamodb.AttributeValue{
		"uuid": {
			S: aws.String(client.UUID),
		},
		"name": {
			S: aws.String(client.Name),
		},
		"cpf": {
			S: aws.String(client.CPF),
		},
		"email": {
			S: aws.String(client.Email),
		},
		"created_at": {
			S: aws.String(client.CreatedAt),
		},
	}

	inputPutItem := &dynamodb.PutItemInput{
		Item:      putItem,
		TableName: aws.String(c.tableName),
	}

	putOut, err := c.Database.PutItem(inputPutItem)

	if err != nil {
		return nil, err
	}

	var out *dto.ClientDB

	if err := dynamodbattribute.UnmarshalMap(putOut.Attributes, &out); err != nil {
		return nil, err
	}

	return out, nil
}
