/*
Copyright © 2022 NAME HERE <ilinux@88.com>

*/
package cmd

import (
	"fmt"
	"kafkacli/internal"

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
		if err := internal.Produce(topic, server, number); err != nil {
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
