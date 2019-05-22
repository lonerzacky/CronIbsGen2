package main

import (
	"CronIbsGen2/database"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"log"
)

//noinspection SqlDialectInspection,SqlNoDataSourceInspection
func DestroyLogin() {
	var db = database.ConnectDB()
	flag := 0
	insForm, err := db.Prepare("UPDATE sys_daftar_user SET flag=?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(flag)
	log.Println("Successfully Destroy All User")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Starting Services Scheduler IbsGen2")
	gocron.Every(1).Day().At("01:00").Do(DestroyLogin)
	<-gocron.Start()
	s := gocron.NewScheduler()
	<-s.Start()
}
