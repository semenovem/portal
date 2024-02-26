package throwtrace

var (
	extractTraceData traceDataExtractor
	extractDesc      descExtractor
	extractReasons   reasonsExtractor
)

type (
	traceDataExtractor func(err error) ([]*Point, map[string]any)
	descExtractor      func(err error) (string, []any)
	reasonsExtractor   func(err error) []*Reason

	XDoNotUse struct{}
)

func (_ XDoNotUse) SetExtractTraceData(f traceDataExtractor) {
	extractTraceData = f
}

func (_ XDoNotUse) SetExtractDesc(f descExtractor) {
	extractDesc = f
}

func (_ XDoNotUse) SetExtractReasons(f reasonsExtractor) {
	extractReasons = f
}
