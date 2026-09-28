package grpcclient

import (
	"apigateway/internal/config"
	// "log/slog"
	"proto/warehouse"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func WarehouseClient(cfg config.Config) (warehousepb.WarehouseProtoClient,error) {
	conn,err:=grpc.NewClient(
		cfg.WarehouseClient,
		grpc.WithTransportCredentials( insecure.NewCredentials() ),
	)
	if err!=nil {
		return nil,err
	}
	return warehousepb.NewWarehouseProtoClient(conn),nil
}