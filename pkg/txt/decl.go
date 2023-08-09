package txt

import (
	"github.com/semenovem/portal/pkg/failing"
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

func GetMessages() map[string]*failing.Message {
	res := make(map[string]*failing.Message)

	for key, val := range messages {
		res[key] = &failing.Message{
			Code:        dig(key),
			DefaultText: text(key),
			Translations: map[string]string{
				EN: val.en,
			},
		}
	}

	return res
}

func GetHTTPStatuses() map[int]*failing.Message {
	res := make(map[int]*failing.Message)

	for key, v := range statuses {
		res[key] = &failing.Message{
			Code:        key,
			DefaultText: v,
			Translations: map[string]string{
				RU: http.StatusText(key),
			},
		}
	}

	return res
}
