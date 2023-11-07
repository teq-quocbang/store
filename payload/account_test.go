package payload

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	assertion := assert.New(t)
	testUsername := "test_user"
	testPassword := "test_password"
	testEmail := "test@teqnological.asia"

	// good case
	{
		req := SignUpRequest{
			Username: testUsername,
			Password: testPassword,
			Email:    testEmail,
		}

		err := req.Validate()

		assertion.NoError(err)
	}

	// bad cases
	{ // missing user name
		requestMissingUsername := SignUpRequest{
			Password: testPassword,
			Email:    testEmail,
		}

		err := requestMissingUsername.Validate()

		assertion.Error(err)
		expected := "Key: 'SignUpRequest.Username' Error:Field validation for 'Username' failed on the 'required' tag"
		assertion.Equal(expected, err.Error())
	}
	{ // password short than 8 char
		requestPasswordShort := SignUpRequest{
			Username: testUsername,
			Password: "bad",
			Email:    testEmail,
		}

		err := requestPasswordShort.Validate()

		assertion.Error(err)
		expected := "Key: 'SignUpRequest.Password' Error:Field validation for 'Password' failed on the 'min' tag"
		assertion.Equal(expected, err.Error())
	}
	{ // invalid email
		requestInvalidEmail := SignUpRequest{
			Username: testUsername,
			Password: testPassword,
			Email:    "bad_email",
		}

		err := requestInvalidEmail.Validate()

		assertion.Error(err)
		expected := "Key: 'SignUpRequest.Email' Error:Field validation for 'Email' failed on the 'email' tag"
		assertion.Equal(expected, err.Error())
	}
}
