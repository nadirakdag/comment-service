package postgres

import (
	"comment-service/config"
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *sqlx.DB
}

func NewDatabase(config config.Database) (*Database, error) {
	connectionString := fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=%s", config.Host, config.Port, config.Username, config.Database, config.Password, config.SslMode)

	dbConnection, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	return &Database{Client: dbConnection}, nil
}

func (d *Database) Ping(ctx context.Context) error {
	return d.Client.PingContext(ctx)
}

func (d *Database) MigrateDB() error {
	fmt.Println("migrating our database")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("could not create the postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://internal/db/postgres/migrations", "postgres", driver)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err := m.Up(); err != nil {
		return fmt.Errorf("could not run up migrations: %v", err)
	}

	fmt.Println("successfully migrated the database")

	return nil
}
