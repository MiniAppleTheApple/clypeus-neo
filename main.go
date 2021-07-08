package main

import (
	"os"
	"fmt"
	"syscall"
	"os/signal"
	"encoding/json"
	io "io/ioutil"
	. "example.com/main/command"
	// "example.com/main/command/data"
	discord "github.com/bwmarrin/discordgo"
)

type Data struct {
	Token string `json:"token"`
	DataManagerType string`json:"datamanger"`
}

func (self *Data) createBot(prefix string) Bot{
	_bot, err := discord.New("Bot " + self.Token)
	
	if err != nil {
		fmt.Println(err)
	}
	bot := Bot{_bot,prefix}
	handler := NewMessageHandler(bot,[]Command{Help{}})
	bot.Bot.AddHandler(func (s *discord.Session, m *discord.MessageCreate) {
		handler.Handle(m)
	})
	bot.Bot.Identify.Intents = discord.IntentsAll
	return bot
}

type Bot struct {
	Bot *discord.Session
	Prefix string
}
func handle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	file,err := io.ReadFile("settings.json")
	if err != nil {
		handle(err)
		file,err = io.ReadFile("settings.example.json")
		if err != nil {
			fmt.Println("pls check if there is settings.example.json or settings.json in your disk")
		}
	}

	data := Data{}

	err = json.Unmarshal(file,&data)
	handle(err)

	myBot := data.createBot("!")
	bot := myBot.Bot
	fmt.Println("Bot is running!")

	err = bot.Open()
	handle(err)

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
	return
}