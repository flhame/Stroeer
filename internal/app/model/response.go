package model

// File for describing responses send by application

type HealthResponse struct {
	Message string `json:"message"`
}

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UsersResponse struct {
	Users []User `json:"users"`
}

type CommentsResponse struct {
	Comments []Comment `json:"comments"`
}