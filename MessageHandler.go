package main

import (
	"fmt"
	"strings"
	discord "github.com/bwmarrin/discordgo"
)

type MessageHandler struct{
	Commands []Command
}

func NewMessageHandler() MessageHandler{
	return MessageHandler{
		commands: []Command{
			Help{},
		},
	}
}

func (self *MessageHandler) handle(prefix string,bot *discord.Session,msg *discord.MessageCreate){
	splited := strings.Split(msg.Content," ")
	command_name := splited[0]

	for i := 0;i < len(self.Commands);i++{
		v := self.Commands[i]
		if name := v.GetCommandName();command_name == prefix + name {
			v.Handle(bot,msg)
			fmt.Println(prefix + name)
		}
	}
}


