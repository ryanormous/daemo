package main

import (
	"log"
)

/* GLOBAL */
var Log *log.Logger
var Conf *Configuration
var Daemo *Daemon

/* INIT */
func init() {
	ConfInit()
	LogInit()
	DaemoInit()
}

/* MAIN */
func main() {
	Conf.Print()
	Daemo.Run()
}
