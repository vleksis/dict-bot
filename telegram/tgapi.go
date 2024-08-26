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
	if upd.EditedMessage != nil {
		go handleEditedMessage(bot, upd.EditedMessage)
	}
	if upd.Message != nil {
		if upd.Message.IsCommand() {
			go handleCommand(bot, upd.Message)
		} else {
			go handleMessage(bot, upd.Message)
		}
	}
}

func handleEditedMessage(bot *tgbotapi.BotAPI, eMessage *tgbotapi.Message) {
	chatID := eMessage.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Editing is useless")
	msg.ReplyToMessageID = eMessage.MessageID
	bot.Send(msg)
}

func handleCommand(bot *tgbotapi.BotAPI, cMessage *tgbotapi.Message) {
	cmd := cMessage.Command()
	args := cMessage.CommandArguments()

	for _, possibleCommand := range availableCommands {
		if possibleCommand.name == cmd {
			go func() {
				msg := tgbotapi.NewMessage(cMessage.Chat.ID,
					possibleCommand.handler(args),
				)
				bot.Send(msg)
			}()
			return
		}
	}

	msg := tgbotapi.NewMessage(cMessage.Chat.ID, "invalid command")
	bot.Send(msg)
}

// do the /lookup command
func handleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	go func() {
		msg := tgbotapi.NewMessage(message.Chat.ID,
			availableCommands[0].handler(message.Text),
		)
		bot.Send(msg)
	}()
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
