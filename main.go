package main

import (
	"context"
	"github.com/shamu00/chatgpt-telegram/ping"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shamu00/chatgpt-telegram/src"
	"github.com/shamu00/chatgpt-telegram/src/tgbot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	libconfig "github.com/shamu00/chatgpt-telegram/src/config"
)

func mustInit() {
	rand.Seed(time.Now().UnixNano())
	libconfig.InitConfigurationFetcher()
	return
}

func main() {
	mustInit()
	ctx := src.PrepareContext()
	ping.StartPingServer()
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		ctx.Bot.StopReceivingUpdates()
		log.Println("Exit")
		ping.StopPingServer(context.Background())
		os.Exit(0)
	}()

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := ctx.Bot.GetUpdatesChan(updateConfig)

	log.Println("start handling messages...")
	tgbot.HandleBotMessage(ctx, updates)

}
