package DB_conn

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "db"
	port = 5432
	user = "postgres"
	password = "12340"
	dbname = "testtask"
)


func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil{
		panic(err)
	}
	err = conn.Ping()
	if err!= nil {
		panic(err)
	}
	return conn
}
