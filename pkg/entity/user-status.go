package entity

type UserStatus string

const (
	UserStatusInactive UserStatus = "inactive"
	UserStatusActive   UserStatus = "active"
	UserStatusBlocked  UserStatus = "blocked"
)
