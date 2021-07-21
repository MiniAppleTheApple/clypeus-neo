package main

import (
	"fmt"

	"example.com/main/command"
	discord "github.com/bwmarrin/discordgo"
)

type Commander struct {
	commands map[string]command.Command // usabble command
}

// CommandHandler constructor
func NewCommander(commands map[string]command.Command) *Commander {
	return &Commander{
		commands,
	}
}

// handling content as command
func (commander Commander) Handle(args []string, msg *discord.MessageCreate) {
	command_name := args[0]
	fmt.Println(command_name)
	var bot *discord.Session = GetBot()

	var coolDownCounter *CoolDownCounter = NewCoolDownCounter()
	if coolDownCounter.Handle(msg) {
		bot.ChannelMessageSend(msg.ChannelID, "at cooldown")
		return
	}
	var command command.Command = commander.commands[command_name]
	// fmt.Println(name)
	err := command.ToArguments(args)
	if err != nil {
		bot.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("%v", err))
	}
	err = command.Handle(GetBot(), msg)
	if err != nil {
		bot.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("%v", err))
	}
	fmt.Println(GetSettings().Prefix + command_name)
}
