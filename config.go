package main

import (
	"io/ioutil"
	"log"

	"github.com/tidwall/gjson"
)

func configInit() gjson.Result {

	b, err := ioutil.ReadFile("config.json") // just pass the file name

	if err != nil {
		log.Panicln(err)
	}

	return gjson.ParseBytes(b)

}
