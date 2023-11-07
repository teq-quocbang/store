package times

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLessThanThreeMonth(t *testing.T) {
	assertion := assert.New(t)

	// good case
	{
		start, err := StringToTime("2023-01-01T08:37:01+07:00")
		assertion.NoError(err)
		end, err := StringToTime("2023-04-01T08:37:01+07:00")
		assertion.NoError(err)

		ok := IsLessThan(start, end, ThreeMonth)
		assertion.False(ok)
	}

	// bad case
	{
		start, err := StringToTime("2023-01-01T08:37:01+07:00")
		assertion.NoError(err)
		end, err := StringToTime("2023-03-01T08:37:01+07:00")
		assertion.NoError(err)

		ok := IsLessThan(start, end, ThreeMonth)
		assertion.True(ok)
	}
}
