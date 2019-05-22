package main

import (
	"CronIbsGen2/database"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"github.com/vjeantet/jodaTime"
	"log"
	"net"
	"os"
	"time"
)

var db = database.ConnectDB()
var dbSys = database.ConnectDBSys()
var mDestroyLogin = "Successfully Destroy All user"

//noinspection SqlDialectInspection,SqlNoDataSourceInspection
func DestroyLogin() {
	flag := 0
	stmt, err := dbSys.Prepare("UPDATE sys_daftar_user SET flag=?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(flag)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		InsertLogCron("DestroyLogin")
		log.Println(mDestroyLogin)
	}
	//noinspection GoUnhandledErrorResult
	defer dbSys.Close()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Starting Services Scheduler IbsGen2")
	gocron.Every(1).Day().At("20:36").Do(DestroyLogin)
	<-gocron.Start()
	s := gocron.NewScheduler()
	<-s.Start()
}

//noinspection SqlDialectInspection,SqlNoDataSourceInspection
func InsertLogCron(scheduler string) {
	stmt, err := db.Prepare("INSERT INTO logcron(scheduler,ip_address, message,tgl_proses) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	ipAdd := GetIpAdd()
	tglProses := jodaTime.Format("YYYY-MM-dd HH:mm:ss", time.Now())
	stmt.Exec(scheduler, ipAdd, mDestroyLogin, tglProses)
	//noinspection GoUnhandledErrorResult
	defer db.Close()
}

func GetIpAdd() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		_, _ = os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}
	var ipadd = ""
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipadd = ipnet.IP.String()
			}
		}
	}
	return ipadd

}
