package graph


import "apigateway/proto/userpb"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct{
	UserClient userpb.UserServiceClient
}
