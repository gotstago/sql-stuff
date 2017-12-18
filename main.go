package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", false, "enable debugging")
	password      = flag.String("password", "Datacert1", "the database password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "ELM-AMACDONA\\SQL2012R1DEV", "the database server")
	user          = flag.String("user", "sa", "the database user")
	//instance          = flag.String("instance", "SQL2012R1DEV", "the database user")
)

func main() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
		//fmt.Printf(" instance:%s\n", *instance)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d", *server, *user, *password, *port)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("USE [master]	RESTORE DATABASE [db-passport-2.5.191] FROM  DISK = N'C:\\Program Files\\Microsoft SQL Server\\MSSQL11.SQL2012R1DEV\\MSSQL\\Backup\\db-passport-2.5.187-baseline.bak' WITH  FILE = 1,  MOVE N'matter_management' TO N'C:\\Program Files\\Microsoft SQL Server\\MSSQL11.SQL2012R1DEV\\MSSQL\\DATA\\db-passport-2.5.191.MDF',  MOVE N'matter_management_log' TO N'C:\\Program Files\\Microsoft SQL Server\\MSSQL11.SQL2012R1DEV\\MSSQL\\DATA\\db-passport-2.5.191_1.LDF',  NOUNLOAD,  REPLACE,  STATS = 5")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	stmt.QueryRow()
	// var somenumber int64
	// var somechars string
	// err = row.Scan(&somenumber, &somechars)
	// if err != nil {
	// 	log.Fatal("Scan failed:", err.Error())
	// }
	// fmt.Printf("somenumber:%d\n", somenumber)
	// fmt.Printf("somechars:%s\n", somechars)

	fmt.Printf("bye\n")
}
