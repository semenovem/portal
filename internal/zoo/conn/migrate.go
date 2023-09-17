package conn

import (
	//"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/semenovem/portal/pkg"
)

func Migrate(logger pkg.Logger, dbStr, path string) error {
	if path == "" {
		return nil
	}

	var (
		ll = logger.Named("Migrate")
		p  = fmt.Sprint("file://", path)
	)

	m, err := migrate.New(p, dbStr)
	if err != nil {
		ll.Named("New").DB(err)
		return err
	}

	defer m.Close()

	if err = m.Up(); err != nil {
		if err.Error() == "no change" {
			ll.Debugf(err.Error())
			return nil
		}

		ll.Named("Up").DB(err)
		return err
	}

	ll.Info("applying new migrations")

	return nil
}
