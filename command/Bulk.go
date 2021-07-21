package command

import (
	// "example.com/main/command/tool"

	"fmt"
	"strconv"

	discord "github.com/bwmarrin/discordgo"
)

type Bulk struct {
	times int
}

func AddBulk() *Bulk {
	return &Bulk{100}
}

func (bulk Bulk) Handle(bot *discord.Session, msg *discord.MessageCreate) error {
	fmt.Printf("BULK: times = %v\n", bulk.times)
	messages_id := []string{}
	messages, err := bot.ChannelMessages(msg.ChannelID, bulk.times, "", "", "")
	if err != nil {
		return err
	}
	for i := range messages {
		messages_id = append(messages_id, messages[i].ID)
	}
	err = bot.ChannelMessagesBulkDelete(msg.ChannelID, messages_id)
	return err
}

func (bulk *Bulk) ToArguments(args []string) error {
	var err error
	if len(args) > 1 {
		bulk.times, err = strconv.Atoi(args[1])
		fmt.Println(bulk.times)
		return err
	}
	bulk.times = 100
	return err
}
