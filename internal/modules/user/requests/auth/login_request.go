package auth

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email,min=5,max=50"`
	Password string `form:"password" binding:"required,min=8"`
}
