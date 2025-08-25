package main

import (
	"fmt"
	"time"

	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xtime"
)

func main() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	log.Infof("current time: %s", xtime.Now().Format("2006-01-02 15:04:05"))

	for {
		<-ticker.C

		fmt.Println(xtime.Now().Format("2006-01-02 15:04:05"))

		// log.Infof("current time: %s", now.Format("2006-01-02 15:04:05"))
	}
}
