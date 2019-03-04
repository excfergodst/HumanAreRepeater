package main

import (
	"fmt"

	"./char"
	"github.com/tidwall/gjson"
)

var config gjson.Result
var score [100]int64

func main() {

	config = configInit()
	list := char.ListInit()

	var index int64
	index = 0

	for ; index < config.Get("round_times").Int(); index++ {

		for i, iv := range list {

			for j, jv := range list {

				if i != j {

					g := Game{}
					g.Init()
					a, b := g.Run(iv, jv)
					score[i] += a
					score[j] += b

					fmt.Println(iv.Name(), a, jv.Name(), b)

				}

			}

		}
	}

	for i, v := range list {
		fmt.Println(v.Name(), score[i])
	}

}
