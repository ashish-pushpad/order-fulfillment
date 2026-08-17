package postgre

import (
	"database/sql"
	"inventory/internal/config"
	"log/slog"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDb(cfg config.Config)(*sql.DB,error) {

	db,err:= sql.Open(cfg.Db,os.Getenv("DB_URI"))

	if err != nil {
		slog.Error("Error to connect the db ","error",err)
		return  nil,err
	}

	if err:=db.Ping();err!=nil{
		slog.Error("Error to connect the db ","error",err)
		return  nil,err
	}


	slog.Info("Database connected successfully")

	return db, nil


}