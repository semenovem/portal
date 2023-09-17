package it

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type LoginPasswdAuthenticator interface {
	Hashing(passwd string) string
	Matching(hash, passwd string) bool
}

type userPasswdAuth struct {
	salt []byte
}

func NewLoginPasswdAuthenticator(salt string) LoginPasswdAuthenticator {
	return &userPasswdAuth{
		salt: []byte(salt),
	}
}

func (u *userPasswdAuth) Hashing(passwd string) string {
	if passwd == "" {
		return ""
	}

	var (
		b    = []byte(fmt.Sprint(passwd))
		hash = sha1.Sum(append(b, u.salt...))
	)

	return hex.EncodeToString(hash[:])
}

func (u *userPasswdAuth) Matching(hash, passwd string) bool {
	if hash == "" || passwd == "" {
		return false
	}

	return u.Hashing(passwd) == hash
}
