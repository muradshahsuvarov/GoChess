package kafka

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func Consume(server string, topic string, partition int32) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V1_0_0_0

	consumer, err := sarama.NewConsumer([]string{server}, config)
	if err != nil {
		log.Fatalln("Failed to start consumer:", err)
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		log.Fatalln("Failed to start partition consumer:", err)
	}

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("Consumed message with offset %d\n", msg.Offset)
			fmt.Printf("Message: %s\n", string(msg.Value))
		}
	}
}
