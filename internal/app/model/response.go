package model

// File for describing responses send by application

type (
	HealthResponse struct {
		Message string `json:"message"`
	}

	ApiError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	UsersResponse struct {
		Users []User `json:"users"`
	}

	CommentsResponse struct {
		Comments []Comment `json:"comments"`
	}

	UsersAndCommentsResponse struct {
		Result []UserAndComments `json:"result"`
	}
)
