package scheduler

import (
	"CronIbsGen2/database"
	"CronIbsGen2/functions"
	"fmt"
	"log"
)

//noinspection SqlDialectInspection,SqlNoDataSourceInspection
func DestroyLogin() {
	var dbSys = database.ConnectDBSys()
	flag := 0
	stmt, err := dbSys.Prepare("UPDATE sys_daftar_user SET flag=?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(flag)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		functions.InsertLogCron("DestroyLogin")
		log.Println("Successfully Destroy All user")
	}
	//noinspection GoUnhandledErrorResult
	defer dbSys.Close()
}
