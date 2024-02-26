package throw

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		msgOrErr interface{}
	}
	tests := []struct {
		name string
		args args
		want *Throw
	}{
		{
			name: "Error arg",
			args: args{msgOrErr: errors.New("test error")},
			want: &Throw{kind: Unknown, msg: "test error"},
		},
		{
			name: "String arg",
			args: args{msgOrErr: "test error"},
			want: &Throw{kind: Unknown, msg: "test error"},
		},
		{
			name: "Integer arg",
			args: args{msgOrErr: 1},
			want: &Throw{kind: Unknown, msg: "1"},
		},
	}
	for _, tt := range tests {
		got := New(tt.args.msgOrErr)
		assert.Equalf(t, tt.want, got, "Case: %s | NewUnknown() = %v", tt.name, got)
	}
}

func TestNewDeny(t *testing.T) {
	type args struct {
		msgOrErr any
	}
	tests := []struct {
		name string
		args args
		want *Throw
	}{
		{
			name: "Error arg",
			args: args{msgOrErr: errors.New("test error")},
			want: &Throw{kind: Deny, msg: "test error"},
		},
		{
			name: "String arg",
			args: args{msgOrErr: "test error"},
			want: &Throw{kind: Deny, msg: "test error"},
		},
		{
			name: "Integer arg",
			args: args{msgOrErr: 1},
			want: &Throw{kind: Deny, msg: "1"},
		},
	}
	for _, tt := range tests {
		assert.Equalf(t, tt.want, NewDeny(tt.args.msgOrErr), "NewDeny(%v)", tt.args.msgOrErr)
	}
}

func TestNewDuplicate(t *testing.T) {
	type args struct {
		msgOrErr any
	}
	tests := []struct {
		name string
		args args
		want *Throw
	}{
		{
			name: "Error arg",
			args: args{msgOrErr: errors.New("test case 1")},
			want: &Throw{kind: Duplicate, msg: "test case 1"},
		},
		{
			name: "String arg",
			args: args{msgOrErr: "test case 2"},
			want: &Throw{kind: Duplicate, msg: "test case 2"},
		},
		{
			name: "Integer arg",
			args: args{msgOrErr: 3},
			want: &Throw{kind: Duplicate, msg: "3"},
		},
	}
	for _, tt := range tests {
		got := NewDuplicate(tt.args.msgOrErr)
		assert.Equalf(t, tt.want, got, "NewDuplicate(%v)", tt.args.msgOrErr)
	}
}

func TestNewNotFound(t *testing.T) {
	type args struct {
		msgOrErr any
	}
	tests := []struct {
		name string
		args args
		want *Throw
	}{
		{
			name: "Error arg",
			args: args{msgOrErr: errors.New("test case 1")},
			want: &Throw{kind: NotFound, msg: "test case 1"},
		},
		{
			name: "String arg",
			args: args{msgOrErr: "test case 2"},
			want: &Throw{kind: NotFound, msg: "test case 2"},
		},
		{
			name: "Integer arg",
			args: args{msgOrErr: 3},
			want: &Throw{kind: NotFound, msg: "3"},
		},
	}
	for _, tt := range tests {
		got := NewNotFound(tt.args.msgOrErr)
		assert.Equalf(t, tt.want, got, "NewNotFound(%v)", tt.args.msgOrErr)
	}
}

func TestNewBadRequest(t *testing.T) {
	type args struct {
		msgOrErr any
	}
	tests := []struct {
		name string
		args args
		want *Throw
	}{
		{
			name: "Error arg",
			args: args{msgOrErr: errors.New("test case 1")},
			want: &Throw{kind: BadRequest, msg: "test case 1"},
		},
		{
			name: "String arg",
			args: args{msgOrErr: "test case 2"},
			want: &Throw{kind: BadRequest, msg: "test case 2"},
		},
		{
			name: "Integer arg",
			args: args{msgOrErr: 3},
			want: &Throw{kind: BadRequest, msg: "3"},
		},
	}
	for _, tt := range tests {
		got := NewBadRequest(tt.args.msgOrErr)
		assert.Equalf(t, tt.want, got, "NewBadRequest(%v)", tt.args.msgOrErr)
	}
}

func TestNewAuth(t *testing.T) {
	type args struct {
		msgOrErr any
	}
	tests := []struct {
		name string
		args args
		want *Throw
	}{
		{
			name: "Error arg",
			args: args{msgOrErr: errors.New("test case 1")},
			want: &Throw{kind: Auth, msg: "test case 1"},
		},
		{
			name: "String arg",
			args: args{msgOrErr: "test case 2"},
			want: &Throw{kind: Auth, msg: "test case 2"},
		},
		{
			name: "Integer arg",
			args: args{msgOrErr: 3},
			want: &Throw{kind: Auth, msg: "3"},
		},
	}
	for _, tt := range tests {
		got := NewAuth(tt.args.msgOrErr)
		assert.Equalf(t, tt.want, got, "NewAuth(%v)", tt.args.msgOrErr)
	}
}

