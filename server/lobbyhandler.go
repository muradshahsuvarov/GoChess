package server

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func JoinRoom(roomId string) (bool, error) {
	// TODO: Room joining logic
	return false, fmt.Errorf("Couldn't join the room %s", roomId)
}

func ListRooms(_server string) {

	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0 // Adjust according to your Kafka version

	client, err := sarama.NewClient([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	partitions, err := client.Partitions("rooms")
	if err != nil {
		log.Fatalf("Failed to get partitions: %v", err)
	}

	for _, partition := range partitions {

		// TOOD: Get parameters of a partition
		var white bool = false
		var black bool = false

		fmt.Printf("Room: (%d), White: (%t), Black: (%t) \n", partition, white, black)
	}

}
