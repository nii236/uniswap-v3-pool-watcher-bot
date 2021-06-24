package main

import (
	"log"
	"os"
	"strconv"
	utils "uniswap-v3-pool-watcher-bot/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/urfave/cli/v2"
)

func main() {
	var bot_key, geth_url, timeout string
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "botkey",
				Aliases:     []string{"k"},
				Value:       "BOT_KEY",
				Destination: &bot_key,
				Usage:       "Bot key for the telegram bot",
			},
			&cli.StringFlag{
				Name:        "timeout",
				Aliases:     []string{"t"},
				Value:       "TIMEOUT",
				Destination: &timeout,
				Usage:       "timeout for the telegram bot",
			},
			&cli.StringFlag{
				Name:        "url",
				Aliases:     []string{"u"},
				Value:       "GETH_URL",
				Destination: &geth_url,
				Usage:       "Geth Url for the telegram bot",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	bot, err := tgbotapi.NewBotAPI(bot_key)
	if err != nil {
		log.Println("Bot creation error:", err)
		return
	}
	// Uncomment below line for complete debug output
	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout, err = strconv.Atoi(timeout)
	if err != nil {
		log.Println("Error converting timeout: ", err)
		return
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Println("Error getting channel updates:", err)
		return
	}
	for update := range updates {
		if update.Message == nil || update.Message.Text != "/status" { // ignore any non-Message Updates
			continue
		}
		updated_msg := utils.HandleStatusCmd(geth_url)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, updated_msg)
		msg.ReplyToMessageID = update.Message.MessageID

		_, err := bot.Send(msg)
		if err != nil {
			log.Println("Bot send error: ", err)
		}
	}
}
