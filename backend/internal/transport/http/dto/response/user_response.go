package response

type UserResponse struct {
	ID        string  `json:"id"`
	UserName  string  `json:"user_name"`
	UserLogin float32 `json:"user_login"`
	Image     string  `json:"image_url,omitempty"`
}
