package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Dealer struct {
	card Card
}

type Player struct {
	card [2]int
}

func (p *Player) getPriority() (priority int) {
	priority = (p.card[0] + p.card[1]) % 10
	if p.card[0] == p.card[1] {
		priority = 10 * p.card[0]
	}
	return
}

type Card [20]int

func createCard() Card {
	var card Card
	for i, _ := range card {
		card[i] = i%10 + 1
	}
	return card
}

func (card *Card) show() {
	fmt.Println("card : ", *card)
}

func (card *Card) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i, value := range card {
		t := rand.Intn(20)
		temp := value
		card[i] = card[t]
		card[t] = temp
	}
}

func (dealer *Dealer) deal() (p1, p2 Player) {
	dealer.card.shuffle()
	p1.card[0] = dealer.card[0]
	p1.card[1] = dealer.card[1]

	p2.card[0] = dealer.card[2]
	p2.card[1] = dealer.card[3]

	return
}

type Game struct {
	history [100]int
}

func (game *Game) save(index, result int) {
	game.history[index] = result
}

func (game *Game) check(p1, p2 *Player) (result int) {
	if p1.getPriority() == p2.getPriority() {
		result = 0
	} else if p1.getPriority() < p2.getPriority() {
		result = 1
	} else {
		result = 2
	}
	return
}

func (game *Game) show() {

	var result = make(map[int]int)
	for times := 0; times < 100; times++ {
		if value, ok := result[game.history[times]]; ok {
			result[game.history[times]] = value + 1
		} else {
			result[game.history[times]] = 1
		}
	}
	fmt.Println(game.history)
	fmt.Println("Result of Game", result)
	fmt.Printf("draw : %0.2f%%\n", float32(100*result[0]/100))
	fmt.Printf("Player1 : %0.2f%%\n", float32(100*result[1]/100))
	fmt.Printf("Player2 : %0.2f%%\n", float32(100*result[2]/100))
}

func main() {
	var dealer Dealer
	dealer.card = createCard()

	var game Game
	for i := 0; i < 100; i++ {
		p1, p2 := dealer.deal()
		result := game.check(&p1, &p2)
		game.save(i, result)
	}
	game.show()
}
