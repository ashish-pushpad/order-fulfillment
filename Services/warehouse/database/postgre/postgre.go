package postgre

import (
	"database/sql"
	"log/slog"
	"warehouse/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDb(cfg config.Config)(*sql.DB,error) {

	db,err:= sql.Open("pgx",cfg.DB_URI)

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