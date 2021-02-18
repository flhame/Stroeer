package userservice

import (
	"Strooer/internal/app/model"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

const (
	UserEndpoint    = "https://jsonplaceholder.typicode.com/users"
	CommentEndpoint = "https://jsonplaceholder.typicode.com/posts"
)

func (us *UserService) GetUsersWithComments(userId *int) ([]model.UserAndComments, error) {
	userChan := make(chan []model.User)
	commentChan := make(chan []model.Comment)

	go us.FetchUsers(userId, userChan)
	go us.FetchComments(userId, commentChan)

	// wait for the go routines sending value
	users := <-userChan
	comments := <-commentChan

	if users == nil || comments == nil {
		return nil, errors.New("failed fetching users or comments")
	}

	combinations := us.MergeUserAndComments(users, comments)
	return combinations, nil
}

func (us *UserService) MergeUserAndComments(users []model.User, comments []model.Comment) []model.UserAndComments {
	var combinations []model.UserAndComments
	for _, user := range users {
		var matchingComments []model.Comment
		for _, comment := range comments {
			if comment.UserId == user.Id {
				// adding comments to slice
				matchingComments = append(matchingComments, comment)
			}
		}
		// adding completed combinations to slice
		combi := model.UserAndComments{
			User:     user,
			Comments: matchingComments,
		}
		combinations = append(combinations, combi)
	}

	return combinations
}

func (us *UserService) FetchUsers(userId *int, channel chan []model.User) {
	log.Printf("NOTICE start fetching users")
	endPoint := UserEndpoint
	if userId != nil && *userId > 0 {
		endPoint += fmt.Sprintf("/%d", userId)
	}

	req, err := http.NewRequest(http.MethodGet, endPoint, nil)
	if err != nil {
		log.Printf("ERROR: could not create request for fetching users with error: %v", err)
		channel <- nil
		return
	}

	resp, err := us.DoRequest(req)
	if err != nil {
		log.Printf("ERROR: could not fetching users with error: %v", err)
		channel <- nil
		return
	}

	users := []model.User{}
	err = json.Unmarshal(resp, &users)
	if err != nil {
		log.Printf("ERROR: failed to unmarshal response from getting user with err: %v", err)
		channel <- nil
		return
	}

	log.Printf("NOTICE successfully fetched users")
	channel <- users
}

func (us *UserService) FetchComments(userId *int, channel chan []model.Comment) {
	log.Printf("NOTICE start fetching comments")

	endPoint := CommentEndpoint
	if userId != nil && *userId > 0 {
		endPoint += fmt.Sprintf("?userId=%d", userId)
	}

	req, err := http.NewRequest(http.MethodGet, endPoint, nil)
	if err != nil {
		log.Printf("ERROR: could not create request for fetching comments with error: %v", err)
		channel <- nil
		return
	}

	resp, err := us.DoRequest(req)
	if err != nil {
		log.Printf("ERROR: could not fetching comments with error: %v", err)
		channel <- nil
		return
	}

	comments := []model.Comment{}
	err = json.Unmarshal(resp, &comments)
	if err != nil {
		log.Printf("ERROR: failed to unmarshal response from getting comments with err: %v", err)
		channel <- nil
		return
	}

	log.Printf("NOTICE successfully fetched comments")
	channel <- comments
}

func (us *UserService) DoRequest(req *http.Request) ([]byte, error) {
	requestDump, _ := httputil.DumpRequest(req, true)
	log.Println(string(requestDump))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	log.Println(string(body))

	return body, nil
}

func (app *Application) NewUserService() *UserService {
	return &UserService{}
}
