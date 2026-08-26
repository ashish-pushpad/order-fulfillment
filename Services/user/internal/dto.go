package internal


type CreateUserBody struct{
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}


type GetLoginRequest struct {
	Email string 	`json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetLoginResponse struct {
	Id int64 	`json:"id" binding:"required"`
	Role string `json:"string" binding:"required"`
}

type GetUserResponse struct {
	Id int64 	`json:"id"`
	Name string `json:"name" `
	Email string `json:"email"`
}

type UpdateUserBody struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}


type UpdateUserResponse struct{
	ID    int    `json:"id"`
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Role string `json:"role"`
}




type UserResponse struct{
	ID    int    `json:"id"`
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Role string `json:"role"`
}