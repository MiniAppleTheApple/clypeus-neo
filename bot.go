package main

import (
	discord "github.com/bwmarrin/discordgo"
)

var bot *discord.Session

func GetBot() *discord.Session {
	return bot
}

func SetBot(_bot *discord.Session) {
	bot = _bot
}
