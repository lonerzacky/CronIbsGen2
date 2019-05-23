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
	IntervalHarianTabunganTime := functions.ParseTimeScheduler(os.Getenv("INTERVALHARIANTABUNGAN_TIME"))
	IntervalHarianDepositoTime := functions.ParseTimeScheduler(os.Getenv("INTERVALHARIANDEPOSITO_TIME"))
	IntervalHarianKreditTime := functions.ParseTimeScheduler(os.Getenv("INTERVALHARIANKREDIT_TIME"))
	IntervalHarianABPTime := functions.ParseTimeScheduler(os.Getenv("INTERVALHARIANABP_TIME"))
	IntervalHarianAkuntingTime := functions.ParseTimeScheduler(os.Getenv("INTERVALHARIANAKUNTING_TIME"))
	IntervalHarianInventarisTime := functions.ParseTimeScheduler(os.Getenv("INTERVALHARIANINVENTARIS_TIME"))
	IntervalHarianABATime := functions.ParseTimeScheduler(os.Getenv("INTERVALHARIANABA_TIME"))
	functions.Logger().Info("Starting Scheduler IbsGen2")
	gocron.Every(1).Day().At(string(DestoyLoginTime)).Do(scheduler.DestroyLogin)
	gocron.Every(1).Day().At(string(IntervalHarianTabunganTime)).Do(scheduler.IntervalHarianTabungan)
	gocron.Every(1).Day().At(string(IntervalHarianDepositoTime)).Do(scheduler.IntervalHarianDeposito)
	gocron.Every(1).Day().At(string(IntervalHarianKreditTime)).Do(scheduler.IntervalHarianKredit)
	gocron.Every(1).Day().At(string(IntervalHarianABPTime)).Do(scheduler.IntervalHarianABP)
	gocron.Every(1).Day().At(string(IntervalHarianAkuntingTime)).Do(scheduler.IntervalHarianAkunting)
	gocron.Every(1).Day().At(string(IntervalHarianInventarisTime)).Do(scheduler.IntervalHarianInventaris)
	gocron.Every(1).Day().At(string(IntervalHarianABATime)).Do(scheduler.IntervalHarianABA)
	<-gocron.Start()
}
