/*
Copyright © 2022 NAME HERE <ilinux@88.com>

*/
package internal

import (
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

var (
	wg sync.WaitGroup
)

func Conmuse(topic string, server string) {
	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_1
	var err error
	consumer, err := sarama.NewConsumer([]string{server}, config)
	if err != nil {
		fmt.Println("Failed to start consumer: ", err)
		return
	}
	//设置分区
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		fmt.Println("Failed to get the list of partitions: ", err)
		return
	}

	//循环分区
	for partition := range partitionList {
		//获取消费耗时
		start := time.Now()
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		later := time.Since(start)
		if err != nil {
			fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s , 程序消费Kafka用时=[%s]\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value), later)
				fmt.Println()
			}

		}(pc)
	}
	//time.Sleep(time.Hour)
	wg.Wait()
	consumer.Close()
}
