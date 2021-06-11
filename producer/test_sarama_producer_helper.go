package main

import (
	"fmt"

	"gopkg.in/Shopify/sarama.v1"
)

// NOTE: this file only contains the producer and consumer code
// 			 in other words, this is only the client side used to connect to the Kafka brokers
// 			 we still need to use docker to spin up the brokers

func initProducer(brokerURLs []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	prod, err := sarama.NewSyncProducer(brokerURLs, config)
	if err != nil {
		return nil, err
	}
	return prod, nil
}

func toQueue(topic string, message []byte) error {
	// basically takes a topic and a message array of bytes, returning error only if one occurred, nil otherwise

	// important aspects:
	// 		1. setup brokerURLs
	// 		2. init producer
	// 		3. defer closing of producer
	// 		4. initialize a producer message
	// 		5. push message to queue
	// 		6. print the Kafka specific partition, offset, and topic
	brokerURLs := []string{"192.168.4.25:9092"} // check this to make sure kafka will take this

	producer, err := initProducer(brokerURLs)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("partition: %d, offset: %d, topic: %s\n", partition, offset, topic)
	return nil
}
