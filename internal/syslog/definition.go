package syslog

import (
	"log"
	"log/syslog"
)

// Master after running the program, run journalctl -xe
func Master(show bool) {
	if show {
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
