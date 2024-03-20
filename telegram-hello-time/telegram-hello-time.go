package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := os.Getenv("TELEGRAM_TIME_HELLO_BOTAPI_KEY")
	log.Printf("Starting bot with token %s", token)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Print("The bot is available at t.me/TimeHelloBot")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	// interval started for user
	intervalStarted := make(map[int64]chan struct{})

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			switch update.Message.Command() {
			case "close":
				id := update.Message.Chat.ID
				if intervalStarted, ok := intervalStarted[id]; ok {
					log.Printf("Chat %d cancelling previous interval", id)
					msg := tgbotapi.NewMessage(id, "Cancelling previous interval.")
					bot.Send(msg)
					close(intervalStarted)
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "No interval started for this chat.")
					bot.Send(msg)
				}
			case "hello":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Hello! Current time is: %s", time.Now().Format("2006-01-02 15:04:05")))
				bot.Send(msg)
			case "interval":
				if len(update.Message.CommandArguments()) > 0 {
					interval, err := strconv.Atoi(update.Message.CommandArguments())
					if err != nil {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please specify a valid number of seconds.")
						bot.Send(msg)
						continue
					}
					id := update.Message.Chat.ID
					if intervalStarted, ok := intervalStarted[id]; ok {
						log.Printf("Chat %d cancelling previous interval", id)
						msg := tgbotapi.NewMessage(id, "Cancelling previous interval.")
						bot.Send(msg)
						close(intervalStarted)
					}
					ch := make(chan struct{})
					intervalStarted[id] = ch
					log.Printf("Chat %d started interval for %d seconds", id, interval)

					go func(interval int, ch chan struct{}) {
						for {
							select {
							case <-ch:
								return
							case <-time.After(time.Duration(interval) * time.Second):
								log.Printf("Sending message to chat %d interval %d", id, interval)
								msg := tgbotapi.NewMessage(id, fmt.Sprintf("Current time is: %s", time.Now().Format("2006-01-02 15:04:05")))
								bot.Send(msg)
							}
						}
					}(interval, ch)
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Please specify the interval in seconds.")
					bot.Send(msg)
				}
			}
		}
	}
}
