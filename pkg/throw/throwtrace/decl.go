package throwtrace

import (
	"fmt"
	"strings"
)

const (
	tracePointsName = "trace_points"
)

type (
	Point struct {
		Name     string
		LineCode string
	}
	Reason struct {
		Reason   string
		Metadata map[string]string
	}
)

// Format выдаст данные для добавления в logger
// trace_points: "(1)Get[provider/user.go:21] <-<- (2)userRepo.Get[action/user.go:24] <-<- (3)Get.Provider"
// для добавленных полей
// name1: value1,
// name2: value2,
// name3: value3,
func Format(err error) map[string]any {
	var (
		points, with = extractTraceData(err)
		m            = make(map[string]any)
	)

	for k, v := range with {
		m[k] = v
	}

	m[tracePointsName] = formatPoints(points)

	return m
}

func formatPoints(ps []*Point) string {
	s := make([]string, 0, len(ps))

	for i, v := range ps {
		s = append(s, fmt.Sprintf("((%d))%s[%s]", i, v.Name, v.LineCode))
	}

	return strings.Join(s, " <-<- ")
}

func Desc(err error) (string, []any) {
	return extractDesc(err)
}

func Reasons(err error) []*Reason {
	return extractReasons(err)
}
