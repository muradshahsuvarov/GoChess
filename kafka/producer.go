package kafka

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func Produce(server string, topic string, targetPartition int32, message string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{server}, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(message),
		Partition: targetPartition, // Specify the target partition
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalln("Failed to send message:", err)
	}

	fmt.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}

func main() {
	Produce("localhost:9092", "my_topic", 1, "Hello, World!")
}
