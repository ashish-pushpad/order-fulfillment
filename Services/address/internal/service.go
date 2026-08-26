package internal

import (
	pb "proto/user"
	"context"
	"errors"
	"log"
)

type Service struct {
	repo *Repository
	userClient pb.UserServiceClient
}

func NewService(repo *Repository  , userClient pb.UserServiceClient) *Service {
	return &Service{
		repo: repo,
		userClient: userClient,
	}
}

func (s *Service) CreateAddress(ctx context.Context, userID int64, address CreateAddressBody) error {

	// Business logic can go here
	// Example:
	// - Validate phone number
	// - Ensure only one default address
	// - Verify user exists

	resp,err:=s.userClient.GetUserById(ctx,&pb.GetUserByIdRequest{
		Id: userID,
	})

	if err!=nil{
		return  err
	}

	return s.repo.CreateAddress(resp.Id, address)
}

func (s *Service) GetAddress(ctx context.Context,userID int64,addressID int64,) (AddressResponse, error) {

	address, err := s.repo.GetAddress(ctx,addressID)  //s.repo.GetAddress(addressID)
	log.Println("Address User id",address.UserId)
	log.Println("given add id",addressID)
	if err != nil {
		return AddressResponse{}, err
	}

	if address.UserId != userID {
		return AddressResponse{},errors.New("address does not belong to user")
	}

	return address, nil
}

func (s *Service) GetAddresses(page, limit *int) ([]AddressResponse, error) {

	addresses, err := s.repo.GetAddresses(page, limit)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (s *Service) UpdateAddress(id int64, data UpdateAddressBody) (AddressResponse, error) {

	updatedAddress, err := s.repo.UpdateAddress(id, data)
	if err != nil {
		return AddressResponse{}, err
	}

	return updatedAddress, nil
}

func (s *Service) DeleteAddress(id int64) error {

	return s.repo.DeleteAddress(id)
}