package main
import (
	"fmt"
	"example.com/main/command"
	discord "github.com/bwmarrin/discordgo"
)
type CommandHandler struct {
	commands []command.Command
	bot *discord.Session
	data Data
}
func (self CommandHandler) Handle(args []string,msg *discord.MessageCreate){
	command_name := args[0]
	fmt.Println(command_name)
	for i := 0;i < len(self.commands);i++{
		v := self.commands[i]
		if name := v.GetCommandName();command_name == name {
			err := v.ToArguments(args)
			err = v.Handle(self.bot,msg)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(self.data.Prefix + name)
		}
	}
}