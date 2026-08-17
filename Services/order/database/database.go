package database

import (
	"database/sql"
	"log"
	"order"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDb( cfg order.Config) *sql.DB {
	db ,err:= sql.Open(cfg.DB,os.Getenv("DB_URI"))

	if err!= nil {
		log.Fatal("Error to connect the db ",err)
	}

    if	err:=db.Ping(); err!=nil {
		log.Fatal("Error to connect the db",err)
	}

	return db

}