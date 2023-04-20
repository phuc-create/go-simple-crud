package user

import (
	"github.com/phuc-create/go-simple-crud/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"time"

	"testing"
)

func TestImplement_CreateUser(t *testing.T) {
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
		"success: create user successfully": {

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
				error: ErrMissingInformation,
			},
		},
		"fail: error user already exist": {
			input: models.User{
				ID:        "1",
				Username:  "sam",
				Password:  "sam",
				CreatedAt: datetime,
				UpdatedAt: datetime,
			},
			expectedResult: expected{
				user:  models.User{},
				error: ErrUserAlreadyExist,
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
				error: ErrPasswordContainWhiteSpace,
			},
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
