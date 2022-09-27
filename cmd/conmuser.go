/*
Copyright © 2022 NAME HERE <ilinux@88.com>

*/
package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"

	"github.com/spf13/cobra"
)

// conmuserCmd represents the conmuser command
var conmuserCmd = &cobra.Command{
	Use:   "conmuser",
	Short: "Conmuser messages to a topic",
	Long:  `📖使用这个子命令可以帮你在topic中消费消息`,

	Run: func(cmd *cobra.Command, args []string) {
		server, _ = cmd.Flags().GetString("server")
		topic, _ = cmd.Flags().GetString("topic")
		Conmuse(topic, server)

	},
}

func init() {
	rootCmd.AddCommand(conmuserCmd)
	conmuserCmd.Flags().StringP("server", "s", "192.168.1.1:9092", "🔗连接你的Kafka Server,可以是很多台，也可以是一台 : 192.168.1.1:9092")
	conmuserCmd.Flags().StringP("topic", "t", "test", "🈯️指定你需要测试的Topic名称 : test")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// conmuserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// conmuserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
