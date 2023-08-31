package txt

import (
	"github.com/semenovem/portal/pkg/fail"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	EN = "en"
	RU = "ru"
)

var (
	regDig = regexp.MustCompile(`^[\s\d]+`)
)

func dig(s string) int {
	d, _ := strconv.Atoi(strings.TrimSpace(regDig.FindString(s)))
	return d
}

func text(s string) string {
	return regDig.ReplaceAllString(s, "")
}

func GetMessages() map[string]*fail.Message {
	res := make(map[string]*fail.Message)

	for key, val := range messages {
		res[key] = &fail.Message{
			Code:        dig(key),
			DefaultText: text(key),
			Translations: map[string]string{
				EN: val.en,
			},
		}
	}

	return res
}

func GetHTTPStatuses() map[int]*fail.Message {
	res := make(map[int]*fail.Message)

	for key, v := range statuses {
		res[key] = &fail.Message{
			Code:        key,
			DefaultText: v,
			Translations: map[string]string{
				RU: http.StatusText(key),
			},
		}
	}

	return res
}
