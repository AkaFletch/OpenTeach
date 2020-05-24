package main

import (
	log "github.com/jeanphorn/log4go"
)

func main() {
	log.LoadConfiguration("logging.json")
	log.Info("OpenTeach starting")
	log.Close()
}
