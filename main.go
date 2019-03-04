package main

import (
	"fmt"

	"./char"
)

var config configStruct
var score [99999]int64

func main() {

	config = configInit()
	list := char.ListInit()

	for index := 0; index < config.RoundTimes; index++ {

		for i, iv := range list {

			for j, jv := range list {

				if i < j {

					g := Game{}
					g.Init()
					a, b := g.Run(&iv, &jv)
					score[i] += a
					score[j] += b

					//fmt.Println(iv.Name(), a, jv.Name(), b)

				}

			}

		}
	}

	for i, v := range list {
		fmt.Println(v.Name(), score[i])
	}

}
