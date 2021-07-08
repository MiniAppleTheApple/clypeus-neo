package main
import (
	"fmt"
	"example.com/main/command"
	discord "github.com/bwmarrin/discordgo"
)
type CommandHandler struct {
	commands []command.Command
	bot Bot
}
func NewCommandHandler() CommandHandler {
	return CommandHandler {}
}
func (self CommandHandler) Handle(command_name string,msg *discord.MessageCreate){
	for i := 0;i < len(self.commands);i++{
		v := self.commands[i]
		if name := v.GetCommandName();command_name == self.bot.Prefix + name {
			v.Handle(self.bot.Bot,msg)
			fmt.Println(self.bot.Prefix + name)
		}
	}
}