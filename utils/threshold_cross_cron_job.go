package utils

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron/v3"
)

// This function runs every 60 seconds to check if unclaimed fees exceeds $1000
func RunStartScheduler(bot *tgbotapi.BotAPI, gethUrl string, threshold float64) error {
	c := cron.New()
	// runs every 60 seconds
	_, err := c.AddFunc("@every 0h0m10s", func() {
		SendMessageOnThresholdBreach(bot, gethUrl, threshold)
	})
	if err != nil {
		return err
	}
	c.Start()
	return nil
}

// This function sends an alert once the threshold for unclaimed fees is breached (ex say >$1000)
func SendMessageOnThresholdBreach(bot *tgbotapi.BotAPI, gethUrl string, threshold float64) error {
	filteredPools := FilterThresholdCrossingPools(gethUrl, threshold)
	if len(filteredPools) > 0 {
		alertMsg := msgToSend(filteredPools, threshold)
		for _, chatID := range WhitelistTelegramAccountIDs {
			msg := tgbotapi.NewMessage(int64(chatID), alertMsg)
			_, err := bot.Send(msg)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// Format the message properly before sending to the bot
func msgToSend(arr []string, threshold float64) string {
	message := fmt.Sprintf("The following pools have exceeded $%f in total unclaimed fees:\n", threshold)
	for _, str := range arr {
		message += fmt.Sprintf("%s\n", str)
	}
	return message
}
