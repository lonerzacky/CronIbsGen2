package scheduler

import (
	"CronIbsGen2/database"
	"CronIbsGen2/functions"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var conn = database.ConnectDB()
var connSys = database.ConnectDBSys()

//noinspection SqlDialectInspection,SqlNoDataSourceInspection
func DestroyLogin() {
	flag := 0
	stmt, err := connSys.Prepare("UPDATE sys_daftar_user SET flag=?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(flag)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		functions.InsertLogCron("DestroyLogin", "Successfully Destroy All user", conn)
		log.Println("Successfully Destroy All user")
	}
}

func AutoLogin() {
	request := gorequest.New()
	resp, _, _ := request.Post(""+os.Getenv("IP_SERVICE")+"/010000").
		Set("Content-Type", "application/json").
		Send(`{"username":"USSI","password":"7376cbab0e92be1a14377f15202424ecde2fbf5f","kode_kantor":""}`).
		End()
	fmt.Println(resp.Body)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		functions.InsertLogCron("AutoLogin", bodyString, conn)
	}
}
