package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/davecgh/go-spew/spew"
)

var (
	svc     = dynamodb.New(session.New())
	iso8601 = "2006-01-02T15:04:05.000Z"
)

func main() {
	giftID := "test10"
	t := time.Now().Format(iso8601)
	count := 30000
	chunk := 25
	for i := 0; i < count; i += chunk {
		lastIndex := i + chunk
		jLastIndex := chunk
		if count < lastIndex {
			lastIndex = count
			jLastIndex = count - i
		}
		items := make([]*dynamodb.WriteRequest, jLastIndex)
		for j := 0; j < jLastIndex; j++ {
			userID := fmt.Sprintf("gift_test_user_%d", i+j)
			items[j] = &dynamodb.WriteRequest{
				PutRequest: &dynamodb.PutRequest{
					Item: map[string]*dynamodb.AttributeValue{
						"giftId": {
							S: aws.String(giftID),
						},
						"userId": {
							S: aws.String(userID),
						},
						"createdAt": {
							S: aws.String(t),
						},
						"updatedAt": {
							S: aws.String(t),
						},
						"status": {
							S: aws.String("applying"),
						},
					},
				},
			}
		}
		_, err := svc.BatchWriteItem(&dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]*dynamodb.WriteRequest{
				"devhoshi.client-api.lottery-gifts": items,
			},
		})
		if err != nil {
			spew.Dump(err)
			return
		}
		fmt.Println("%d ~ %d", i, lastIndex)
	}
}
