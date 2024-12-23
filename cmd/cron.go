/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"TelegramBot/internal/crontab"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"log"

	"github.com/spf13/cobra"
)

type CronZapLogger struct {
}

func (*CronZapLogger) Info(msg string, keysAndValues ...interface{}) {
	zap.S().Debug(msg, keysAndValues)
}

func (*CronZapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	zap.S().Error(err, msg, keysAndValues)
}

func cronFunc(cmd *cobra.Command, args []string) {
	var err error
	c := cron.New(cron.WithLogger(&CronZapLogger{}), cron.WithChain(cron.Recover(&CronZapLogger{})))
	// 每秒检查一次 推送公告
	_, err = c.AddFunc("@every 1s", crontab.CronPushGroupMessage)
	if err != nil {
		log.Fatalf("cron add func error: %#v", err)
	}
	// 启动定时任务
	c.Start()
	select {}
}

// cronCmd represents the cron command
var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: cronFunc,
}

func init() {
	rootCmd.AddCommand(cronCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cronCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cronCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
