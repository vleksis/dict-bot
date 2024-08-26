package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var (
	apiKey = os.Getenv("TELEGRAM_API_KEY")
)

func GetBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		log.Panic(err)
	}

	// TODO: delete debug
	bot.Debug = true
	makeMenu(bot)

	return bot
}

func GetUpdatesChannel(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	return bot.GetUpdatesChan(u)
}

// Take user's input and dispatch it
func HandleInput(bot *tgbotapi.BotAPI, upd tgbotapi.Update) {
	if upd.Message != nil {
		if upd.Message.IsCommand() {
			go handleCommand(bot, upd)
		} else {
			go handleMessage(bot, upd)
		}
	}
}

func handleCommand(bot *tgbotapi.BotAPI, upd tgbotapi.Update) {
	cmd := upd.Message.Command()
	args := upd.Message.CommandArguments()

	for _, possibleCommand := range availableCommands {
		if possibleCommand.name == cmd {
			go func() {
				msg := tgbotapi.NewMessage(upd.Message.Chat.ID,
					possibleCommand.handler(args),
				)
				bot.Send(msg)
			}()
			return
		}
	}

	msg := tgbotapi.NewMessage(upd.Message.Chat.ID, "invalid command")
	bot.Send(msg)
}

// pre-condition: upd.Message != nil
func handleMessage(bot *tgbotapi.BotAPI, upd tgbotapi.Update) {
	chatID := upd.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "todo something")
	msg.ReplyToMessageID = upd.Message.MessageID
	bot.Send(msg)
}

func makeMenu(bot *tgbotapi.BotAPI) {
	config := tgbotapi.NewSetMyCommands()

	for _, cmd := range availableCommands {
		config.Commands = append(config.Commands,
			tgbotapi.BotCommand{
				Command:     cmd.name,
				Description: cmd.shortDescription,
			},
		)
	}

	bot.Send(config)
}
