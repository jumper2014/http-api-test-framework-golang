package main


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DbWorker struct {
	//mysql data source name
	Dsn string
}

func main() {
	dbw := DbWorker{
		Dsn: "user:password@tcp(192.68.1.61:3306)/gother",
	}
	db, err := sql.Open("mysql",
		dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()
}