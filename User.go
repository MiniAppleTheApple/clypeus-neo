package main
import (
	discord "github.com/bwmarrin/discordgo"
)
type User struct{
	message_count int32
	warned bool
	User *discord.User
}
func NewUser(user *discord.User) *User{
	return &User{
		0,
		false,
		user,
	}
}
func (self *User) GetMessageCount() int32{
	return self.message_count
}
func (self *User) AddMessageCount(value int32){
	self.message_count += value
}

func (self *User) ResetMessageCount(){
	self.message_count = 0
}
func (self *User) Warn(){
	self.warned = true
}

func (self *User) IsWarned() bool {
	return self.warned
}