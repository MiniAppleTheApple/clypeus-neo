package main
import (
	"fmt"
	"time"
	discord "github.com/bwmarrin/discordgo"
)
type DangerHandler struct {
	bot *discord.Session
	users map[string]*User
}

func (self DangerHandler) Handle(msg *discord.MessageCreate){
	user := msg.Author
	val,ok := self.users[user.ID]

	if ok {
		val.AddMessageCount(1)

		if val.IsWarned(){
			fmt.Printf("%v is warned\n",msg.Author.ID)
			self.bot.ChannelMessageDelete(msg.ChannelID,msg.ID)
			return
		}
		if val.GetMessageCount() == 1{
			go func() {
				time.Sleep(time.Millisecond * 3000)
				count := val.GetMessageCount()
				fmt.Printf("reseted with %v\n",count)
				if count > 5{
					val.Warn()
					messages_id := []string{}
					messages,err := self.bot.ChannelMessages(msg.ChannelID,int(count),"","","")
					if err != nil {
						return
					}
					for i := range messages {
						messages_id = append(messages_id,messages[i].ID)
					}
					self.bot.ChannelMessagesBulkDelete(msg.ChannelID,messages_id)
					self.bot.ChannelMessageSend(msg.ChannelID,fmt.Sprintf("%v you are warned",user.Mention()))
				}
				val.ResetMessageCount()
			}()
			return
		}
		fmt.Println("user exist in data")
	}else{
		self.users[user.ID] = NewUser(user)
		fmt.Printf("user storaged in the map %v",val)
	}
}
