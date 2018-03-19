package bot

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"fmt"
)

type Robot struct {
	BotAPI      *tgbotapi.BotAPI
	CommandChan chan *tgbotapi.Message
	SendChan    chan tgbotapi.MessageConfig
}

func CreateRobot(config *Config) *Robot {
	robot := Robot{}
	robot.CommandChan = make(chan *tgbotapi.Message, 1)
	robot.SendChan = make(chan tgbotapi.MessageConfig, 1)

	api, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}

	robot.BotAPI = api

	api.Debug = true

	log.Printf("Authorized on account %s", api.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updater := func() {
		updates, err := api.GetUpdatesChan(u)

		if err != nil {
			log.Print("Can not listen an updates server")
		} else {
			for update := range updates {
				if update.Message.IsCommand() {
					robot.CommandChan <- update.Message
				}
			}
		}
	}

	commander := func() {
		for {
			select {
			case cmd := <-robot.CommandChan:
				switch cmd.Text {
				case "/hey":
					robot.SendChan <- tgbotapi.NewMessage(cmd.Chat.ID, "Hey, How you doing?")
				case "/migration":

				}
			case msg := <-robot.SendChan:
				api.Send(msg)
			}
		}
	}

	go commander()
	go updater()

	return &robot
}

func sendError(ch chan tgbotapi.MessageConfig, chatId int64, err error) {
	msg := fmt.Sprintf("Failed with an exception:\n%v", err.Error())
	ch <- tgbotapi.NewMessage(chatId, msg)
}

