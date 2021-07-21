package command

import (
	discord "github.com/bwmarrin/discordgo"
)

type Command interface {
	ToArguments([]string) error
	Handle(*discord.Session, *discord.MessageCreate) error
}
