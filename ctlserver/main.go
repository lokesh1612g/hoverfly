package main

import (
	"github.com/namsral/flag"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
)

var (
	destdbhost           = flag.String("destinationdb-host", "localhost", "Destination database host")
	destdbport           = flag.String("destinationdb-port", "3306", "Destination database port")
	destdbuser           = flag.String("destinationdb-user", "", "Destination database user")
	destdbpass           = flag.String("destinationdb-password", "", "Destinatin database password")
	binlogexecutable     = flag.String("binlog-executable", "", "executable location for binlog command")
	mysqlexecutable      = flag.String("mysql-executable", "", "executable location for mysql command")
	mysqldumpexecutable  = flag.String("mysqldump-executable", "", "executable location for mysqldump command")
	mysqladminexecutable = flag.String("mysqladmin-executable", "", "executable location for mysqladmin command")
)

func checkMysql() {
	if *mysqladminexecutable != "" {
		MysqlCommand = *mysqladminexecutable

		if *destdbhost != "" {
			MysqlCommand = MysqlCommand + " -h" + *destdbhost
		}

		if *destdbport != "" {
			MysqlCommand = MysqlCommand + " -P" + *destdbport
		}

		if *destdbuser != "" {
			MysqlCommand = MysqlCommand + " -u" + *destdbuser
		}

		if *destdbpass != "" {
			MysqlCommand = MysqlCommand + " -p" + *destdbpass
		}
	}

	mysqlcmdparameters := strings.Split(MysqlCommand, " ")

	cmd := mysqlcmdparameters[0]
	prms := mysqlcmdparameters[1:]
	prms = append(prms, "ping")

	for {
		mysql := exec.Command(cmd, prms...)
		mysqlerr := mysql.Run()
		if mysqlerr == nil {
			break
		}
		mysql.Process.Kill()
		log.Print(".")
		time.Sleep(1 * time.Second)
	}
}

var MysqlCommand string

func main() {
	flag.Parse()

	checkMysql()

	router := NewRouter()

	RepoCreateState(State{Command: "/bin/hoverfly -listen-on-host=0.0.0.0 -capture -tls-verification=false"})
	time.Sleep(2 * time.Second)
	RepoCreateState(State{Command: "/bin/hoverctl mode capture --stateful"})

	log.Fatal(http.ListenAndServe(":8080", router))
}
