package main

import (
	// "fmt"
	"example.com/main/command"
	"strings"
	discord "github.com/bwmarrin/discordgo"
)

type MessageHandler struct{
	Command_Handler CommandHandler
}

func NewMessageHandler(bot Bot,commands []command.Command) MessageHandler{
	return MessageHandler{
		CommandHandler{
			commands,
			bot,
		},
	}
}

func (self *MessageHandler) Handle(msg *discord.MessageCreate){
	splited := strings.Split(msg.Content," ")
	command_name := splited[0]
	self.Command_Handler.Handle(command_name,msg)
}


