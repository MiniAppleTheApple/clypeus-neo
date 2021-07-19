package main

import (
	"fmt"
	"time"

	discord "github.com/bwmarrin/discordgo"
)

type CoolDownCounter struct{}

const (
	TimeToReset = time.Millisecond * 1500
)

func NewCoolDownCounter() *CoolDownCounter {
	return &CoolDownCounter{}
}
func (handler *CoolDownCounter) Handle(msg *discord.MessageCreate) bool {
	var author string = msg.Author.ID
	var guild string = msg.GuildID

	var id string = IdForUserMapCalculation(author, guild)
	var userMap *UserMap = GetUserMap()

	var _user *User
	var ok bool

	_user, ok = userMap.GetUserById(id)

	if !ok || _user.IsAtCoolDown() {
		fmt.Println("user is in cooldown")
		return true
	}
	fmt.Println("user is not in cooldown")
	var user *User = NewUser(msg.Author)
	user.CoolDown()
	go func() {
		time.Sleep(TimeToReset)
		user.ResetCoolDown()
	}()
	userMap.StoreUser(id, user)
	return false
}
