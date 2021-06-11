package main

import (
	"fmt"
	"os"

	"gopkg.in/Shopify/sarama.v1"
)

func initConsumer(brokerURLs []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true // turn on errors

	consumer, err := sarama.NewConsumer(brokerURLs, config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func pullQueue(consumer sarama.Consumer, topic string) (*sarama.ConsumerMessage, error) {

	topics, err := consumer.Topics()
	if err != nil {
		return nil, err
	}
	fmt.Println(topics)

	partitions, err2 := consumer.Partitions(topic)
	if err2 != nil {
		return nil, err2
	}
	fmt.Println(partitions)

	// use a lambda function for closing the consumer later
	defer func() {
		if err := consumer.Close(); err != nil {
			panic(err)
		}
	}()

	// like the producer in python, get the oldest record available
	partitionConsumer, err3 := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err3 != nil {
		return nil, err3
	}

	// we have to be able to detect for interrupt
	msgChannel := <-partitionConsumer.Messages()
	return msgChannel, nil
}

func main() {
	fmt.Println("hello")

	brokerURLs := []string{"localhost:9092"}
	consumer, err := initConsumer(brokerURLs)
	if err != nil {
		panic(err)
	}

	consumerMsg, err2 := pullQueue(consumer, "test")
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(consumerMsg.Value)

	os.Exit(0)
}
