package main

import (
	"github.com/staubichsauger/jabberwoky/app"
	"os"
)

func main() {
	player1 := app.App{
		Url: os.Args[1],
	}


	player1.Login()
	//player2.Login()

	go player1.Play()
	//go player2.Play()

	wait := make(chan bool)
	<-wait
}
