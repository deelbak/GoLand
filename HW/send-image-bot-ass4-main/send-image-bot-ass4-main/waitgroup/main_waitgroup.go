package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/deelbak/send-image-bot-ass4/keys"
	"github.com/deelbak/send-image-bot-ass4/random"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func incrementCounter(c *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*c++

}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	bot, err := bot.New(keys.Token(), bot.WithDefaultHandler(handler))
	if err != nil {
		log.Fatal(err)
	}

	bot.Start(ctx)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	var c int
	var wg sync.WaitGroup

	if update.Message.Text == "image" || update.Message.Text == "/image" {
		// for i := 0; i < 100; i++ {
			wg.Add(1)
			go incrementCounter(&c, &wg)
			go func(chatID int64) {
				defer wg.Done()

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
		// }
	}

	wg.Wait()
	fmt.Println("Counter:", c)
}
