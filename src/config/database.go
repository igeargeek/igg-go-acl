package config

import (
	"context"
	"fmt"
	"os"

	"github.com/qiniu/qmgo"
)

var database *qmgo.Database = nil
var err error

func init() {
	connect()
}

func connectMongoDB() (*qmgo.Client, error) {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: os.Getenv("DB_CONNECTION")})
	return client, err
}

func connect() *qmgo.Database {
	if database != nil {
		return database
	}
	client, err := connectMongoDB()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("MongoDB is connected!")
	}
	database = client.Database(os.Getenv("DB_NAME"))
	return database
}

func GetDBClient() *qmgo.Database {
	return connect()
}
