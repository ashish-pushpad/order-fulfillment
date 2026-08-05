package postgre

import (
	"database/sql"
	"log/slog"
	"checkout/internal/config"
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