func TestIs(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Unknown kind error",
			args: args{
				err: &Throw{kind: Unknown, msg: "test error"},
			},
			want: false,
		},
		{
			name: "Non-Unknown kind error",
			args: args{
				err: &Throw{kind: Deny, msg: "test error"},
			},
			want: true,
		},
		{
			name: "Nil Error",
			args: args{
				err: nil,
			},
			want: false,
		},
		{
			name: "Non Throw type error",
			args: args{
				err: errors.New("standard error"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		got := Is(tt.args.err)
		assert.Equalf(t, tt.want, got, "Is(%v)", tt.args.err)
	}
}

func TestCast(t *testing.T) {
	tests := []struct {
		name string
		args error
		want *Throw
	}{
		{
			name: "Nil error",
			args: nil,
			want: &Throw{},
		},
		{
			name: "Standard error",
			args: errors.New("standart error"),
			want: &Throw{
				msg: errors.New("standart error").Error(),
			},
		},
		{
			name: "Throw error",
			args: &Throw{
				kind: NotFound,
				msg:  "test error",
			},
			want: &Throw{
				kind: NotFound,
				msg:  "test error",
			},
		},
	}
	for _, tt := range tests {
		got := Cast(tt.args)
		assert.Equalf(t, tt.want, got, "Cast(%v)", tt.args)
	}
}

func TestIsDuplicate(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Duplicate kind error",
			args: args{
				err: &Throw{kind: Duplicate, msg: "test error"},
			},
			want: true,
		},
		{
			name: "Non Duplicate kind error",
			args: args{
				err: &Throw{kind: Unknown, msg: "test error"},
			},
			want: false,
		},
		{
			name: "Nil error",
			args: args{
				err: nil,
			},
			want: false,
		},
		{
			name: "Non Throw error",
			args: args{
				err: errors.New("standard error"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		got := IsDuplicate(tt.args.err)
		assert.Equalf(t, tt.want, got, "IsDuplicate(%v)", tt.args.err)
	}
}

func TestIsNotFound(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "NotFound kind error",
			args: args{
				err: &Throw{kind: NotFound, msg: "test error"},
			},
			want: true,
		},
		{
			name: "Non NotFound kind error",
			args: args{
				err: &Throw{kind: Duplicate, msg: "test error"},
			},
			want: false,
		},
		{
			name: "Nil error",
			args: args{
				err: nil,
			},
			want: false,
		},
		{
			name: "Non Throw error",
			args: args{
				err: errors.New("standard error"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		got := IsNotFound(tt.args.err)
		assert.Equalf(t, tt.want, got, "IsNotFound(%v)", tt.args.err)
	}
}

func TestIsDeny(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Deny kind error",
			args: args{
				err: &Throw{kind: Deny, msg: "test error"},
			},
			want: true,
		},
		{
			name: "Non Deny kind error",
			args: args{
				err: &Throw{kind: NotFound, msg: "test error"},
			},
			want: false,
		},
		{
			name: "Nil error",
			args: args{
				err: nil,
			},
			want: false,
		},
		{
			name: "Non Throw error",
			args: args{
				err: errors.New("standard error"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		assert.Equalf(t, tt.want, IsDeny(tt.args.err), "IsDeny(%v)", tt.args.err)
	}
}

func TestIsAuth(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Auth kind error",
			args: args{
				err: &Throw{kind: Auth, msg: "test error"},
			},
			want: true,
		},
		{
			name: "Non Auth kind error",
			args: args{
				err: &Throw{kind: NotFound, msg: "test error"},
			},
			want: false,
		},
		{
			name: "Nil error",
			args: args{
				err: nil,
			},
			want: false,
		},
		{
			name: "Non Throw error",
			args: args{
				err: errors.New("standard error"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		got := IsAuth(tt.args.err)
		assert.Equalf(t, tt.want, got, "IsAuth(%v)", tt.args.err)
	}
}

func TestIsBadRequest(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "BadRequest kind error",
			args: args{
				err: &Throw{kind: BadRequest, msg: "test error"},
			},
			want: true,
		},
		{
			name: "Non BadRequest kind error",
			args: args{
				err: &Throw{kind: NotFound, msg: "test error"},
			},
			want: false,
		},
		{
			name: "Nil error",
			args: args{
				err: nil,
			},
			want: false,
		},
		{
			name: "Non Throw error",
			args: args{
				err: errors.New("standard error"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		got := IsBadRequest(tt.args.err)
		assert.Equalf(t, tt.want, got, "IsBadRequest(%v)", tt.args.err)
	}
}

func Test_reasonMetadataKeyReg(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			name: "good",
			args: "email",
			want: true,
		},
		{
			name: "good",
			args: "EMAIL",
			want: true,
		},
	}
	for _, tt := range tests {
		assert.Equalf(t, tt.want, reasonMetadataKeyReg.MatchString(tt.args), "IsBadRequest(%v)", tt.args)
	}
}
