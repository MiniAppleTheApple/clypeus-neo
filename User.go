package main

import (
	discord "github.com/bwmarrin/discordgo"
)

type User struct {
	at_cool_down  bool
	message_count int32
	warn_count    int32
	User          *discord.User
}

func NewUser(user *discord.User) *User {
	return &User{
		false,
		0,
		0,
		user,
	}
}

func (user *User) IsAtCoolDown() bool {
	return user.at_cool_down
}

func (user *User) CoolDown() {
	user.at_cool_down = true
}

func (user *User) ResetCoolDown() {
	user.at_cool_down = false
}

func (user *User) GetMessageCount() int32 {
	return user.message_count
}

func (user *User) AddMessageCount() {
	user.message_count += 1
}

func (user *User) ResetMessageCount() {
	user.message_count = 0
}

func (user *User) GetWarnCount() int32 {
	return user.warn_count
}

func (user *User) Warn() {
	user.warn_count += 1
}

func (user *User) IsWarned() bool {
	return user.warn_count > 0
}
