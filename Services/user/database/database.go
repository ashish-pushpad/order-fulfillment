package database

import (
	"database/sql"
	"log"
	"user/internal/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDb( cfg config.Config) *sql.DB {
	db ,err:= sql.Open("pgx",cfg.DB_URI)

	if err!= nil {
		log.Fatal("Error to connect the db ",err)
	}

    if	err:=db.Ping(); err!=nil {
		log.Fatal("Error to connect the db",err)
	}

	return db

}