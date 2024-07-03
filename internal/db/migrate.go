package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func (d *Database) Migrate() error {
	fmt.Println("Running migrations")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create postgres driver: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrate",
		"postgres",
		driver,
	)
	if err != nil {
		println(err)
		return err
	}

	if err = m.Up(); err != nil {
		return fmt.Errorf("could not run database migrations: %w", err)
	}

	return nil
}
