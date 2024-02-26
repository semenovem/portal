package user_tool

import (
	"strings"
)

type Group string
type Status string

const (
	GroupUnknown    Group = ""
	GroupSuperAdmin Group = "super_admin"
	GroupAdmin      Group = "admin"
	GroupUser       Group = "user"
)

const (
	StatusUnknown       Status = ""
	StatusInactive      Status = "inactive"
	StatusActive        Status = "active"
	StatusBlocked       Status = "blocked"
	StatusRegistration  Status = "registration"
	StatusAwaitActivate Status = "await_activate"
)

func ParseStatus(s string) (Status, bool) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case string(StatusInactive):
		return StatusInactive, true
	case string(StatusActive):
		return StatusActive, true
	case string(StatusBlocked):
		return StatusBlocked, true
	case string(StatusRegistration):
		return StatusRegistration, true
	case string(StatusAwaitActivate):
		return StatusAwaitActivate, true
	}

	return StatusUnknown, false
}

func MustParseUserStatus(s string) Status {
	st, _ := ParseStatus(s)
	return st
}

func ParseGroup(s string) (Group, bool) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case string(GroupSuperAdmin):
		return GroupSuperAdmin, true
	case string(GroupAdmin):
		return GroupAdmin, true
	case string(GroupUser):
		return GroupUser, true
	}

	return GroupUnknown, false
}

func ParseUserGroups(groups []string) ([]Group, bool) {
	var result []Group

	for _, group := range groups {
		r, ok := ParseGroup(group)
		if !ok {
			return nil, false
		}
		result = append(result, r)
	}

	return result, true
}

func StringifyUserGroups(a []Group) []string {
	b := make([]string, len(a))
	for i := range a {
		b[i] = string(a[i])
	}
	return b
}

func InflateUserGroups(a []string) []Group {
	b := make([]Group, len(a))
	for i := range a {
		b[i] = Group(a[i])
	}
	return b
}
