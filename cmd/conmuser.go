/*
Copyright © 2022 NAME HERE <ilinux@88.com>

*/
package cmd

import (
	"kafkacli/internal"

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
		internal.Conmuse(topic, server)

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
