package main

import (
	// "fmt"
	"example.com/main/command"
	"strings"
	discord "github.com/bwmarrin/discordgo"
)

type MessageHandler struct{
	data Data
	bot *discord.Session
	command CommandHandler
	danger DangerHandler
}

func NewMessageHandler(data Data,bot *discord.Session,commands []command.Command) MessageHandler{
	return MessageHandler{
		data,
		bot,
		CommandHandler{
			commands,
			bot,
			data,
		},
		DangerHandler{
			bot,
			make(map[string]*User),
		},
	}
}

func (self *MessageHandler) Handle(msg *discord.MessageCreate){
	if msg.Author.ID == self.bot.State.User.ID {
		return
	}

	self.danger.Handle(msg)

	content := msg.Content

	if len(self.data.Prefix) < len(content){
		splited := strings.Split(content[len(self.data.Prefix):]," ")
		self.command.Handle(splited,msg)
		return
	}
	return
}


