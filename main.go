package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tg/bin/Application"
	"tg/bin/Tg"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("token")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	botStart(bot)
}

func botStart(bot *tgbotapi.BotAPI) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			message := Application.Apple(update.Message.Chat.ID, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, message.Message)
			if message.Keyboard {
				msg.ReplyMarkup = tg.MenuKeyboard
			}
			bot.Send(msg)

		}
	}
}
