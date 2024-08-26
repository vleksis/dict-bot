package main

import "dict-bot/telegram"

func main() {
	bot := telegram.GetBot()
	updates := telegram.GetUpdatesChannel(bot)

	for update := range updates {
		telegram.HandleInput(bot, update)
	}

}
