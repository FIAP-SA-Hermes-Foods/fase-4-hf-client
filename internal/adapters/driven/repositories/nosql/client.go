package reponosql

import (
	"fase-4-hf-client/internal/core/db"
	"fase-4-hf-client/internal/core/domain/entity/dto"
	"fase-4-hf-client/internal/core/domain/repository"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
	attrSearch := map[string]types.AttributeValue{
		":value": &types.AttributeValueMemberS{
			Value: cpf,
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
		if err := attributevalue.UnmarshalMap(item, &c); err != nil {
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

	putItem := map[string]types.AttributeValue{
		"uuid": &types.AttributeValueMemberS{
			Value: client.UUID,
		},
		"name": &types.AttributeValueMemberS{
			Value: client.Name,
		},
		"cpf": &types.AttributeValueMemberS{
			Value: client.CPF,
		},
		"email": &types.AttributeValueMemberS{
			Value: client.Email,
		},
		"createdAt": &types.AttributeValueMemberS{
			Value: client.CreatedAt,
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

	if err := attributevalue.UnmarshalMap(putOut.Attributes, &out); err != nil {
		return nil, err
	}

	out.UUID = client.UUID
	out.Name = client.Name
	out.CPF = client.CPF
	out.Email = client.Email
	out.CreatedAt = client.CreatedAt

	return out, nil
}
