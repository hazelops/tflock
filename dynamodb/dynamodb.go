package dynamodb

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	uuid "github.com/hashicorp/go-uuid"
)

type Info struct {
	ID        string
	Operation string
	Info      string
	Who       string
	Version   string
	Created   time.Time
	Path      string
}

func Lock(path string, region string, profile string) {
	config := aws.NewConfig().WithRegion(region)

	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  *config,
		Profile: profile,
	})

	if err != nil {
		panic(err)
	}

	tableName := "tf-state-lock"

	lockID, err := uuid.GenerateUUID()
	if err != nil {
		panic(err)
	}

	host, _ := os.Hostname()

	info := Info{
		ID:        lockID,
		Created:   time.Now().UTC(),
		Who:       fmt.Sprintf("tflock@%s", host),
		Version:   "tflock",
		Operation: "tflock",
		Path:      path,
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}

	_, err = dynamodb.New(sess).PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"LockID": {S: &path},
			"Info":   {S: aws.String(string(bytes))},
		},
		TableName: &tableName,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("TF lock successfully!")
}
