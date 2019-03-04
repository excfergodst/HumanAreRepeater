package main

import (
	"math/rand"
	"time"

	"./char"
)

type Game struct {
	times   int64
	AStatus char.Status
	BStatus char.Status
}

func (g *Game) Init() {
	rand.Seed(time.Now().UTC().UnixNano())
	g.times = config.Get("game_times").Int()
	g.AStatus.MyPoint = config.Get("init_pt").Int()
	g.BStatus.MyPoint = config.Get("init_pt").Int()
}

func (g *Game) Run(charA char.Char, charB char.Char) (int64, int64) {

	var i int64
	i = 0

	for ; i < g.times; i++ {

		decisionA := charA.Run(g.AStatus)
		decisionB := charB.Run(g.BStatus)

		g.AStatus.RecordStep(decisionA, decisionB)
		g.BStatus.RecordStep(decisionB, decisionA)

		if decisionA == 1 && decisionB == 1 {
			g.AStatus.MyPoint += config.Get("both_win_pt").Int()
			g.BStatus.MyPoint += config.Get("both_win_pt").Int()
		} else if decisionA == 1 && decisionB == 0 {
			g.AStatus.MyPoint += config.Get("lose_pt").Int()
			g.BStatus.MyPoint += config.Get("win_pt").Int()
		} else if decisionA == 0 && decisionB == 1 {
			g.AStatus.MyPoint += config.Get("win_pt").Int()
			g.BStatus.MyPoint += config.Get("lose_pt").Int()
		} else if decisionA == 0 && decisionB == 0 {
			g.AStatus.MyPoint += config.Get("both_lose_pt").Int()
			g.BStatus.MyPoint += config.Get("both_lose_pt").Int()
		}

		g.AStatus.OppoPoint = g.BStatus.MyPoint
		g.BStatus.OppoPoint = g.AStatus.MyPoint

	}

	return g.AStatus.MyPoint, g.BStatus.MyPoint
}
