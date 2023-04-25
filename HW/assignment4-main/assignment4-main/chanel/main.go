package main

import (
	"encoding/json"
	"log"
	"net/http"

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

	chanel := make(chan int)
	count := 1

	go func() {
		for update := range updates {
			if update.Message != nil {
				if update.Message.IsCommand() && update.Message.Command() == "image" || update.Message.Text == "image" {
					chanel <- count
					count++
					log.Println("check")
					photo, err := GetRandomPhoto()
					log.Println("check2")
					if err != nil {
						log.Println(err)
					}
					log.Println("check3")
					file := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(photo))
					log.Println("check4")
					bot.Send(file)
					log.Println("check5")

				}
			}
		}
	}()

	go func() {
		for {
			log.Println("cnt")
			log.Printf("Count: %v", <-chanel)
		}
	}()

	select {}
}
