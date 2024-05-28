package main

import "fmt"

type gambler interface {
	placebet() bool
	chooseGame(game map[string]string) error
	stake(int) error
}

type bet struct {
	// bet details
}

func (b bet) chooseGame(game map[string]string) error {
	fmt.Printf("Choose match: ")
	var match string
	fmt.Scan(&match)
	t, ok := game[match]
	if ok {
		fmt.Printf("Matchday: %s Vs %s ", match, t)
	} else {
		fmt.Println("That match doesn't exist yet!!")
	}
	return nil

}
func (b bet) placebet() bool {
	return (true)
}
func (b bet) stake(amount int) error {
	fmt.Println("usjsjsjsjsjs")
	return nil
}

type teams struct {
	gDe gambler
}

func main() {
	game := map[string]string{
		"Burnley":   "Crystal Palace",
		"Liverpool": "Wolves",
		"Shuffield": "Manchester City",
	}
	GDetails := teams{
		gDe: bet{},
	}
	GDetails.gDe.chooseGame(game)
}
