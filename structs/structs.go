package structs

type Join struct {
	Name string `json:"name"`
}

type GameStatus struct {
	MyTurn        bool     `json:"my_turn"`
	Hand          []Card   `json:"hand"`
	OtherPlayers  []Player `json:"other_players"`
	DiscardedCard Card     `json:"discarded_card"`
}

func (gs *GameStatus) DoTurn() (turn Turn) {
	var possible []Card
	red := 0
	blue := 0
	green := 0
	yellow := 0
	for _, c := range gs.Hand {
		switch c.Color {
		case gs.DiscardedCard.Color:
			possible = append(possible, c)
		case "":
			possible = append(possible, c)
		}
		if c.Value == gs.DiscardedCard.Value && numbers[gs.DiscardedCard.Value] > 0 && numbers[gs.DiscardedCard.Value] < 10 {
			possible = append(possible, c)
		}
		if c.Color != "" {
			switch c.Color {
			case "RED":
				red++
			case "BLUE":
				blue++
			case "GREEN":
				green++
			case "YELLOW":
				yellow++
			}
		}
	}

	turn = Turn{PlayCard:nil}
	if len(possible) == 0 {
		return
	}
	turn.PlayCard = &possible[0]
	for _, c := range possible {
		if numbers[c.Value] > numbers[turn.PlayCard.Value] {
			turn.PlayCard = &c
		}
	}

	if numbers[turn.PlayCard.Value] < -1 {
		var c string
		if red >= yellow && red>=green && red>=blue {
			c = "RED"
		} else if yellow >= green && yellow >= blue {
			c = "YELLOW"
		} else if green >= blue {
			c = "GREEN"
		} else {
			c = "BLUE"
		}

		turn.PlayCard.Color = c
	}

	return
}

type Card struct {
	Color string `json:"color"`
	Value string `json:"value"`
}

type Player struct {
	Name string `json:"player_name"`
	CardCount int `json:"card_count"`
}

type Turn struct {
	PlayCard *Card `json:"play_card"`
}

type Id struct {
	PlayerId string `json:"player_id"`
	PlayerName string `json:"player_name"`
}

var numbers = map[string]int {
	"ONE": 1,
	"TWO": 2,
	"THREE": 3,
	"FOUR": 4,
	"FIVE": 5,
	"SIX": 6,
	"SEVEN": 7,
	"EIGHT": 8,
	"NINE": 9,
	"REVERSE": 0,
	"DRAW_TWO": 10,
	"SKIP": 0,
	"WILD": -2,
	"WILD_DRAW_FOUR": -3,
}