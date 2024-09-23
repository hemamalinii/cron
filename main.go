package main

import (
	"time"

	"github.com/robfig/cron"
)

func main() {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}
	cronJob := cron.NewWithLocation(loc)
	cronJob.AddFunc("0 0 10 * * *", func() {

	})
	cronJob.Start()
	select {}
}
