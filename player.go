package main

type Player struct {
	id    int
	name  string
	score int
}

func (p *Player) ChangeName(name string) {
	p.name = name
}

func (p *Player) AddScore(points int) {
	p.score += points
}
