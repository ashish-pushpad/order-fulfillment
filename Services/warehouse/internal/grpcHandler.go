package internal

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"log/slog"
	pb "proto/warehouse"
	// "strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcHandler struct {
	pb.UnimplementedWarehouseProtoServer
	service *Service
}

func NewGrpcHandler(service *Service) *GrpcHandler {
	return &GrpcHandler{
		service: service,
	}
}

func (h *GrpcHandler) GetWarehouse(ctx context.Context, req *pb.GetWarehouseRequest) (*pb.GetWarehouseResponse, error) {
	resp, err := h.service.GetWarehouse(ctx, req.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "warehouse not found")
		}
		log.Println("error to getWArehouse", err)
		return nil, status.Error(codes.Internal, "internal server error")
	}

	return &pb.GetWarehouseResponse{
		Id:       resp.Id,
		Name:     resp.Name,
		City:     resp.City,
		Addressd: resp.Address,
	}, nil
}

func (h *GrpcHandler) CreateWarehouse(c context.Context, req *pb.CreateWarehouseRequest) (*emptypb.Empty, error) {

	reqBody := CreateWarehouseBody{
		Name:    req.Name,
		City:    req.City,
		Address: req.Address,
	}

	err := h.service.CreateWarehouse(reqBody)
	if err != nil {
		slog.Error("error to add the Warehouse ", "error", err)

		return nil,err
	}

	return &emptypb.Empty{},nil

}

// Id      int64  `json:"id"`
// Name    string `json:"name" binding:"required"`
// City    string `json:"city" binding:"required"`
// Address string `json:"address" binding:"required"`

func (h *GrpcHandler) GetWarehouses(c context.Context, req *pb.GetWarehousesRequest) (*pb.GetWarehousesResponse, error) {

	var (
		err   error
	)
	pageInt :=int(req.Page)
	limitInt := int(req.Limit)

	if pageInt == 0 || limitInt == 0{
		return nil, errors.New("page or limit not provided")
	}


	

	Warehouse, err := h.service.GetWarehouses(&limitInt, &pageInt)
	if err != nil {
		slog.Error("error to add the Warehouse ", "error", err)
		return nil, err
	}

	response := &pb.GetWarehousesResponse{}

	for _, warehouse := range Warehouse {
		response.Warehouses = append(
			response.Warehouses,
			&pb.GetWarehouseResponse{
				Id:      warehouse.Id,
				Name:    warehouse.Name,
				City:    warehouse.City,
				Addressd: warehouse.Address,
			},
		)
	}

	return response, nil

}

// func (h *GrpcHandler) UpdateWarehouse(c context.Context, req *pb.U) {

// 	idParam := c.Param("id")

// 	id, err := strconv.ParseInt(idParam, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid Warehouse id",
// 		})
// 		return
// 	}

// 	var data UpdateWarehouseBody
// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	if data.Name == "" || data.City == "" || data.Address == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "all fields are required",
// 		})
// 		return
// 	}

// 	updatedWarehouse, err := h.service.UpdateWarehouse(id, data)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, updatedWarehouse)
// }

func (h *GrpcHandler) DeleteWarehouse(c context.Context , req *pb.DeleteWarehouseRequest) (*pb.DeleteWarehouseResponse,error) {
	idParam := req.Id

	// id, err := strconv.ParseInt(idParam, 10, 64)
	// if err != nil {
	// 	// c.JSON(http.StatusBadRequest, gin.H{
	// 	// 	"error": "invalid Warehouse id",
	// 	// })
	// 	slog.Error("Error to conver the Id",err)
	// 	return nil,status.Error(codes.InvalidArgument,"invalid Warehouse id")
	// }

	deletedWarehouse, err := h.service.DeleteWarehouse(idParam)
	if err!= nil {		
		return nil, status.Errorf(codes.Internal, "failed to delete warehouse: %v", err)
	}

	// c.JSON(http.StatusOK, gin.H{
	// 	"message":   "Warehouse deleted successfully",
	// 	"Warehouse": deletedWarehouse,
	// })
	return  &pb.DeleteWarehouseResponse{
		Id: deletedWarehouse,
	},nil
}
