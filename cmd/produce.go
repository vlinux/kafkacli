/*
Copyright © 2022 NAME HERE <ilinux@88.com>

*/
package cmd

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"

	"github.com/spf13/cobra"
)

var server string
var topic string
var number int

// produceCmd represents the produce command
var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "Produce messages to a topic",
	Long:  `📖使用这个子命令可以帮你在topic中生产消息`,
	//Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		number, _ = cmd.Flags().GetInt("number")
		server, _ = cmd.Flags().GetString("server")
		topic, _ = cmd.Flags().GetString("topic")
		if err := Produce(topic, server, number); err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(produceCmd)
	produceCmd.Flags().StringP("server", "s", "192.168.1.1:9092", "🔗连接你的Kafka Server,可以是很多台，也可以是一台 : 192.168.1.1:9092")
	produceCmd.Flags().StringP("topic", "t", "test", "🈯️指定你需要测试的Topic名称 : test")
	produceCmd.Flags().IntP("number", "n", 0, "🔢指定你需要生产几条消息到该Topic : 1")

}

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
