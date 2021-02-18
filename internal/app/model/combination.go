package model

type (
	UserAndComments struct {
		User     User      `json:"user"`
		Comments []Comment `json:"comments"`
	}
)
