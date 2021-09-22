package Game

type Game struct{
	id int
}

type Gamer interface {
	GetId() int
	SetId(int)
}

func (g *Game) GetId() int{
	return g.id
}

func (g *Game) SetId(i int) {
	g.id = i
}
