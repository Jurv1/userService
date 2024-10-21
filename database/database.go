package database

import (
	"fmt"
	"github.com/jackc/pgx"
	"os"
)

type db struct {
	dbConn *pgx.ConnPool
}

func GetNewDB() *pgx.Conn {
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     "",
		Port:     0,
		Database: "",
		User:     "",
		Password: "",
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}
