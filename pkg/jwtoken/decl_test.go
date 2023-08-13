package jwtoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_taskItem_timeHasCome(t *testing.T) {
	_, err := time.LoadLocation("UTC")
	if !assert.NoError(t, err) {
		return
	}

}
