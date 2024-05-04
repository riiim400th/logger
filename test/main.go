package main

import (
	"log"

	"github.com/riiim400th/logger"
)

func main() {
	type ina struct {
		NickName string
	}
	type a struct {
		ina
		Name  string
		Email string
	}
	var a_ a
	a_.Name = "hoge"
	a_.Email = "hoge@example.com"
	a_.NickName = "hogehoge"
	log.Default().Print(a_)
	logger.Log(logger.Debug, a_)
	logger.Log(logger.Info, "bbbb", "hogehoge")
	logger.Log(logger.Error, "gefdksjal;df")
	logger.Log(logger.Panic, "hoge?")
}
