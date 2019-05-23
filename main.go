package main

import (
	"CronIbsGen2/functions"
	"CronIbsGen2/scheduler"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DestoyLoginTime := functions.ParseTimeScheduler(os.Getenv("DESTROYLOGIN_TIME"))
	IntervalHarianDepositoTime := functions.ParseTimeScheduler(os.Getenv("INTERVALHARIANDEPOSITO_TIME"))
	functions.Logger().Info("Starting Scheduler IbsGen2")
	gocron.Every(1).Day().At(string(DestoyLoginTime)).Do(scheduler.DestroyLogin)
	gocron.Every(1).Day().At(string(IntervalHarianDepositoTime)).Do(scheduler.IntervalHarianDeposito)
	<-gocron.Start()
	s := gocron.NewScheduler()
	<-s.Start()
}
