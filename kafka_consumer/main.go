package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"kafka_consumer/dal"
	"kafka_consumer/handlers"
	"sync"
)

const Topic = "moni"

var wg sync.WaitGroup

func main() {
	defer dal.Close()
	dal.Init()

	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		fmt.Println("consumer connect err:", err)
		return
	}
	defer consumer.Close()

	//获取 kafka 主题
	partitions, err := consumer.Partitions(Topic)
	if err != nil {
		fmt.Println("get partitions failed, err:", err)
		return
	}

	fmt.Printf("get %v partitions\n", len(partitions))

	for _, p := range partitions {
		//sarama.OffsetNewest：从当前的偏移量开始消费，sarama.OffsetOldest：从最老的偏移量开始消费
		partitionConsumer, err := consumer.ConsumePartition(Topic, p, sarama.OffsetNewest)
		if err != nil {
			fmt.Println("partitionConsumer err:", err)
			continue
		}
		wg.Add(1)
		go func() {
			for m := range partitionConsumer.Messages() {
				// fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
				handlers.ItemChangeHandler(context.Background(), m.Value)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
