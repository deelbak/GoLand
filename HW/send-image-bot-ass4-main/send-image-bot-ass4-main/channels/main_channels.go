package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/deelbak/send-image-bot-ass4/keys"
	"github.com/deelbak/send-image-bot-ass4/random"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func incrementCounter(c chan int) {
	c <- 1
}
func main() {
	c := make(chan int)
	// for i := 0; i < 100; i++ {

		ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
		defer cancel()

		bot, err := bot.New(keys.Token(), bot.WithDefaultHandler(handler))
		if err != nil {
			log.Fatal(err)
		}
		go incrementCounter(c)
		bot.Start(ctx)
	// }
	var sum int
	for i := 0; i < 100; i++ {
		sum += <-c
	}
	fmt.Println("Sum:", c)

}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	message := make(chan bool, 1)

	if update.Message.Text == "image" || update.Message.Text == "/image" {
		message <- true

		go func(chatID int64) {
			defer func() { <-message }()

			photo, err := random.RandomPhoto(keys.AccessKey())
			if err != nil {
				log.Fatal(err)
			}

			params := &bot.SendPhotoParams{
				ChatID: update.Message.Chat.ID,
				Photo:  &models.InputFileString{Data: photo.URLs.Regular},
			}
			if _, err := b.SendPhoto(ctx, params); err != nil {
				log.Println(err)
			}
		}(update.Message.Chat.ID)

	}

}
