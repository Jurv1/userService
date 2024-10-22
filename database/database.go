package database

import (
	"fmt"
	"github.com/Jurv1/userService/config"
	"github.com/jackc/pgx"
	"os"
)

type db struct {
	dbConn *pgx.ConnPool
}

func getDBConnStr() string {
	conf := config.GetNewConfig()
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Password, conf.Database)
}

func GetNewDB() *pgx.Conn {
	conf := config.NewConfig
	conn, err := pgx.Connect(pgx.ConnConfig{
		Host:     conf.Host,
		Port:     conf.Port,
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
