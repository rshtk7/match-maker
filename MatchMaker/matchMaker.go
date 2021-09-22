package MatchMaker

import (
	"fmt"
	"log"
	"match-maker/Game"
	"match-maker/Player"
	"match-maker/Team"
	"math/rand"
)
type MatchMaker struct{
	caseGameMap map[int][]int // map of case_id with game_id
	gameMap map[int]*Game.Game
	playerMap map[int]*Player.Player
	gamePlayerMap map[int][]int
	gameTeamMap map[int][]int
}

func (m *MatchMaker) GameTeamMap() map[int][]int {
	return m.gameTeamMap
}

func (m *MatchMaker) SetGameTeamMap(gameTeamMap map[int][]int) {
	m.gameTeamMap = gameTeamMap
}

type MatchMakerer interface {
	RegisterGame(int, int)
	RegisterPlayer(int, int, int, int)
	AssignOpponent(int, int)
}

func (m *MatchMaker) RegisterGame(caseId int, gameId int){
	if m.CaseGameMap() == nil{
		caseGameMap := make(map[int][]int, 1)
		gameList := make([]int,1)
		gameList[0] = gameId
		caseGameMap[caseId] = gameList
		m.SetCaseGameMap(caseGameMap)
	}else{
		caseGameMap := m.CaseGameMap()
		_, caseExist := caseGameMap[caseId]
		if caseExist==false{
			gameList := caseGameMap[caseId]
			gameList = append(gameList, gameId)
			caseGameMap[caseId] = gameList
			m.SetCaseGameMap(caseGameMap)
		}
	}
	log.Printf("registered game : %d\n", gameId)
}

func (m *MatchMaker) RegisterPlayer(playerId int, score int, caseId int, gameId int) {
	//making a player entity and setting the player details
	var player Player.Player
	player.SetId(playerId)
	player.SetCaseId(caseId)
	player.SetGameId(gameId)
	player.SetStatus(true)
	player.SetScore(score)

	//storing the player details in player map
	if m.PlayerMap() == nil{
		playerMap := make(map[int]*Player.Player, 1)
		playerMap[playerId] = &player
		m.SetPlayerMap(playerMap)
	}else{
		playerMap := m.PlayerMap()
		playerMap[playerId] = &player
		m.SetPlayerMap(playerMap)
	}

	//storing the mapping of game and player
	if m.GamePlayerMap() == nil{
		gamePlayerMap := make(map[int][]int, 1)
		playerList := make([]int, 1)
		playerList[0] = playerId
		gamePlayerMap[gameId] = playerList
		m.SetGamePlayerMap(gamePlayerMap)
	}else{
		gamePlayerMap := m.GamePlayerMap()
		playerList := gamePlayerMap[gameId]
		playerList = append(playerList, playerId)
		gamePlayerMap[gameId] = playerList
		m.SetGamePlayerMap(gamePlayerMap)
	}

	log.Printf("registered player : %d, game : %d\n", playerId, gameId)
}

func (m *MatchMaker) AssignOpponent(playerId int, delta int) error{
	playerMap := m.PlayerMap()
	_, playerExist := playerMap[playerId]
	if playerExist==false{
		return fmt.Errorf("Player not registered : %d\n", playerId)
	}else{
		player := playerMap[playerId]
		score := player.Score()
		gameId := player.GameId()
		minScore := score - delta
		maxScore := score + delta
		targetPlayerList := make([]int,0)
		gamePlayerMap := m.GamePlayerMap()
		_, gameExist := gamePlayerMap[gameId]
		if gameExist==true{
			playerList := gamePlayerMap[gameId]
			for i:=0; i< len(playerList); i++{
				playerObj := playerMap[playerList[i]]
				if playerObj.Id() == playerId{
					continue
				}
				if playerObj.Score() >= minScore && playerObj.Score() <= maxScore && playerObj.Status()==true {
					var team Team.Team
					playerIdList := make([]int, 2)
					playerIdList[0] = playerId
					playerIdList[1] = playerObj.Id()
					team.SetId(rand.Intn(100))
					team.SetPlayerIdList(playerIdList)
					//storing the mapping of game and team
					//storing the mapping of game and player
					if m.GameTeamMap() == nil{
						gameTeamMap := make(map[int][]int, 1)
						teamList := make([]int, 1)
						teamList[0] = team.Id()
						gameTeamMap[gameId] = teamList
						m.SetGameTeamMap(gameTeamMap)
					}else{
						gameTeamMap := m.GameTeamMap()
						teamList := gameTeamMap[gameId]
						teamList = append(teamList, team.Id())
						gameTeamMap[gameId] = teamList
						m.SetGameTeamMap(gameTeamMap)
					}
					player.SetStatus(false)
					playerObj.SetStatus(false)
					fmt.Printf("Assigned Opponent : %d for player : %d\n", playerObj.Id(), playerId)
					targetPlayerList = append(targetPlayerList, playerObj.Id())
					//for 1-1 case
					return nil
				}
			}
			return fmt.Errorf("Could not find opponent for player : %d, delta : %d \n", playerId, delta)
		}
	}
	return nil
}

func (m *MatchMaker) CaseGameMap() map[int][]int {
	return m.caseGameMap
}

func (m *MatchMaker) SetCaseGameMap(caseGameMap map[int][]int) {
	m.caseGameMap = caseGameMap
}

func (m *MatchMaker) GameMap() map[int]*Game.Game {
	return m.gameMap
}

func (m *MatchMaker) SetGameMap(gameMap map[int]*Game.Game) {
	m.gameMap = gameMap
}

func (m *MatchMaker) PlayerMap() map[int]*Player.Player {
	return m.playerMap
}

func (m *MatchMaker) SetPlayerMap(playerMap map[int]*Player.Player) {
	m.playerMap = playerMap
}

func (m *MatchMaker) GamePlayerMap() map[int][]int {
	return m.gamePlayerMap
}

func (m *MatchMaker) SetGamePlayerMap(gamePlayerMap map[int][]int) {
	m.gamePlayerMap = gamePlayerMap
}
