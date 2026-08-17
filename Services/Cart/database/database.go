package database

import (
	"cartservice/internal/config"
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDb( cfg config.Config) *sql.DB {
	db,err:= sql.Open(cfg.DB,os.Getenv("DB_URI"))

	if err!= nil {
		log.Fatal("Error to connect the db ",err)
	}

    if	err:=db.Ping(); err!=nil {
		log.Fatal("Error to connect the db",err)
	}

	return db

}