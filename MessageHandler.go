package main

import (
	// "fmt"
	"strings"

	"example.com/main/command"
	discord "github.com/bwmarrin/discordgo"
)

type MessageHandler struct {
	commander *Commander
	antispam  *AntiSpam
}

func NewMessageHandler(commands map[string]command.Command) MessageHandler {
	return MessageHandler{
		NewCommander(commands),
		NewAntiSpam(),
	}
}

func (handler *MessageHandler) Handle(msg *discord.MessageCreate) {
	if msg.Author.ID == GetBot().State.User.ID {
		return
	}

	handler.antispam.Handle(msg)

	var content string = msg.Content

	if len(GetSettings().Prefix) < len(content) {
		splited := strings.Split(content[len(GetSettings().Prefix):], " ")
		handler.commander.Handle(splited, msg)
		return
	}
}
