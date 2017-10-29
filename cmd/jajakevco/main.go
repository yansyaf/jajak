package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/toshim45/jajak/config"

	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
)

var (
	consumerGrupID string = "jajak-poll-update-evco"
)

func main() {
	envConfig := config.NewEnv()
	topics := []string{envConfig.EventTopic}
	zookeeperNodes := []string{envConfig.ZookeeperHost}
	log.Println("starting evco")
	evtConfig := consumergroup.NewConfig()
	evtConfig.Offsets.Initial = sarama.OffsetNewest
	evtConfig.Offsets.ProcessingTimeout = 5 * time.Second

	consumer, err := consumergroup.JoinConsumerGroup(consumerGrupID, topics, zookeeperNodes, evtConfig)
	if err != nil {
		log.Printf("fail to join consumer group %v\n", err)
	}

	// set terminate service
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if err := consumer.Close(); err != nil {
			sarama.Logger.Println("Error closing the consumer", err)
		}
	}()

	go func() {
		for err := range consumer.Errors() {
			log.Println(err)
		}
	}()

	// get event message
	eventCount := 0
	offsets := make(map[string]map[int32]int64)
	for message := range consumer.Messages() {
		if offsets[message.Topic] == nil {
			offsets[message.Topic] = make(map[int32]int64)
		}

		eventCount += 1
		if offsets[message.Topic][message.Partition] != 0 && offsets[message.Topic][message.Partition] != message.Offset-1 {
			log.Printf("Unexpected offset on %s:%d. Expected %d, found %d, diff %d.\n", message.Topic, message.Partition, offsets[message.Topic][message.Partition]+1, message.Offset, message.Offset-offsets[message.Topic][message.Partition]+1)
		}

		// Simulate processing
		log.Printf("Processing %v\n", string(message.Value))

		offsets[message.Topic][message.Partition] = message.Offset
		consumer.CommitUpto(message)
	}

	log.Printf("Processed %d events.", eventCount)
	log.Printf("%+v", offsets)
}
