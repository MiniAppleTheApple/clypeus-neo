package command

import (
	"fmt"
	"strconv"

	discord "github.com/bwmarrin/discordgo"
)

type Purge struct {
	times int
}

func AddPurge() *Purge {
	return &Purge{100}
}
func (purge Purge) Handle(bot *discord.Session, msg *discord.MessageCreate) error {
	fmt.Printf("PURGE: times = %v\n", purge.times)
	messages, err := bot.ChannelMessages(msg.ChannelID, purge.times, "", "", "")
	for i := range messages {
		err := bot.ChannelMessageDelete(msg.ChannelID, messages[i].ID)
		if err != nil {
			return err
		}
	}
	if err != nil {
		return err
	}
	return err
}

func (self Purge) GetCommandName() string {
	return "purge"
}

func (self *Purge) ToArguments(args []string) error {
	var err error
	if len(args) > 1 {
		self.times, err = strconv.Atoi(args[1])
		fmt.Println(self.times)
		return err
	}
	self.times = 100
	return err
}
