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

var _ repository.ClientRepository = (*clientDB)(nil)

type clientDB struct {
	Database  db.NoSQLDatabase
	tableName string
}

func NewClientRepository(database db.NoSQLDatabase, tableName string) *clientDB {
	return &clientDB{Database: database, tableName: tableName}
}

func (p *clientDB) GetClientByID(uuid string) (*dto.ClientDB, error) {
	partitionKeyName := "uuid"

	input := &dynamodb.QueryInput{
		TableName:              aws.String(p.tableName),
		KeyConditionExpression: aws.String("#pk = :value"),
		ExpressionAttributeNames: map[string]string{
			"#pk": partitionKeyName,
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":value": &types.AttributeValueMemberS{Value: uuid},
		},
	}

	result, err := p.Database.Query(input)
	if err != nil {
		return nil, err
	}

	var clientList = make([]dto.ClientDB, 0)
	if result != nil {
		for _, item := range result.Items {
			var pDb dto.ClientDB
			if err := attributevalue.UnmarshalMap(item, &pDb); err != nil {
				return nil, err
			}
			clientList = append(clientList, pDb)
		}
	}

	if len(clientList) > 0 {
		return &clientList[0], nil
	}

	return nil, nil
}

func (c *clientDB) GetClientByCPF(cpf string) (*dto.ClientDB, error) {
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

	var clientList = make([]dto.ClientDB, 0)
	if result != nil {
		for _, item := range result.Items {
			if item == nil {
				continue
			}
			var c dto.ClientDB
			if err := attributevalue.UnmarshalMap(item, &c); err != nil {
				return nil, err
			}
			clientList = append(clientList, c)
		}
	}

	if len(clientList) > 0 {
		return &clientList[0], nil
	}

	return nil, nil
}

func (c *clientDB) SaveClient(client dto.ClientDB) (*dto.ClientDB, error) {

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
