package mock

import (
	"Strooer/internal/app/model"
	"log"
	"net/http"
)

type ApiMock struct{}
var (
	CommentSeed = []model.Comment{
	{
		UserId: 1,
		Id:     1,
		Title:  "sample",
		Body:   "some text",
	},
	{
		UserId: 1,
		Id:     2,
		Title:  "sample",
		Body:   "some text",
	},
	{
		UserId: 1,
		Id:     3,
		Title:  "sample",
		Body:   "some text",
	},
	{
		UserId: 1,
		Id:     4,
		Title:  "sample",
		Body:   "some text",
	},
	{
		UserId: 2,
		Id:     5,
		Title:  "sample",
		Body:   "some text",
	},
	{
		UserId: 2,
		Id:     6,
		Title:  "sample",
		Body:   "some text",
	},
}
	UserSeed = []model.User{
		{
			Id:       1,
			Name:     "Leanne Graham",
			Username: "Bret",
			Email:    "Sincere@april.biz",
			Address: model.Address{
				Street:  "Kulas Light",
				Suite:   "Apt. 556",
				City:    "Gwenborough",
				Zipcode: "92998-3874",
				Geographic: model.Geo{
					Latitude:  "-37.3159",
					Longitude: "81.1496",
				},
			},
			Phone:   "1-770-736-8031 x56442",
			Website: "hildegard.org",
			Company: model.Company{
				Name: "Romaguera-Crona",
				CatchPhrase: "Multi-layered client-server neural-net",
				Businesses: "harness real-time e-markets",
			},
		},
		{
			Id:       2,
			Name:     "Ervin Howell",
			Username: "Antonette",
			Email:    "Shanna@melissa.tv",
			Address: model.Address{
				Street:  "Victor Plain",
				Suite:   "Suite 879",
				City:    "Wisokyburgh",
				Zipcode: "90566-7771",
				Geographic: model.Geo{
					Latitude:  "-43.9509",
					Longitude: "-34.4618",
				},
			},
			Phone:   "010-692-6593 x09125",
			Website: "anastasia.net",
			Company: model.Company{
				Name: "Deckow-Crist",
				CatchPhrase: "Proactive didactic contingency",
				Businesses: "synergize scalable supply-chains",
			},
		},
	}
)
func (api *ApiMock) FetchComments(userId *int, channel chan []model.Comment) {
	log.Print("NOTICE: using mocked function for fetching comments")
	var comments []model.Comment

	if userId != nil && *userId > 0 {
		for _, comment := range CommentSeed {
			if comment.UserId == *userId {
				comments = append(comments, comment)
			}
		}
	} else {
		comments = CommentSeed
	}

	channel <- comments
}

func (api *ApiMock) FetchUsers(userId *int, channel chan []model.User) {
	log.Print("NOTICE: using mocked function for fetching users")
	var users []model.User

	if userId != nil && *userId > 0 {
		for _, user := range UserSeed {
			if user.Id == *userId {
				users = append(users, user)
			}
		}
	} else {
		users = UserSeed
	}

	channel <- users
}

func (api *ApiMock) DoRequest(req *http.Request) ([]byte, error) {
	log.Print("NOTICE: using mocked function for doing request")
	// currently no function needed for mock
	return nil, nil
}
