package main

import (
	"os"
	"fmt"
	"syscall"
	"os/signal"
	"encoding/json"
	io "io/ioutil"
	discord "github.com/bwmarrin/discordgo"
)

type Data struct {
	Token string `json:"token"`
	DataManagerType string`json:"datamanger"`
}

func (self *Data) createBot() *discord.Session{
	bot, err := discord.New("Bot " + self.Token)

	if err != nil {
		fmt.Println(err)
		return bot
	}
	return bot
}

func main() {
	file,err := io.ReadFile("settings.json")
	if err != nil {
		handle(err)
		file,err = io.ReadFile("settings.example.json")
	}

	data := Data{}

	err = json.Unmarshal(file,&data)
	handle(err)

	bot := data.createBot()

	bot.AddHandler(func (s *discord.Session, m *discord.MessageCreate) {
		handler := NewMessageHandler()
		handler.handle("!",s,m)
	})

	fmt.Println("Bot is running!")

	bot.Identify.Intents = discord.IntentsAll

	err = bot.Open()
	handle(err)

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
	return
}