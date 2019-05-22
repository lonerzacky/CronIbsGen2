package scheduler

import (
	"CronIbsGen2/database"
	"CronIbsGen2/functions"
	"fmt"
	"log"
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
		functions.InsertLogCron("DestroyLogin", conn)
		log.Println("Successfully Destroy All user")
	}
	//noinspection GoUnhandledErrorResult
	defer connSys.Close()
}
