/*
Copyright © 2022 NAME HERE <ilinux@88.com>

*/
package internal

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

// 基于sarama第三方库开发的kafka client

func Produce(topic string, server string, number int) (err error) {
	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_1
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Key = sarama.StringEncoder("hello")
	msg.Value = sarama.StringEncoder("vlinux")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{server}, config)

	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	// 发送消息
	for i := 0; i < number; i++ {
		start := time.Now()
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			panic(err)
		}
		later := time.Since(start)
		fmt.Printf("Pid:%v Offset:%v,程序发送Kafka用时=[%s]\n", pid, offset, later)
	}
	return
}
