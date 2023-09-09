package provider

import (
	"database/sql"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

const (
	MsgErrNoRecordRedis = "no record in redis"
)

// IsDuplicateKeyErr является ли ошибки БД следствием ограничения уникальности значения в поле
func IsDuplicateKeyErr(err error) bool {
	if e, ok := err.(*pgconn.PgError); ok {
		return e.Code == "23505"
	}

	return false
}

// IsNoRow является ли ошибка БД следствием отсутствия запрошенной строки
func IsNoRow(err error) bool {
	return errors.Is(err, pgx.ErrNoRows) || errors.Is(err, sql.ErrNoRows)
}

// IsNoRec является ли ошибка Redis следствием отсутствия записи
func IsNoRec(err error) bool {
	return errors.Is(err, redis.Nil)
}

// IsCheckErr ограничивающее условие для значения поля в таблице
func IsCheckErr(err error) bool {
	if e, ok := err.(*pgconn.PgError); ok {
		return e.Code == "23514"
	}

	return false
}

// CheckConstraintErr ограничивающее условие для значения поля в таблице
func CheckConstraintErr(err error) (nameConstraint string, ok bool) {
	if e, ok := err.(*pgconn.PgError); ok && e.Code == "23514" {
		return e.ConstraintName, true
	}

	return "", false
}

// IsConstrainForeignKeyErr ограничение удаления записи
func IsConstrainForeignKeyErr(err error) bool {
	if e, ok := err.(*pgconn.PgError); ok {
		return e.Code == "23503"
	}

	return false
}

// OID 16495

// IsUnknownTypeErr неизвестный тип
func IsUnknownTypeErr(err error) bool {
	if e, ok := err.(*pgconn.PgError); ok {
		return e.Code == "16495"
	}

	return false
}

// ConstraintErr является ли ошибки БД следствием ограничения уникальности значения в поле
func ConstraintErr(err error) (nameConstraint string, ok bool) {
	if e, ok := err.(*pgconn.PgError); ok && (e.Code == "23505" || e.Code == "23514") {
		return e.ConstraintName, true
	}

	return "", false
}
