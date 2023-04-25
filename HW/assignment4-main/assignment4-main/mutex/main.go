package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/deelbak/assignment4/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Photo struct {
	Urls Urls `json:"urls"`
}
type Urls struct {
	Regular string `json:"regular"`
}

func GetRandomPhoto() (string, error) {
	url := "https://api.unsplash.com/photos/random?client_id=" + config.AccessKey
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var photo Photo
	err = json.NewDecoder(resp.Body).Decode(&photo)
	if err != nil {
		return "", err
	}
	return photo.Urls.Regular, nil
}

func main() {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	mu := sync.Mutex{}
	count := 0
	go func() {
		mu.Lock()
		for update := range updates {
			if update.Message != nil {
				if update.Message.IsCommand() && update.Message.Command() == "image" || update.Message.Text == "image" {
					count++
					log.Printf("Count: %v", count)
					photo, err := GetRandomPhoto()
					if err != nil {
						log.Println(err)
					}
					file := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(photo))
					bot.Send(file)
				}
			}
		}
		mu.Unlock()
	}()
	select {}
}
