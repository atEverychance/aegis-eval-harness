package db

import (
    "database/sql"
    "embed"
    "fmt"
    "os"
    "path/filepath"

    _ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var schemaFS embed.FS

// Init opens a SQLite database at dbPath, runs any pending migrations, and
// returns the ready-to-use connection.
func Init(dbPath string) (*sql.DB, error) {
    if err := os.MkdirAll(filepath.Dir(dbPath), 0o755); err != nil {
        return nil, fmt.Errorf("ensure path: %w", err)
    }

    db, err := sql.Open("sqlite3", dbPath)
    if err != nil {
        return nil, fmt.Errorf("open sqlite: %w", err)
    }

    db.SetMaxOpenConns(1)

    if err := db.Ping(); err != nil {
        db.Close()
        return nil, fmt.Errorf("ping sqlite: %w", err)
    }

    if err := initializeSchema(db); err != nil {
        db.Close()
        return nil, err
    }

    return db, nil
}

func initializeSchema(db *sql.DB) error {
    schema, err := schemaFS.ReadFile("schema.sql")
    if err != nil {
        return fmt.Errorf("read schema: %w", err)
    }

    if _, err := db.Exec(string(schema)); err != nil {
        return fmt.Errorf("exec schema: %w", err)
    }

    return nil
}
