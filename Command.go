package main
import (
	discord "github.com/bwmarrin/discordgo"
)
type Command interface {
	GetCommandName() string
	Handle(*discord.Session,*discord.MessageCreate)
}