package functions

import (
	"CronIbsGen2/database"
	"github.com/vjeantet/jodaTime"
	"net"
	"os"
	"time"
)

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

//noinspection SqlDialectInspection,SqlNoDataSourceInspection
func InsertLogCron(scheduler string) {
	var db = database.ConnectDB()
	stmt, err := db.Prepare("INSERT INTO logcron(scheduler,ip_address, message,tgl_proses) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	ipAdd := GetIpAdd()
	tglProses := jodaTime.Format("YYYY-MM-dd HH:mm:ss", time.Now())
	stmt.Exec(scheduler, ipAdd, "Successfully Destroy All user", tglProses)
	//noinspection GoUnhandledErrorResult
	defer db.Close()
}
