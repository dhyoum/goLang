package main

import (
	"fmt"
	"math/rand"
)

const (
	GAME_NUM   int = 5
	PLAYER_NUM int = 2
)

var cards = make([]int, 20)
var players [PLAYER_NUM][2]int

func compare(p int, last_winner int) bool {

	player_result := (players[p][0] + players[p][1]) % 10
	if players[p][0] == players[p][1] {
		player_result += 10 * (players[p][0])
	}
	last_result := (players[last_winner][0] + players[last_winner][1]) % 10
	if players[last_winner][0] == players[last_winner][1] {
		last_result += 10 * (players[last_winner][0])
	}
	if last_result < player_result {
		return true
	}
	return false
}

func main() {
	var result [GAME_NUM]int

	for i := 0; i < 10; i++ {
		cards[i] = i + 1
		cards[i+10] = i + 1
	}
	fmt.Println("Cards : ", cards)

	for num := 0; num < GAME_NUM; num++ {
		for p := 0; p < PLAYER_NUM; p++ {
			players[p][0] = cards[rand.Intn(20)]
			players[p][1] = cards[rand.Intn(20)]
		}
		fmt.Printf(" ======  GAME : %d =======\n", num)
		fmt.Println("The card of player : ", players)
		winner := 0
		for p := 0; p < PLAYER_NUM; p++ {
			if compare(p, winner) {
				winner = p
			}
		}
		result[num] = winner
		fmt.Println("Winner : ", winner)
	}
	fmt.Println("Total Result : ", result)
}
