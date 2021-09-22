package Player

type Player struct{
	id int
	score int
	gameId int
	caseId int
	teamId int
	status bool
}

func (p *Player) Id() int {
	return p.id
}

func (p *Player) SetId(id int) {
	p.id = id
}

func (p *Player) Score() int {
	return p.score
}

func (p *Player) SetScore(score int) {
	p.score = score
}

func (p *Player) GameId() int {
	return p.gameId
}

func (p *Player) SetGameId(gameId int) {
	p.gameId = gameId
}

func (p *Player) CaseId() int {
	return p.caseId
}

func (p *Player) SetCaseId(caseId int) {
	p.caseId = caseId
}

func (p *Player) TeamId() int {
	return p.teamId
}

func (p *Player) SetTeamId(teamId int) {
	p.teamId = teamId
}

func (p *Player) Status() bool {
	return p.status
}

func (p *Player) SetStatus(status bool) {
	p.status = status
}



