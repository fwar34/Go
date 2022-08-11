package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func main() {
	CreateTopic("TEST_FL")
	Producer()
	Consumer()
}

func CreateTopic(topicName string) {
	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver([]string{"192.168.32.228:9876"})))
	if err != nil {
		fmt.Printf("connection error: %s\n", err.Error())
	}
	err = testAdmin.CreateTopic(context.Background(), admin.WithTopicCreate(topicName), admin.WithBrokerAddrCreate("192.168.32.228:10909"))
	if err != nil {
		fmt.Printf("CreateTopic error: %s\n", err.Error())
	}
}

func Producer() {
	p, _ := rocketmq.NewProducer(producer.WithNameServer([]string{"192.168.32.228:9876"}), producer.WithRetry(2), producer.WithGroupName("ProducerGroupName"))
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s\n", err.Error())
		os.Exit(1)
	}

	for i := 0; i < 10; i++ {
		msg := &primitive.Message{
			Topic: "TEST_FL",
			Body:  []byte("Hello RocketMQ Go Client! " + strconv.Itoa(i)),
		}
		msg.WithTag("cdts")
		msg.WithKeys([]string{"test"})
		result, err := p.SendSync(context.Background(), msg)

		if err != nil {
			fmt.Printf("send message error: %s\n", err.Error())
		} else {
			fmt.Printf("send message success: result = %s\n", result.String())
		}
	}
}

func Consumer() {
	c, _ := rocketmq.NewPushConsumer(consumer.WithNameServer([]string{"192.168.32.228:9876"}), consumer.WithConsumerModel(consumer.Clustering), consumer.WithGroupName("ConsumerGroupName"))
	err := c.Subscribe("TEST_FL", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			fmt.Printf("subscribe callback : %v \n", msgs[i])
		}
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		fmt.Printf("subscribe message error: %s\n", err.Error())
	}

	err = c.Start()
	if err != nil {
		fmt.Printf("consumer start error: %s\n", err.Error())
		os.Exit(-1)
	}

	sig := make(chan os.Signal)
	<-sig

	err = c.Shutdown()
	if err != nil {
		fmt.Printf("consumer shutdown error: %s\n", err.Error())
	}
}
