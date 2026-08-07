package internal

// import (
// 	pb "checkout/proto/checkoutpb"
// 	"context"
// )

// type GrpcHandler struct {
// 	pb.UnimplementedAddressServiceServer
// 	service *Service
// }


// func NewGrpcHandler(service *Service ) *GrpcHandler{
// 	return &GrpcHandler{
// 		service: service,
// 	}
// }


// func (h *GrpcHandler) GetAddress (ctx context.Context,req *pb.GetAddressRequest) (*pb.GetAddressResponse,error){
// 	res,err:=h.service.GetAddress(ctx,req.UserId,req.AddressId)

// 	if err!=nil {
// 		return  nil,err
// 	}

// 	return  &pb.GetAddressResponse{
// 		 Id:res.Id,
// 		 FullName: res.FullName,
// 		 PhoneNumber: res.PhoneNumber,
// 		 HouseNo: res.HouseNo,
// 		 Apartment: res.Apartment,
// 		 Colony: res.Colony,
// 		 City: res.City,
// 		 State: res.State,
// 		 UserId: res.UserId,
// 		 PinCode: res.PinCode,
// 		 IsDefault: res.IsDefault,
// 	},nil
// }