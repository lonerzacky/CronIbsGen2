package main

import (
	"CronIbsGen2/scheduler"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Starting Services Scheduler IbsGen2")
	gocron.Every(1).Day().At("21:48").Do(scheduler.DestroyLogin)
	<-gocron.Start()
	s := gocron.NewScheduler()
	<-s.Start()
}
