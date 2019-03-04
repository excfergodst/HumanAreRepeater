package main

import (
	"math/rand"
	"time"

	"./char"
)

type Game struct {
	times   int
	AStatus char.Status
	BStatus char.Status
}

func (g *Game) Init() {
	rand.Seed(time.Now().UTC().UnixNano())
	g.times = config.GameTimes
	g.AStatus.MyPoint = config.GameInitPoint
	g.BStatus.MyPoint = config.GameInitPoint
}

func (g *Game) Run(charA *char.Char, charB *char.Char) (int64, int64) {

	for i := 0; i < g.times; i++ {

		g.RunOnce(charA, charB)

	}

	return g.AStatus.MyPoint, g.BStatus.MyPoint
}

func (g *Game) Judge(dA, dB int64) {

	if dA == 1 && dB == 1 {
		g.AStatus.MyPoint += config.BothWinPoint
		g.BStatus.MyPoint += config.BothWinPoint
	} else if dA == 1 && dB == 0 {
		g.AStatus.MyPoint += config.LosePoint
		g.BStatus.MyPoint += config.WinPoint
	} else if dA == 0 && dB == 1 {
		g.AStatus.MyPoint += config.WinPoint
		g.BStatus.MyPoint += config.LosePoint
	} else if dA == 0 && dB == 0 {
		g.AStatus.MyPoint += config.BothLosePoint
		g.BStatus.MyPoint += config.BothLosePoint
	}

}

func (g *Game) RunOnce(charA *char.Char, charB *char.Char) {

	decisionA := (*charA).Run(g.AStatus)
	decisionB := (*charB).Run(g.BStatus)

	g.Judge(decisionA, decisionB)

	g.AStatus.RecordStep(decisionA, decisionB)
	g.BStatus.RecordStep(decisionB, decisionA)

	g.AStatus.OppoPoint = g.BStatus.MyPoint
	g.BStatus.OppoPoint = g.AStatus.MyPoint

}
