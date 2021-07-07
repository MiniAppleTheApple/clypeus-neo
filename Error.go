package main

import "fmt"

func handle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}