package db

import (
	gen "bookAPI/gen/book"
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Required to use file-based migrations
	"github.com/jmoiron/sqlx"
)

// Storage defines the methods required by the repository layer.
type Storage interface {
	CreateBook(ctx context.Context, book *gen.Book) error
	GetBookByID(ctx context.Context, id string) (*gen.Book, error)
	UpdateBook(ctx context.Context, book *gen.Book) error
	DeleteBook(ctx context.Context, id string) error
}

// BookStorage implements the Storage interface using sqlx.
type BookStorage struct {
	db *sqlx.DB
}

// NewStorage returns a new instance of BookStorage.
func NewStorage() (*BookStorage, error) {
	// Load DSN from environment or use default
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "root:password@tcp(localhost:3306)/bookapi?parseTime=true"
	}

	// Connect to the database
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}
	fmt.Println("Connected to MySQL successfully!")

	databaseURL := fmt.Sprintf("mysql://%s", dsn)
	// Run database migrations
	if err := runMigrations("file://./db/migrations", databaseURL); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
		return nil, err
	}

	return &BookStorage{db: db}, nil
}

// runMigrations runs database migrations using the golang-migrate package.
func runMigrations(migrationURL, databaseURL string) error {
	// Create a new migrate instance
	m, err := migrate.New(migrationURL, databaseURL)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %v", err)
	}

	// Run the migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		// ErrNoChange is returned when the database is already up-to-date.
		return fmt.Errorf("failed to run migrations: %v", err)
	}

	log.Println("Migrations ran successfully")
	return nil
}
