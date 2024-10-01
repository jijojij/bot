package tg

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var MenuKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("add"),
		tgbotapi.NewKeyboardButton("read all"),
	),
)
