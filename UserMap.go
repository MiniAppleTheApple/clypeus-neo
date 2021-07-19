package main

import "fmt"

type UserMap struct {
	users map[string]*User
}

var warningHandler *UserMap = &UserMap{make(map[string]*User)}

func GetUserMap() *UserMap {
	return warningHandler
}

func IdForUserMapCalculation(guildID string, authorID string) string {
	return fmt.Sprintf("%v %v", guildID, authorID)
}

func (handler *UserMap) StoreUser(id string, user *User) {
	handler.users[id] = user
}
func (handler *UserMap) GetUserById(id string) (*User, bool) {
	var val *User
	var ok bool

	val, ok = handler.users[id]

	return val, ok
}
