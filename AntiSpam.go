package main

import (
	"fmt"
	"time"

	discord "github.com/bwmarrin/discordgo"
)

// handler which handles SPAM and others problem
type AntiSpam struct{}

const (
	// time that handler reset the count of message
	timeToReset = time.Millisecond * 3000
	// the max message that the user can send at {TimeToRest},
	// if the user send more than the max message,
	// then the user get warned and can't send message
	maxMessageCount = 5
)

// DangerHandler constructor
func NewAntiSpam() *AntiSpam {
	return &AntiSpam{}
}

// handling problem
func (_antispam AntiSpam) Handle(msg *discord.MessageCreate) {
	var (
		warningHandler *UserMap         = GetUserMap()
		bot            *discord.Session = GetBot()

		author *discord.User = msg.Author
		id     string        = IdForUserMapCalculation(author.ID, msg.GuildID)

		user *User
		ok   bool
	)

	user, ok = warningHandler.GetUserById(id)

	if !ok {
		var user *User = NewUser(author)

		warningHandler.StoreUser(id, user)
		// fmt.Printf("user storaged in the map %v", user)
		return
	}

	user.AddMessageCount()

	if user.IsWarned() {
		// fmt.Printf("%v is warned\n", msg.Author.ID)
		GetBot().ChannelMessageDelete(msg.ChannelID, msg.ID)
		return
	}

	if user.GetMessageCount() == 1 {
		go func() {
			time.Sleep(timeToReset)

			var count int32 = user.GetMessageCount()

			// fmt.Printf("reseted with %v\n", count)
			if count > maxMessageCount-1 {
				user.Warn()

				var messages_id []string

				var messages []*discord.Message
				var err error

				messages, err = bot.ChannelMessages(msg.ChannelID, int(count), "", "", "")

				if err != nil {
					return
				}

				for i := range messages {
					messages_id = append(messages_id, messages[i].ID)
				}

				bot.ChannelMessagesBulkDelete(msg.ChannelID, messages_id)
				bot.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("%v you are warned", author.Mention()))
			}
			user.ResetMessageCount()
		}()
		return
	}
	// fmt.Println("user exist in data")
}
