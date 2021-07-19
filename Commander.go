package main

import (
	"fmt"

	"example.com/main/command"
	discord "github.com/bwmarrin/discordgo"
)

type Commander struct {
	commands []command.Command // usabble command
}

// CommandHandler constructor
func NewCommander(commands []command.Command) *Commander {
	return &Commander{
		commands,
	}
}

// handling content as command
func (handler Commander) Handle(args []string, msg *discord.MessageCreate) {
	command_name := args[0]
	fmt.Println(command_name)
	var bot *discord.Session = GetBot()
	for i := 0; i < len(handler.commands); i++ {
		v := handler.commands[i]
		if name := v.GetCommandName(); command_name == name {
			var coolDownCounter *CoolDownCounter = NewCoolDownCounter()
			if coolDownCounter.Handle(msg) {
				bot.ChannelMessageSend(msg.ChannelID, "at cooldown")
				return
			}

			// fmt.Println(name)
			err := v.ToArguments(args)
			if err != nil {
				bot.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("%v", err))
			}
			err = v.Handle(GetBot(), msg)
			if err != nil {
				bot.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("%v", err))
			}
			fmt.Println(GetSettings().Prefix + name)
		}
	}
}
