package users

import (
	"context"
	userHandler "github.com/phuc-create/go-simple-crud/internal/handlers/users"
	"github.com/phuc-create/go-simple-crud/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestImplement_CreateUser(t *testing.T) {
	ctx := context.Background()
	i := implement{}

	datetime := time.Date(2023, 04, 19, 0, 0, 0, 0, time.UTC)

	type expected struct {
		user  models.User
		error error
	}

	type arg struct {
		input          models.User
		expectedResult expected
	}

	tsc := map[string]arg{
		"success: create users successfully": {

			input: models.User{
				ID:        "1",
				Username:  "sam",
				Password:  "sam",
				CreatedAt: datetime,
				UpdatedAt: datetime,
			},
			expectedResult: expected{
				user: models.User{
					ID:        "1",
					Username:  "sam",
					Password:  "sam",
					CreatedAt: datetime,
					UpdatedAt: datetime,
				},
				error: nil,
			},
		},
		"fail: error missing information": {
			input: models.User{
				ID:        "",
				Username:  "sam",
				Password:  "sam",
				CreatedAt: datetime,
				UpdatedAt: datetime,
			},
			expectedResult: expected{
				user:  models.User{},
				error: userHandler.ErrMissingInformation,
			},
		},
		"fail: error users already exist": {
			input: models.User{
				ID:        "1",
				Username:  "sam",
				Password:  "sam",
				CreatedAt: datetime,
				UpdatedAt: datetime,
			},
			expectedResult: expected{
				user:  models.User{},
				error: userHandler.ErrUserAlreadyExist,
			},
		},
		"fail: error password contain white space": {
			input: models.User{
				ID:        "1",
				Username:  "sam2",
				Password:  "sam sam",
				CreatedAt: datetime,
				UpdatedAt: datetime,
			},
			expectedResult: expected{
				user:  models.User{},
				error: userHandler.PasswordContainWhiteSpace,
			},
		},
	}

	for _, tc := range tsc {
		usr, err := i.CreateUserController(ctx, tc.input)
		if err != nil {
			require.EqualError(t, err, tc.expectedResult.error.Error())
		}
		assert.Equal(t, tc.expectedResult.user, usr)
	}

}
