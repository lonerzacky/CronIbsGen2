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
		functions.Logger().Error(err.Error())
	} else {
		functions.InsertLogCron("DestroyLogin", "Successfully Destroy All User Login", conn)
		functions.Logger().Info("Successfully Destroy All User Login")
	}
}

func IntervalHarianTabungan() {
	functions.Logger().Info("Starting Interval Harian Tabungan")
	jsonData := ProsesIntervalHarian{
		TglAwal:     dateNowYmd,
		TglAkhir:    dateNowYmd,
		JenisProduk: "tabungan",
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
		functions.InsertLogCron("IntervalHarianTabungan", bodyString, conn)
		functions.Logger().Info("Successfully Processing Interval Tabungan")
	} else {
		functions.Logger().Error("An Error Occured, Service is Not Running or Timeout")
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
	} else {
		functions.Logger().Error("An Error Occured, Service is Not Running or Timeout")
	}
}

func IntervalHarianKredit() {
	functions.Logger().Info("Starting Interval Harian Kredit")
	jsonData := ProsesIntervalHarian{
		TglAwal:     dateNowYmd,
		TglAkhir:    dateNowYmd,
		JenisProduk: "kredit",
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
		functions.InsertLogCron("IntervalHarianKredit", bodyString, conn)
		functions.Logger().Info("Successfully Processing Interval Kredit")
	} else {
		functions.Logger().Error("An Error Occured, Service is Not Running or Timeout")
	}
}

func IntervalHarianABP() {
	functions.Logger().Info("Starting Interval Harian ABP")
	jsonData := ProsesIntervalHarian{
		TglAwal:     dateNowYmd,
		TglAkhir:    dateNowYmd,
		JenisProduk: "abp",
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
		functions.InsertLogCron("IntervalHarianABP", bodyString, conn)
		functions.Logger().Info("Successfully Processing Interval ABP")
	} else {
		functions.Logger().Error("An Error Occured, Service is Not Running or Timeout")
	}
}

func IntervalHarianAkunting() {
	functions.Logger().Info("Starting Interval Harian Akunting")
	jsonData := ProsesIntervalHarian{
		TglAwal:     dateNowYmd,
		TglAkhir:    dateNowYmd,
		JenisProduk: "akunting",
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
		functions.InsertLogCron("IntervalHarianAkunting", bodyString, conn)
		functions.Logger().Info("Successfully Processing Interval Akunting")
	} else {
		functions.Logger().Error("An Error Occured, Service is Not Running or Timeout")
	}
}

func IntervalHarianInventaris() {
	functions.Logger().Info("Starting Interval Harian Inventaris")
	jsonData := ProsesIntervalHarian{
		TglAwal:     dateNowYmd,
		TglAkhir:    dateNowYmd,
		JenisProduk: "inventaris",
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
		functions.InsertLogCron("IntervalHarianInventaris", bodyString, conn)
		functions.Logger().Info("Successfully Processing Interval Inventaris")
	} else {
		functions.Logger().Error("An Error Occured, Service is Not Running or Timeout")
	}
}

func IntervalHarianABA() {
	functions.Logger().Info("Starting Interval Harian ABA")
	jsonData := ProsesIntervalHarian{
		TglAwal:     dateNowYmd,
		TglAkhir:    dateNowYmd,
		JenisProduk: "aba",
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
		functions.InsertLogCron("IntervalHarianABA", bodyString, conn)
		functions.Logger().Info("Successfully Processing Interval ABA")
	} else {
		functions.Logger().Error("An Error Occured, Service is Not Running or Timeout")
	}
}
