package app

import (
	"bytes"
	"encoding/json"
	"github.com/staubichsauger/jabberwoky/structs"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type App struct {
	Url string
	Id string
}

func (a *App) Login() {
	me := structs.Join{
		Name: "jabberwoky",
	}
	reqBytes, err := json.Marshal(&me)
	if err != nil {
		log.Print(err)
	}

	res, err := http.Post(a.Url + "/join", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Print(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Print(res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
	}

	id := structs.Id{}
	err = json.Unmarshal(body, &id)
	if err != nil {
		log.Print(err)
	}

	a.Id = id.PlayerId
}

func (a *App) Play() {
	for _ = range time.Tick(time.Millisecond * 20) {
		res, err := http.Get(a.Url + "/games?id=" + a.Id)
		if err != nil {
			log.Print(err)
		}
		if res.StatusCode != http.StatusOK {
			log.Print("Error getting games endpoint: " + strconv.Itoa(res.StatusCode))
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Print(err)
		}

		gs := structs.GameStatus{}
		err = json.Unmarshal(body, &gs)
		if err != nil {
			log.Print(err)
		}

		if gs.MyTurn {
			turn := gs.DoTurn()
			bj, err := json.Marshal(&turn)
			if err != nil {
				log.Print(err)
			}
			res, err := http.Post(a.Url + "/games?id=" + a.Id, "application/json", bytes.NewBuffer(bj))
			if err != nil {
				log.Print(err)
			}
			if res.StatusCode != http.StatusOK {
				//log.Print(*turn.PlayCard)
				body, _ := ioutil.ReadAll(res.Body)
				log.Print("Error posting turn: " + strconv.Itoa(res.StatusCode) + "-> " + string(body))
			}
		}
	}
}