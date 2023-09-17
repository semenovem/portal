package people

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestValidateUserName(t *testing.T) {
	t.Parallel()

	t.Run("valid", func(t *testing.T) {
		cases := []string{
			"sdfsadfsa",
			"34234",
			"Фио",
		}

		for _, v := range cases {
			assert.NoError(t, ValidateUserName(v), "username:[%s]", v)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		casesInvalid := []string{
			"",
			"1",
			"   a ",
			strings.Repeat("a", 129),
		}

		for _, v := range casesInvalid {
			assert.Error(t, ValidateUserName(v), "username:[%s]", v)
		}
	})
}
