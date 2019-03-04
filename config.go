package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type configStruct struct {
	RoundTimes    int   `json:"round_times"`
	GameTimes     int   `json:"game_times"`
	GameInitPoint int64 `json:"game_init_pt"`
	CharInitPoint int64 `json:"char_init_pt"`
	WinPoint      int64 `json:"win_pt"`
	LosePoint     int64 `json:"lose_pt"`
	BothWinPoint  int64 `json:"both_win_pt"`
	BothLosePoint int64 `json:"both_lose_pt"`
}

func configInit() (c configStruct) {

	b, err := ioutil.ReadFile("config.json") // just pass the file name

	if err != nil {
		log.Panicln(err)
	}

	err = json.Unmarshal(b, &c)

	if err != nil {
		log.Panicln(err)
	}

	return c

}
