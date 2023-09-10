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

func MigrateDev(logger pkg.Logger, dbStr, path string) error {
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

	// Под таким именем монтируются файлы в scripts/local-stand/stand.sh
	if err = m.Force(20250525141415); err != nil {
		ll.Named("force").DB(err)
		return err
	}

	if err = m.Down(); err != nil {
		if err.Error() == "no change" {
			ll.Debugf(err.Error())
			return nil
		}

		ll.Named("Down").DB(err)
		return err
	}

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
