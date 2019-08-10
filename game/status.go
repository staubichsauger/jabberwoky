package game

type Status struct {
	MyTurn        bool     `json:"my_turn"`
	Hand          []Card   `json:"hand"`
	OtherPlayers  []Player `json:"other_players"`
	DiscardedCard Card     `json:"discarded_card"`
	possibleCol []Card
	possibleVal []Card
	numOfColors map[string]int
}

func (gs *Status) DoTurn() (turn Turn) {
	gs.analiseHand()

	turn = Turn{
		PlayCard: gs.getBestCard(),
	}

	return turn
}

func (gs *Status) analiseHand() {
	gs.numOfColors = make(map[string]int)
	gs.numOfColors["RED"] = 0
	gs.numOfColors["BLUE"] = 0
	gs.numOfColors["GREEN"] = 0
	gs.numOfColors["YELLOW"] = 0
	for _, c := range gs.Hand {
		switch c.Color {
		case gs.DiscardedCard.Color:
			gs.possibleCol = append(gs.possibleCol, c)
		case "":
			gs.possibleCol = append(gs.possibleCol, c)
		}
		if c.Value == gs.DiscardedCard.Value && numbers[gs.DiscardedCard.Value] > 0 && numbers[gs.DiscardedCard.Value] < 10 {
			gs.possibleVal = append(gs.possibleVal, c)
		}
		if c.Color != "" {
			gs.numOfColors[c.Color] = gs.numOfColors[c.Color] + 1
		}
	}
}

func (gs *Status) getBestCard() (card *Card) {
	if len(gs.possibleCol) == 0 && len(gs.possibleVal) == 0 {
		return nil
	}

	if len(gs.possibleCol) > 0 {
		card = &gs.possibleCol[0]
		for _, c := range gs.possibleCol {
			if numbers[c.Value] > numbers[card.Value] {
				card = &c
			}
		}
	} else {
		card = &gs.possibleVal[0]
		for _, c := range gs.possibleVal {
			if numbers[c.Value] > numbers[card.Value] {
				card = &c
			}
		}
	}


	if numbers[card.Value] < -1 {
		max := 0
		for c, n := range gs.numOfColors {
			if n > max {
				card.Color = c
				max = n
			}
		}
	}

	return card
}