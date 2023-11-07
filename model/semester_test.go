package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildUpdateFields(t *testing.T) {
	assertion := assert.New(t)

	semester := Semester{
		MinCredits: 8,
	}

	result := semester.BuildUpdateFields()

	assertion.NotNil(result)
	expected := map[string]interface{}{
		"MinCredits": 8,
	}
	assertion.Equal(expected, result)
}
