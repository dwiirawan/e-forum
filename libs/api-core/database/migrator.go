package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"
)

type DBConfig struct {
	Driver string
	Source string
}

func MigratorNew() DBConfig {
	driver := os.Getenv("DB_DRIVER")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	requiredEnvVars := map[string]string{
		"DB_NAME":     dbName,
		"DB_USER":     dbUser,
		"DB_PASSWORD": dbPassword,
		"DB_HOST":     dbHost,
		"DB_PORT":     dbPort,
		"DB_DRIVER":   driver,
	}

	for envVar, value := range requiredEnvVars {
		if len(value) == 0 {
			panic("ENV " + envVar + " IS NOT SET")
		}
	}

	return DBConfig{
		Driver: driver,
		Source: driver + "://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable"}
}

func (d DBConfig) DBGenerate(name string) {

	migrations := &migrate.FileMigrationSource{
		Dir: "../migrations",
	}

	migrations.FindMigrations()

}

func (d DBConfig) DBUp() error {
	fmt.Println("Migration Starting")
	db, err := sql.Open(d.Driver, d.Source)
	if err != nil {
		panic(err)

	}
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	migrate.SetTable("history_migrations")

	_, err = migrate.Exec(db, d.Driver, migrations, migrate.Up)
	if err != nil {
		fmt.Println("Error in MIGRATE Function", err)
		panic(err)
	}
	defer db.Close()
	fmt.Println("Migration success")
	return nil
}

func (d DBConfig) DBDown() error {
	db, err := sql.Open(d.Driver, d.Source)
	if err != nil {
		return err
	}
	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}
	migrate.SetTable("history_migrations")
	_, err = migrate.Exec(db, d.Driver, migrations, migrate.Down)
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Println("Migration down success")
	return nil
}

func (d DBConfig) DBStatus() string {
	db, err := sql.Open(d.Driver, d.Source)
	if err != nil {
		return "Error connecting to database"
	}
	defer db.Close()

	migrate.SetTable("history_migrations")

	records, err := migrate.GetMigrationRecords(db, d.Driver)
	if err != nil {
		return "Error getting migration status"
	}

	fmt.Println("Migration Status:")
	for _, record := range records {
		fmt.Printf("Applied: %s\n", record.Id)
	}
	return "Migration status check completed"
}
