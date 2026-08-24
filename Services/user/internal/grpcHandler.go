package internal

import (
	"context"
	"log"
	pb "user/proto/userpb"
)

	type GrpchHandler struct {
		pb.UnimplementedUserServiceServer
		service *Service
	}

	func NewGrpcHandler (service *Service) *GrpchHandler{
		return &GrpchHandler{
			service: service,
		}
	}


func (h *GrpchHandler) GetUserById(ctx context.Context,req *pb.GetUserByIdRequest)(*pb.GetUserByIdResponse,error){

		user,err:=h.service.GetUser(ctx,req.Id)

		if err!=nil {
			log.Println("error to get the user ",err)
			return  &pb.GetUserByIdResponse{},err
		}
		return  &pb.GetUserByIdResponse{
			Id:user.Id,
			Name:user.Name,
			Email:user.Email,
		},nil
}