package main

import (
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

/* DAEMON */
type Daemon struct {
	Pid int
}

func (d *Daemon) Stop() {
	Log.Println("STOPING daemo...")
	RemovePidfile()
	os.Exit(0)
}

func (d *Daemon) AwaitSignal() {
	await := make(chan os.Signal, 1)
	signal.Notify(
		await,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	sig := <-await
	Log.Printf("RECIEVED SIGNAL: %s\n", sig)
	switch sig {
	// 1) SIGHUP
	case syscall.SIGHUP:
		// RELOAD CONFIG
		Conf.Load()
		return

	// 2) SIGINT
	// 3) SIGQUIT
	// 15) SIGTERM
	case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
		d.Stop()
	}
}

func (d *Daemon) Run() {
	Log.Printf("...STARTING daemo. PID: %d\n", Daemo.Pid)
	for {
		d.AwaitSignal()
	}
	Log.Fatalln("daemo EXITING BAD...")
}

/* PID FILE*/
func (d *Daemon) WritePidfile() {
	p := []byte(strconv.Itoa(d.Pid) + "\n")
	Log.Printf("WRITING \"%s\"...", Conf.Pidfile)
	err := ioutil.WriteFile(Conf.Pidfile, p, 0644)
	if err != nil {
		Log.Fatalln(err)
	}
}

func RemovePidfile() {
	if _, err := os.Stat(Conf.Pidfile); err == nil {
		Log.Printf("REMOVING \"%s\"...", Conf.Pidfile)
		os.Remove(Conf.Pidfile)
	}
}

func DaemoInit() {
	// GLOBAL `Daemo`
	Daemo = &Daemon{
		os.Getpid(),
	}
	Daemo.WritePidfile()
}
