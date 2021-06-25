package utils

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron/v3"
)

// This function runs every 30 seconds to check if unclaimed fees exceeds $1000
func RunCronJob(bot *tgbotapi.BotAPI, gethUrl string) error {
	c := cron.New()
	// runs every 30 seconds
	_, err := c.AddFunc("@every 0h0m10s", func() {
		SendMessageOnThresholdBreach(bot, gethUrl)
	})
	if err != nil {
		return err
	}
	c.Start()
	return nil
}

func FilterThresholdCrossingPools()

// This function sends an alert once the threshold for unclaimed fees is breached (>$1000)
func SendMessageOnThresholdBreach(bot *tgbotapi.BotAPI, gethUrl string) error {
	for _, chatID := range WhitelistTelegramAccountIDs {
		// Write logic here
		log.Println(chatID)
	}
	return nil
}
