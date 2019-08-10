package bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/staubichsauger/jabberwoky/game"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	Url url.URL
	Id string
	Stop chan error
}

func (a *Client) Login() error {
	me := game.Join{
		Name: "jabberwoky",
	}
	reqBytes, err := json.Marshal(&me)
	if err != nil {
		return err
	}

	res, err := http.Post(a.Url.String() + "/join", "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("Got status: " + strconv.Itoa(res.StatusCode))
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	id := game.Id{}
	err = json.Unmarshal(body, &id)
	if err != nil {
		return err
	}

	a.Id = id.PlayerId
	return nil
}

func (a *Client) Play(stop chan error) {
	a.Stop = stop
	for _ = range time.Tick(time.Millisecond * 40) {
		res, err := http.Get(a.Url.String() + "/games?id=" + a.Id)
		if err != nil {
			stop <- err
			return
		}
		if res.StatusCode != http.StatusOK {
			stop <- errors.New("Error getting games endpoint, Status: " + strconv.Itoa(res.StatusCode))
			continue
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			stop <- err
			return
		}

		gs := game.Status{}
		err = json.Unmarshal(body, &gs)
		if err != nil {
			stop <- err
			return
		}

		if gs.MyTurn {
			turn := gs.DoTurn()
			bj, err := json.Marshal(&turn)
			if err != nil {
				stop <- err
				return
			}
			res, err := http.Post(a.Url.String() + "/games?id=" + a.Id, "application/json", bytes.NewBuffer(bj))
			if err != nil {
				stop <- err
				return
			}
			if res.StatusCode != http.StatusOK {
				//log.Print(*turn.PlayCard)
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					stop <- err
					return
				}
				stop <- errors.New("Error posting turn: " + strconv.Itoa(res.StatusCode) + "-> " + string(body))
			}
		}
	}
}