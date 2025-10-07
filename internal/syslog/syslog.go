package syslog

import (
	"fmt"
	"log"
	"log/syslog"
)

// Index after running the program, run journalctl -xe
func Run(show bool) {
	if show {
		fmt.Println("-- Syslog in Go")

		sysLog, err := syslog.New(syslog.LOG_SYSLOG, "main.go")
		if err != nil {
			log.Println(err)
			return
		} else {
			log.SetOutput(sysLog)
			log.Print("Everything is fine!")
		}
	}
}
