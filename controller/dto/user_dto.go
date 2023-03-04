package dto

type UpdateUserRequest struct {
	Name     string `json:"name"`
}

type UpdateUserResponse struct {
	Id       string `json:"id"`
	Name		 string `json:"name"`
}