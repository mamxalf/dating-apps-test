package dto

type UserResponse struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	IsVerified bool   `json:"is_verified"`
}
