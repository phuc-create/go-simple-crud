package controllers

import (
	"github.com/phuc-create/go-simple-crud/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"time"

	"testing"
)

func TestImplement_CreateUser(t *testing.T) {
	i := implement{}

	type expected struct {
		user  UserResponse
		error error
	}

	type arg struct {
		input          models.User
		expectedResult expected
		expectedError  error
	}
	datetime := time.Date(2023, 04, 19, 0, 0, 0, 0, time.UTC)

	tsc := []arg{
		{
			input: models.User{
				ID:        "1",
				Username:  "sam",
				Password:  "sam",
				CreatedAt: datetime,
				UpdatedAt: datetime,
			},
			expectedResult: expected{
				user: UserResponse{
					ID:        "1",
					Username:  "sam",
					Password:  "sam",
					CreatedAt: datetime,
					UpdatedAt: datetime,
				},
				error: nil,
			},
			expectedError: nil,
		},
	}

	for _, tc := range tsc {
		usr, err := i.CreateUser(tc.input)
		if err != nil {
			require.EqualError(t, err, tc.expectedResult.error.Error())
		}
		assert.Equal(t, tc.expectedResult.user, usr)
	}
}
