package main

import (
	"flag"
	"log"
	"os"
	"syscall"
)

// AppInfo - stores information about current Gogling instance
type AppInfo struct {
	Name       string
	Version    string
	Done       chan bool
	ConfigFile string
	Config     *Config
}

// StdoutFile - file wrapper for stdout
var StdoutFile = os.NewFile(uintptr(syscall.Stdout), "stdout")

func main() {
	ai := &AppInfo{
		Name:       "Gogling",
		Version:    "0.0.2-pre2",
		Done:       make(chan bool),
		ConfigFile: "./config.json",
		Config:     nil,
	}

	flag.StringVar(&ai.ConfigFile, "config", "./config.json", "Config file")
	flag.Parse()

	l := log.New(StdoutFile, "main ", -1)
	l.Printf("This is %s v%s", ai.Name, ai.Version)

	ai.Config = ConfigLoad(ai.ConfigFile)

	l.Printf("Will listen on %s:%s", ai.Config.Host, ai.Config.Port)

	go Serve(ai)

	<-ai.Done
	l.Println("Exiting")
}
