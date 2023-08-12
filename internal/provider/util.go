package provider

import (
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// IsDuplicateKeyError является ли ошибки БД следствием ограничения дублирования
func IsDuplicateKeyError(err error) bool {
	if e, ok := err.(*pgconn.PgError); ok {
		return e.Code == "23505"
	}

	return false
}

// IsNoRows является ли ошибка БД следствием отсутствия запрошенной строки
func IsNoRows(err error) bool {
	return errors.Is(err, sql.ErrNoRows) || errors.Is(err, pgx.ErrNoRows)
}

// IsConstrainForeignKeyError ограничение удаления записи
func IsConstrainForeignKeyError(err error) bool {
	if e, ok := err.(*pgconn.PgError); ok {
		return e.Code == "23503"
	}

	return false
}
