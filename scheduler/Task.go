package scheduler

import (
	"CronIbsGen2/database"
	"CronIbsGen2/functions"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/vjeantet/jodaTime"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var conn = database.ConnectDB()
var connSys = database.ConnectDBSys()
var dateNowYmd = jodaTime.Format("YYYYMMdd", time.Now())

type ProsesIntervalHarian struct {
	TglAwal     string `json:"tgl_awal"`
	TglAkhir    string `json:"tgl_akhir"`
	JenisProduk string `json:"jenis_produk"`
	KodeKantor  string `json:"kode_kantor"`
}

//noinspection SqlDialectInspection,SqlNoDataSourceInspection
func DestroyLogin() {
	functions.Logger().Info("Starting Destory User Login")
	flag := 0
	stmt, err := connSys.Prepare("UPDATE sys_daftar_user SET flag=?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(flag)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		functions.InsertLogCron("DestroyLogin", "Successfully Destroy All User Login", conn)
		functions.Logger().Info("Successfully Destroy All User Login")
	}
}

func IntervalHarianDeposito() {
	functions.Logger().Info("Starting Interval Harian Deposito")
	jsonData := ProsesIntervalHarian{
		TglAwal:     dateNowYmd,
		TglAkhir:    dateNowYmd,
		JenisProduk: "deposito",
		KodeKantor:  "",
	}
	request := gorequest.New()
	resp, _, _ := request.Post(""+os.Getenv("IP_SERVICE_HARIAN_AKHIR")+"/09003").
		Set("Content-Type", "application/json").
		Send(jsonData).
		End()
	fmt.Println(resp.Body)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		functions.InsertLogCron("IntervalHarianDeposito", bodyString, conn)
		functions.Logger().Info("Successfully Processing Interval Deposito")
	}
}
