package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {

}

func CreateTopic(topicName string) {
    endPoint := []string{"127.0.0.1:9876"}
    testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(endPoint)))
	if err != nil {
        fmt.Printf("connection error: %s\n", err.Error())
	}
    err = testAdmin.CreateTopic(context.Background(), admin.WithTopicCreate(topicName))
    if err != nil {
        fmt.Printf("CreateTopic error: %s\n", err.Error())
    }
}
