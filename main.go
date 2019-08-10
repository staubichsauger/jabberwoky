package main

import (
	"flag"
	"github.com/staubichsauger/jabberwoky/bot"
	"log"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	host := flag.String("host", "localhost", "uno server host")
	port := flag.Int("port", 3000, "uno server port")
	numPlayers := flag.Int("players", 2, "number of clients to spawn")
	flag.Parse()

	url, err := url.Parse("http://" + *host + ":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatal("Invalid hostname and port supplied: ", err)
	}

	var players []*bot.Client

	for i:=0; i < *numPlayers; i++ {
		players = append(players, &bot.Client{
			Url: *url,
		})
	}

	stop := make(chan error)

	for _, p := range players {
		if err := p.Login(); err != nil {
			log.Fatal(err)
		}
		go p.Play(stop)
	}


	for err := range stop {
		log.Print(err)
		if strings.Contains(err.Error(), "Status: 500") {
			return
		}
	}
}
