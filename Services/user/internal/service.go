package internal

import "context"

// "errors"
// "user/auth"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateUser(ctx context.Context,user CreateUserBody) error {

	// Business rules

	// Example:
	// Validate name
	// Check duplicate email
	// Hash password
	// Send event

	return s.repo.CreateUser(ctx, user)
}

func (s *Service) GetUser(ctx context.Context,id int64) (GetUserResponse, error) {
	user, err := s.repo.GetUser(ctx,id)

	if err != nil {
		return GetUserResponse{}, err
	}

	return user, nil
}

func (s *Service) UpdateUser(ctx context.Context,id int64, data UpdateUserBody) (UpdateUserResponse, error) {

	user, err := s.repo.UpdateUser(ctx, id, data)

	if err != nil {
		return UpdateUserResponse{}, err
	}

	return user, nil

}

func (s *Service) DeleteUser(ctx context.Context, id int64) (UserResponse, error) {
	deletedUser, err := s.repo.DeleteUser(ctx,id)
	if err != nil {
		return UserResponse{}, err
	}

	return deletedUser, nil
}

// func (s *Service) LoginUser(email, password string) (string, error) {
// 	// verify email verify password helper
// 	user, err := s.repo.VerifyUser(email, password)

// 	if err != nil {
// 		return "", errors.New("Verificatio fail")
// 	}
// 	token,err:=auth.GenerateToken(user.Id,user.Role)
// 	if err!=nil {
// 		return  "",err
// 	}
// 	return  token ,nil
// }