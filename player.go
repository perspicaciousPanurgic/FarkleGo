package main

type Player struct {
	id     int
	name   string
	points int
	score  int
}

func (p *Player) ChangeName(name string) {
	p.name = name
}

func (p *Player) AddScore() {
	p.score += p.points
}

func (p *Player) AddPoints(points int) {
	p.points += points
}

func (p *Player) ResetPoints() {
	p.points = 0
}
