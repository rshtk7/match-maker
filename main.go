package main

import (
	"log"
	"match-maker/MatchMaker"
)

func main(){
	var matchMaker MatchMaker.MatchMaker
	matchMaker.RegisterGame(1,1)
	//matchMaker.RegisterGame(2,4)
	//matchMaker.RegisterGame(3,5)
	matchMaker.RegisterPlayer(1,1,1, 1)
	matchMaker.RegisterPlayer(2,3,2, 1)
	matchMaker.RegisterPlayer(3,0,2,1)
	matchMaker.RegisterPlayer(4,0,2,1)
	assignOpponent(&matchMaker, 1, 1, 0)
	assignOpponent(&matchMaker, 4, 1, 0)
	//matchMaker.AssignOpponent(3,1)
}

func assignOpponent(matchMaker *MatchMaker.MatchMaker, playerId int, delta int, count int) {
	if err := matchMaker.AssignOpponent(playerId, delta); err != nil && count<3{
		log.Println(err)
		count++
		assignOpponent(matchMaker, playerId, delta+1, count)
	} else if count == 3 {
		log.Printf("Please wait for the suitable opponent\n")
	}

}