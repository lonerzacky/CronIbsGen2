package scheduler

import (
	"CronIbsGen2/database"
	"CronIbsGen2/functions"
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"github.com/vjeantet/jodaTime"
	"io/ioutil"
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
		//panic(err.Error())
		functions.Logger().Error(err.Error())
		functions.InsertLogCron("DestroyLogin", "", err.Error(), conn)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(flag)
	if err != nil {
		functions.Logger().Error(err.Error())
	} else {
		functions.InsertLogCron("DestroyLogin", "", "Successfully Destroy All User Login", conn)
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
	requestParams, _ := json.Marshal(jsonData)
	request := gorequest.New()
	resp, _, _ := request.Post(""+os.Getenv("IP_SERVICE_HARIAN_AKHIR")+"/09003").
		Set("Content-Type", "application/json").
		Send(jsonData).
		End()
	fmt.Println(resp.Body)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			functions.Logger().Error(err.Error())
			functions.InsertLogCron("IntervalHarianTabungan", string(requestParams), err.Error(), conn)
			return
		}
		bodyString := string(bodyBytes)
		functions.InsertLogCron("IntervalHarianTabungan", string(requestParams), bodyString, conn)
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
	requestParams, _ := json.Marshal(jsonData)
	request := gorequest.New()
	resp, _, _ := request.Post(""+os.Getenv("IP_SERVICE_HARIAN_AKHIR")+"/09003").
		Set("Content-Type", "application/json").
		Send(jsonData).
		End()
	fmt.Println(resp.Body)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			functions.Logger().Error(err.Error())
			functions.InsertLogCron("IntervalHarianDeposito", string(requestParams), err.Error(), conn)
			return
		}
		bodyString := string(bodyBytes)
		functions.InsertLogCron("IntervalHarianDeposito", string(requestParams), bodyString, conn)
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
	requestParams, _ := json.Marshal(jsonData)
	request := gorequest.New()
	resp, _, _ := request.Post(""+os.Getenv("IP_SERVICE_HARIAN_AKHIR")+"/09003").
		Set("Content-Type", "application/json").
		Send(jsonData).
		End()
	fmt.Println(resp.Body)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			functions.Logger().Error(err.Error())
			functions.InsertLogCron("IntervalHarianKredit", string(requestParams), err.Error(), conn)
			return
		}
		bodyString := string(bodyBytes)
		functions.InsertLogCron("IntervalHarianKredit", string(requestParams), bodyString, conn)
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
	requestParams, _ := json.Marshal(jsonData)
	request := gorequest.New()
	resp, _, _ := request.Post(""+os.Getenv("IP_SERVICE_HARIAN_AKHIR")+"/09003").
		Set("Content-Type", "application/json").
		Send(jsonData).
		End()
	fmt.Println(resp.Body)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			functions.Logger().Error(err.Error())
			functions.InsertLogCron("IntervalHarianABP", string(requestParams), err.Error(), conn)
			return
		}
		bodyString := string(bodyBytes)
		functions.InsertLogCron("IntervalHarianABP", string(requestParams), bodyString, conn)
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
	requestParams, _ := json.Marshal(jsonData)
	request := gorequest.New()
	resp, _, _ := request.Post(""+os.Getenv("IP_SERVICE_HARIAN_AKHIR")+"/09003").
		Set("Content-Type", "application/json").
		Send(jsonData).
		End()
	fmt.Println(resp.Body)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			functions.Logger().Error(err.Error())
			functions.InsertLogCron("IntervalHarianAkunting", string(requestParams), err.Error(), conn)
			return
		}
		bodyString := string(bodyBytes)
		functions.InsertLogCron("IntervalHarianAkunting", string(requestParams), bodyString, conn)
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
	requestParams, _ := json.Marshal(jsonData)
	request := gorequest.New()
	resp, _, _ := request.Post(""+os.Getenv("IP_SERVICE_HARIAN_AKHIR")+"/09003").
		Set("Content-Type", "application/json").
		Send(jsonData).
		End()
	fmt.Println(resp.Body)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			functions.Logger().Error(err.Error())
			functions.InsertLogCron("IntervalHarianInventaris", string(requestParams), err.Error(), conn)
			return
		}
		bodyString := string(bodyBytes)
		functions.InsertLogCron("IntervalHarianInventaris", string(requestParams), bodyString, conn)
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
	requestParams, _ := json.Marshal(jsonData)
	request := gorequest.New()
	resp, _, _ := request.Post(""+os.Getenv("IP_SERVICE_HARIAN_AKHIR")+"/09003").
		Set("Content-Type", "application/json").
		Send(jsonData).
		End()
	fmt.Println(resp.Body)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			functions.Logger().Error(err.Error())
			functions.InsertLogCron("IntervalHarianABA", string(requestParams), err.Error(), conn)
			return
		}
		bodyString := string(bodyBytes)
		functions.InsertLogCron("IntervalHarianABA", string(requestParams), bodyString, conn)
		functions.Logger().Info("Successfully Processing Interval ABA")
	} else {
		functions.Logger().Error("An Error Occured, Service is Not Running or Timeout")
	}
}
