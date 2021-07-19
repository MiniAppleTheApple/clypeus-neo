package command

import (
	discord "github.com/bwmarrin/discordgo"
)

type Command interface {
	ToArguments([]string) error
	GetCommandName() string
	Handle(*discord.Session, *discord.MessageCreate) error
}
