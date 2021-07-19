package main

type Command interface {
	GetCommandName() string
	Handle([]string)
}