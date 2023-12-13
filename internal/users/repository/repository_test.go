package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/martinyonatann/go-unit-test/config"
	"github.com/martinyonatann/go-unit-test/internal/users"
	"github.com/martinyonatann/go-unit-test/internal/users/entities"
	"github.com/martinyonatann/go-unit-test/internal/users/repository"
	"github.com/martinyonatann/go-unit-test/pkg/databases"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	cfg config.Config
	db  *sqlx.DB
)

const (
	MIGRATION_PATH = "file://../../database/migrations"
)

func init() {
	var err error

	cfg, err = config.LoadConfigPath("../../../config/config.integration.test")
	if err != nil {
		panic(errors.Wrap(err, "config.LoadConfigPath"))
	}

	db, err = databases.NewMySQLDB(context.Background(), cfg.Database)
	if err != nil {
		panic(errors.Wrap(err, cfg.Database.DBName))
	}

	migrationUP(MIGRATION_PATH, db.DB)
}

func migrationUP(path string, db *sql.DB) error {
	// Set the database instance for the "mysql" driver
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	// Set up the migrate instance with the MySQL database driver and file source
	m, err := migrate.NewWithDatabaseInstance(path, "mysql", driver)
	if err != nil {
		return err
	}

	defer m.Close()

	// Apply migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func migrationDOWN(path string, db *sql.DB) error {
	// Set the database instance for the "mysql" driver
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	// Set up the migrate instance with the MySQL database driver and file source
	m, err := migrate.NewWithDatabaseInstance(path, "mysql", driver)
	if err != nil {
		return err
	}

	defer m.Close()

	// Apply migrations
	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

type repositoryTestSuite struct {
	repo users.Repositories
	suite.Suite
}

func TestSuiteRepository(t *testing.T) {
	defer func() {
		time.Sleep(5 * time.Second)

		err := migrationDOWN(MIGRATION_PATH, db.DB)
		require.NoError(t, err)
	}()

	suite.Run(t, &repositoryTestSuite{repo: repository.NewRepository(db)})
}

func (r *repositoryTestSuite) Test_Repo_Repositories() {
	var users = entities.Users{
		ID:        uuid.NewString(),
		Name:      "go-unit-test",
		Password:  faker.Password(),
		CreatedAt: time.Now(),
	}

	err := r.repo.Create(context.Background(), users)
	r.Assert().NoError(err)

	detail, err := r.repo.Detail(context.Background(), users.ID)
	r.Assert().NoError(err)
	r.Equal(users, detail)
}
