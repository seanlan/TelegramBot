/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	tele "gopkg.in/telebot.v3"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := tele.NewBot(tele.Settings{
			Token:   "6935398271:AAF3tkwnlyHeuPRSVWGPVcWH-GXYBfNlyno",
			Offline: true,
		})
		if err != nil {
			panic(err)
		}
		_, err = b.Send(&tele.User{
			ID: -4253239938,
		}, "hello world")
		if err != nil {
			zap.S().Infof("send message error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
