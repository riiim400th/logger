package main

import (
	"log"

	l "github.com/riiim400th/logger"
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
	l.Log(0, a_)
	l.Log(l.Info, "bbbb", "hogehoge")
	l.P(l.Error, "gefdksjal;df")
	l.P(l.Panic, "hoge?")
}
