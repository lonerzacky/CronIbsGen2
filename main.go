package main

import (
	"CronIbsGen2/database"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"log"
	"os"
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
	gocron.Every(1).Day().At(os.Getenv("PROCESSING_TIME")).Do(DestroyLogin)
	<-gocron.Start()
	s := gocron.NewScheduler()
	<-s.Start()
}
