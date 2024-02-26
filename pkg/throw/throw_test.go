package throw

import (
	"errors"
	"github.com/semenovem/portal/pkg/throw/throwtrace"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestDesc(t *testing.T) {
	t.Parallel()

	err := func() error {
		return NewBadRequest("test_err").
			SetDesc("desc_test %s %s", "arg_for_1", "arg_for_2")
	}()

	th := Cast(err)

	assert.Equal(t, "test_err", th.msg)
	assert.Equal(t, "desc_test %s %s", th.desc)
	assert.True(t, reflect.DeepEqual(th.descArgs, []any{"arg_for_1", "arg_for_2"}))
	assert.Equal(t, Kind(BadRequest), th.kind)
}

func TestThrowSetDesc(t *testing.T) {
	t.Parallel()

	th := &Throw{}

	th = th.SetDesc("tetwt", "arg2")

	str, arg := th.getDesc()
	assert.Equal(t, "tetwt", str)
	require.Len(t, arg, 1)
	assert.Equal(t, "arg2", arg[0])
}

func TestTrace(t *testing.T) {
	t.Parallel()

	err3 := func() error {
		err := func() error {
			return NewBadRequest("test_err").
				AddTrace("func1", map[string]any{"arg1": "3333"})
		}()

		if err != nil {
			return Cast(err).AddTrace("func2", map[string]any{"arg2": "222"})
		}

		return nil
	}()

	//fmt.Println("format = ", throwtrace.Format(err3))

	t.Run("login", func(t *testing.T) {
		var (
			//exp = map[string]any{
			//	"arg1":         "3333",
			//	"arg2":         "222",
			//	"trace_points": "((0))func1[rec/throw_test.go:45] <- ((1))func2[rec/throw_test.go:49]]",
			//}

			res = throwtrace.Format(err3)
		)

		assert.Equal(t, res["arg1"], "3333")
		assert.Equal(t, res["arg2"], "222")

		_, ok := res["trace_points"]
		assert.True(t, ok)
	})
}

func TestThrow_Error(t *testing.T) {
	tests := []struct {
		name string
		rec  *Throw
		want string
	}{
		{
			name: "NilThrow",
			rec:  nil,
			want: "",
		},
		{
			name: "NonNilThrow",
			rec: &Throw{
				msg: "test error",
			},
			want: "test error",
		},
	}
	for _, tt := range tests {
		got := tt.rec.Error()
		assert.Equalf(t, tt.want, got, "Case: %s | %v.Error()", tt.name, tt.rec)
	}
}

func TestThrow_Kind(t *testing.T) {
	tests := []struct {
		name string
		rec  *Throw
		want Kind
	}{
		{
			name: "NilThrow",
			rec:  nil,
			want: Unknown,
		},
		{
			name: "DenyKind",
			rec:  &Throw{kind: Deny},
			want: Deny,
		},
		{
			name: "DuplicateKind",
			rec:  &Throw{kind: Duplicate},
			want: Duplicate,
		},
		{
			name: "NotFoundKind",
			rec:  &Throw{kind: NotFound},
			want: NotFound,
		},
		{
			name: "BadRequestKind",
			rec:  &Throw{kind: BadRequest},
			want: BadRequest,
		},
		{
			name: "AuthKind",
			rec:  &Throw{kind: Auth},
			want: Auth,
		},
	}
	for _, tt := range tests {
		got := tt.rec.Kind()
		assert.Equalf(t, tt.want, got, "Case: %s | %v.Kind()", tt.name, tt.rec)
	}
}

func TestThrow_getDesc(t *testing.T) {
	tests := []struct {
		name         string
		rec          *Throw
		wantDesc     string
		wantDescArgs []any
	}{
		{
			name:         "NilThrow",
			rec:          nil,
			wantDesc:     "",
			wantDescArgs: nil,
		},
		{
			name:         "EmptyDescAndArgs",
			rec:          &Throw{desc: "", descArgs: nil},
			wantDesc:     "",
			wantDescArgs: nil,
		},
		{
			name:         "NonEmptyDescEmptyArgs",
			rec:          &Throw{desc: "test error", descArgs: nil},
			wantDesc:     "test error",
			wantDescArgs: nil,
		},
		{
			name:         "NonEmptyDescNonEmptyArgs",
			rec:          &Throw{desc: "test error", descArgs: []any{"arg1", "arg2"}},
			wantDesc:     "test error",
			wantDescArgs: []any{"arg1", "arg2"},
		},
	}
	for _, tt := range tests {
		gotDesc, gotDescArgs := tt.rec.getDesc()
		assert.Equalf(t, tt.wantDesc, gotDesc, "Case: %s | Desc mismatch", tt.name)
		assert.Equalf(t, tt.wantDescArgs, gotDescArgs, "Case: %s | DescArgs mismatch", tt.name)
	}
}

func TestThrow_SetDesc(t *testing.T) {
	tests := []struct {
		name     string
		rec      *Throw
		desc     any
		descArgs []any
		want     *Throw
	}{
		{
			name:     "NilThrow",
			rec:      nil,
			desc:     "test error",
			descArgs: []any{"arg1", "arg2"},
			want:     &Throw{kind: Unknown, msg: "", desc: "test error", descArgs: []any{"arg1", "arg2"}},
		},
		//{
		//	name:     "NonNilThrow",
		//	rec:      &Throw{kind: Deny, msg: "test"},
		//	desc:     "new error desc",
		//	descArgs: []any{"newArg1", "newArg2"},
		//	want:     &Throw{kind: Deny, msg: "test", desc: "new error desc", descArgs: []any{"newArg1", "newArg2"}},
		//},
	}
	for _, tt := range tests {
		got := tt.rec.SetDesc(tt.desc, tt.descArgs...)
		assert.Equalf(t, tt.want.desc, got.desc, "Case: %s | Desc mismatch", tt.name)
		assert.ElementsMatchf(t, tt.want.descArgs, got.descArgs, "Case: %s | DescArgs mismatch", tt.name)
	}
}

func TestThrow_traceData(t *testing.T) {
	tracePoints := []*throwtrace.Point{
		{Name: "methodA", LineCode: "path/to/file.go:123"},
		{Name: "methodB", LineCode: "path/to/file.go:456"},
	}
	with := map[string]any{"key": "value"}
	tests := []struct {
		name   string
		rec    *Throw
		wantTP []*throwtrace.Point
		wantW  map[string]any
	}{
		{
			name:   "NilThrow",
			rec:    nil,
			wantTP: nil,
			wantW:  nil,
		},
		{
			name:   "NonNilThrow",
			rec:    &Throw{kind: Deny, msg: "test", tracePoints: tracePoints, with: with},
			wantTP: tracePoints,
			wantW:  with,
		},
	}
	for _, tt := range tests {
		gotTracePoints, gotWith := tt.rec.traceData()
		if tt.rec == nil {
			assert.Nilf(t, gotTracePoints, "Case: %s | Expected nil TracePoints", tt.name)
			assert.Nilf(t, gotWith, "Case: %s | Expected nil With", tt.name)
		} else {
			assert.Equalf(t, tt.wantTP, gotTracePoints, "Case: %s | TracePoints mismatch", tt.name)
			assert.Equalf(t, tt.wantW, gotWith, "Case: %s | With mismatch", tt.name)
		}
	}
}
func TestThrow_AddWith(t *testing.T) {
	tests := []struct {
		name string
		rec  *Throw
		key  string
		val  any
		want map[string]any
	}{
		{
			name: "NilWith",
			rec:  &Throw{with: nil},
			key:  "key1",
			val:  "value1",
			want: map[string]any{"key1": "value1"},
		},
		{
			name: "EmptyWith",
			rec:  &Throw{with: make(map[string]any)},
			key:  "key1",
			val:  "value1",
			want: map[string]any{"key1": "value1"},
		},
		{
			name: "ExistingWith",
			rec:  &Throw{with: map[string]any{"key2": "value2"}},
			key:  "key3",
			val:  "value3",
			want: map[string]any{"key2": "value2", "key3": "value3"},
		},
	}
	for _, tt := range tests {
		tt.rec.addWith(tt.key, tt.val)
		if !reflect.DeepEqual(tt.rec.with, tt.want) {
			t.Errorf("With mismatch. Got = %+v, want = %+v", tt.rec.with, tt.want)
		}
	}
}
func TestThrow_AddTrace(t *testing.T) {
	tests := []struct {
		name  string
		th    *Throw
		trace string
		with  map[string]any
		want  []*throwtrace.Point
	}{
		{
			name:  "NilThrow",
			th:    nil,
			trace: "test",
			with: map[string]any{
				"key1": "value1",
			},
			want: []*throwtrace.Point{
				{
					Name: "test",
				},
			},
		},
		{
			name:  "EmptyThrow",
			th:    &Throw{},
			trace: "point1",
			with: map[string]any{
				"key1": "value1",
			},
			want: []*throwtrace.Point{
				{
					Name: "point1",
				},
			},
		},
		{
			name:  "ThrowWithExistingTracePoints",
			th:    &Throw{tracePoints: []*throwtrace.Point{{Name: "existingTracePoint"}}},
			trace: "newTrace",
			with: map[string]any{
				"key1": "value1",
			},
			want: []*throwtrace.Point{
				{
					Name: "existingTracePoint",
				},
				{
					Name: "point1",
				},
			},
		},
	}

	for _, tt := range tests {
		th := tt.th.AddTrace(tt.trace, tt.with)

		if th == nil {
			if tt.want != nil {
				t.Errorf("Case: %s | Expected: nil, Got:%v", tt.name, th.tracePoints)
			}
			continue
		}

		if len(tt.want) != len(th.tracePoints) {
			t.Errorf("Case: %s | Expected: %v, Got:%v", tt.name, tt.want, th.tracePoints)
			continue
		}

		for i, v := range tt.want {
			if v.Name != tt.want[i].Name {
				t.Errorf("Case: %s | Expected: %v, Got:%v", tt.name, tt.want, th.tracePoints)
			}
		}
	}
}

func TestThrow_new(t *testing.T) {
	type args struct {
		k   Kind
		msg any
	}
	tests := []struct {
		name string
		args args
		want *Throw
	}{
		{
			name: "Error kind and error msg",
			args: args{k: Deny, msg: errors.New("test error")},
			want: &Throw{kind: Deny, msg: "test error"},
		},
		{
			name: "Error kind and string msg",
			args: args{k: NotFound, msg: "test error"},
			want: &Throw{kind: NotFound, msg: "test error"},
		},
		{
			name: "Error kind and integer msg",
			args: args{k: NotFound, msg: 1},
			want: &Throw{kind: NotFound, msg: "1"},
		},
		{
			name: "Non error kind and error msg",
			args: args{k: Unknown, msg: errors.New("test error")},
			want: &Throw{kind: Unknown, msg: "test error"},
		},
	}
	for _, tt := range tests {
		got := newThrow(tt.args.k, tt.args.msg)
		assert.Equalf(t, tt.want, got, "newThrow(%v, %v)", tt.args.k, tt.args.msg)
	}
}
