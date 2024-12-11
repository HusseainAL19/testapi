package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type SqlConnectionCountStruct struct {
	ConnectinoCount int
	Connections     *sql.DB
}

func GetConnection(connectionObj SqlConnectionCountStruct) {
	//return connectionObj.Connections[connectionObj.ConnectinoCount-1]
}

func InitConnection() SqlConnectionCountStruct {
	// set default connection struct
	SqlConnectionStatusStruct := SqlConnectionCountStruct{}

	// connection config;
	db, err := sql.Open(
		"mysql",
		"zabi:uFfHqAUKNRJHjHRjuE39cEn7rnkvSNmJ@tcp(localhost:3306)/schoolsystem",
	)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	openConnections := db.Stats().OpenConnections
	//println("open connection is : ", openConnections)

	if openConnections > 0 {
		println("connection is more than one")
		SqlConnectionStatusStruct.ConnectinoCount = openConnections
		SqlConnectionStatusStruct.Connections = db
		return SqlConnectionStatusStruct
	}

	// check ping the db as simple db test
	//errPing := db.Ping()
	//if errPing != nil {
	//	fmt.Println("cannot connect to sql database")
	//}

	// check simple connection if ther is any error
	if err != nil {
		fmt.Println("cannot init the db")
		fmt.Println(err)
	}
	// return the struct
	// setup the struct point
	SqlConnectionStatusStruct.ConnectinoCount++
	SqlConnectionStatusStruct.Connections = db
	// check open connections
	// finalizying the return
	//fmt.Println(SqlConnectionStatusStruct)
	return SqlConnectionStatusStruct
}
