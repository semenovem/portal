package conn

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	pg "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/semenovem/portal/pkg"
)

func Migrate(logger pkg.Logger, dbx *sql.DB, path string) error {
	var (
		ll = logger.Named("Migrate")
		p  = fmt.Sprint("file://", path)
	)

	driver, err := pg.WithInstance(dbx, &pg.Config{})
	if err != nil {
		ll.Named("WithInstance").Error(err.Error())
		return err
	}

	mid, err := migrate.NewWithDatabaseInstance(p, "postgres", driver)
	if err != nil {
		ll.Named("NewWithDatabaseInstance").Error(err.Error())
		return err
	}

	if err = mid.Up(); err != nil {
		if err.Error() == "no change" {
			ll.With("migrations", err).Info("no change")
			return nil
		}

		ll.Named("Up").Error(err.Error())
		return err
	}

	ll.Info("applying new migrations")

	return nil
}
