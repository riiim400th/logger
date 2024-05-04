package main

import "github.com/riiim400th/logger"

func main() {
	logger.Log(logger.Debug, "aaa")
	logger.Log(logger.Info, "bbbb", "hogehoge")
	logger.Log(logger.Error, "gefdksjal;df")
	logger.Log(logger.Panic, "hoge?")
}
