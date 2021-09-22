package Team

type Team struct{
	id int
	gameId int
	status bool
	playerIdList []int
}

func (t *Team) Id() int {
	return t.id
}

func (t *Team) SetId(id int) {
	t.id = id
}

func (t *Team) GameId() int {
	return t.gameId
}

func (t *Team) SetGameId(gameId int) {
	t.gameId = gameId
}

func (t *Team) Status() bool {
	return t.status
}

func (t *Team) SetStatus(status bool) {
	t.status = status
}

func (t *Team) PlayerIdList() []int {
	return t.playerIdList
}

func (t *Team) SetPlayerIdList(playerIdList []int) {
	t.playerIdList = playerIdList
}
