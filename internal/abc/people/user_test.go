package people

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
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

func Test1232(t *testing.T) {
	t.Parallel()

	id := uuid.New()

	fmt.Println(">>>>>>>>>> uuid=", id.String())

	str := base64.URLEncoding.EncodeToString(id[:])
	fmt.Println(">>>>>>>>>> uuid str=", str)
	fmt.Println(">>>>>>>>>> uuid len(str) = ", len(id.String()))
	fmt.Println(">>>>>>>>>> uuid len(str) = ", len(str))

	dec, err := base64.URLEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("[ERR] >> ", err)
		return
	}

	id2, err := uuid.FromBytes(dec)
	if err != nil {
		fmt.Println("[ERR] 222 >> ", err)
		return
	}

	fmt.Println(">>>>>>>>>> uuid=", id2.String())
}
