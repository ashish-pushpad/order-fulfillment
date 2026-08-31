package internal

import (
	"context"
	"log"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "proto/user"
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

func (h *GrpchHandler) CreateUser(ctx context.Context,data *pb.CreateUserRequest) (*emptypb.Empty,error){
	 userData :=CreateUserBody{
		data.User.Email,
		data.User.Name,
		data.User.Password,
	}

	err:=h.service.CreateUser(ctx,userData)
	if err!=nil {
		return  &emptypb.Empty{},err
	}
	return  &emptypb.Empty{},nil
}

func (h *GrpchHandler) UpdateUser(ctx context.Context, data *pb.UpdateUserRequest) (*pb.UpdateUserResponse,error){
		var updatData=UpdateUserBody{
			Email:data.Data.Email,
			Name: data.Data.Name,
			Password:data.Data.Password,
		}
		// log.Println("data for updateUser",data.Data)
		res,err:=h.service.UpdateUser(ctx,data.Id,updatData)
		if err!=nil {
			return  &pb.UpdateUserResponse{},err
		}
		return  &pb.UpdateUserResponse{
			Id: int64(res.ID),
			Name: res.Name,
			Email: res.Email,
			Password: res.Role,
		},nil
}

func (h *GrpchHandler) DeleteUser(ctx context.Context, data *pb.DeleteUserRequest) (*pb.DeleteUserResponse,error){
	id:=data.Id
	userRes,err:=h.service.DeleteUser(ctx,id)
	if err!=nil{
		return  nil,err
	}
	return  &pb.DeleteUserResponse{
		Id: int64(userRes.ID),
		Name: userRes.Name,
		Email: userRes.Email,
		Password: userRes.Role,
	},nil
}

// func CreateUser ( ctx context.Context,)

	// GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	// CreateUser(context.Context, *CreateUserRequest) (*emptypb.Empty, error)
	// UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	// DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
