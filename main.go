package main

import (
	"encoding/json"
	"fmt"
	io "io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"example.com/main/command"
	discord "github.com/bwmarrin/discordgo"
	// "example.com/main/command/data"
)

// function that handle error
func handle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	file, err := io.ReadFile("settings.json")
	if err != nil {
		handle(err)
		file, err = io.ReadFile("settings.example.json")
	}

	if err != nil {
		fmt.Println("pls check if there is settings.example.json or settings.json in your disk")
		return
	}

	err = json.Unmarshal(file, GetSettings())
	handle(err)

	bot, err := discord.New("Bot " + settings.Token)

	if err != nil {
		fmt.Println(err)
	}
	func() {
		handler := NewMessageHandler([]command.Command{command.AddHelp(), command.AddPurge(), command.AddBulk()})
		bot.AddHandler(func(s *discord.Session, m *discord.MessageCreate) {
			handler.Handle(m)
		})
	}()
	bot.Identify.Intents = discord.IntentsAll
	SetBot(bot)
	fmt.Println("Bot is running!")

	err = bot.Open()
	handle(err)

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
}
