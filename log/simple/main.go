package main

import (
	"time"

	"github.com/dobyte/due/v2/log"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		now := <-ticker.C

		log.Infof("current time: %s", now.Format("2006-01-02 15:04:05"))
	}
}
