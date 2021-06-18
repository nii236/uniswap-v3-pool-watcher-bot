package main

import (
	"log"
	"os"
	utils "uniswap-v3-pool-watcher-bot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_KEY"))
	if err != nil {
		log.Panic(err)
	}
	// Uncomment below line for complete debug output
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.Text != "/status" { // ignore any non-Message Updates
			continue
		}

		updatedMsg := utils.HandleStatusCmd()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, updatedMsg)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
