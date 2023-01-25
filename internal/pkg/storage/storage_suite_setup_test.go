package storage

import (
	"database/sql"
	"embed"
	"io/fs"
	"os"
	"testing"

	embeddedpg "github.com/fergusstrange/embedded-postgres"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

// RepositorySuite - this struct holds the shared test state and helper methods for tests.
type RepositorySuite struct {
	suite.Suite
	pg *embeddedpg.EmbeddedPostgres

	repo Repository
}

// TestSuite - uses the suite.Run method to run the tests in the suite.
func TestSuite(t *testing.T) {
	t.Parallel()

	// Create the test suite and run the tests.
	suite.Run(t, new(RepositorySuite))
}

// SetupTest - defines method that will be called before each test in the suite.
func (s *RepositorySuite) SetupTest() {
	// Use ZeroLog as a default logger.
	logger := zerolog.New(os.Stdout)

	epConf := embeddedpg.DefaultConfig().Logger(logger)
	epConf = epConf.Port(9337)

	// Create the embedded Postgres instance.
	s.pg = embeddedpg.NewDatabase(epConf)

	// Use the embedded Postgres instance to connect to the database.
	if err := s.pg.Start(); err != nil {
		s.T().Fatal(err)
	}

	db, err := connect()
	if err != nil {
		s.T().Fatal(err)
	}

	s.repo = Repository{
		db:       db,
		tenantDB: db,
	}

	err = migrateDB(db)
	if err != nil {
		s.T().Fatal(err)
	}
}

func (s *RepositorySuite) TearDownTest() {
	// Close the embedded Postgres instance.
	if err := s.pg.Stop(); err != nil {
		s.T().Fatal(err)
	}
}

// connect - opens a new test DB connection.
func connect() (*bun.DB, error) {
	sqldb, err := sql.Open("postgres", "host=localhost port=9337 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	db := bun.NewDB(sqldb, pgdialect.New())

	return db, err
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

// migrateDB - reads and applies migrations from embedded *.sql files.
func migrateDB(db *bun.DB) error {
	// // Read the migration files
	migrationFiles, _ := fs.ReadDir(embedMigrations, "migrations")
	for _, cppFile := range migrationFiles {
		data, err := embedMigrations.ReadFile("migrations/" + cppFile.Name())
		if err != nil {
			return err
		}

		err = createTables(db, string(data))
		if err != nil {
			return err
		}
	}

	return nil
}

// createTables - executes DB migrations.
func createTables(db *bun.DB, script string) error {
	_, err := db.Exec(script)
	if err != nil {
		return err
	}

	return nil
}
