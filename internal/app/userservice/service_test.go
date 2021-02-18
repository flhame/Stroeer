package userservice_test

import (
	"Strooer/internal/app/mock"
	"Strooer/internal/app/model"
	"Strooer/internal/app/userservice"
	"github.com/stretchr/testify/assert"
	"testing"
)

var us userservice.UserService

func init() {
	us = userservice.UserService{
		UserConnection: &mock.ApiMock{},
	}
}

func TestApplication_GetUsersWithComments(t *testing.T) {
	t.Run("Get all users with their comments", func(t *testing.T) {
		result, err := us.GetUsersWithComments(nil)
		assert.Nil(t, err)
		assert.NotNil(t, result)

		var expected []model.UserAndComments
		for _, user := range mock.UserSeed {
			var matchingComments []model.Comment
			for _, comment := range mock.CommentSeed {
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
			expected = append(expected, combi)
		}

		assert.Equal(t, expected, result)


	})

	t.Run("Get single user with its comments", func(t *testing.T) {
		result, err := us.GetUsersWithComments(&mock.UserSeed[0].Id)
		assert.NotNil(t, result)
		assert.Nil(t, err)

		var expectedComments []model.Comment
		for _, comment := range mock.CommentSeed {
			if comment.UserId == mock.UserSeed[0].Id {
				expectedComments = append(expectedComments, comment)
			}
		}
		expected := []model.UserAndComments{
			{
				User: mock.UserSeed[0],
				Comments: expectedComments,
			},
		}
		assert.Equal(t, expected, result)

	})
}
