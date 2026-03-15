package run

import (
    "context"
    "os"
    "path/filepath"
    "testing"
    "time"

    "github.com/atEverychance/aegis-eval-harness/db"
)

func TestManager_CreateRunAndStoreEpisode(t *testing.T) {
    tmp := t.TempDir()
    dbPath := filepath.Join(tmp, "test.db")

    database, err := db.Init(dbPath)
    if err != nil {
        t.Fatalf("failed to init db: %v", err)
    }
    defer database.Close()

    suiteID := "suite-test"
    fixtureID := "fixture-test"
    now := time.Now().UTC().Format(time.RFC3339)

    if _, err := database.ExecContext(context.Background(),
        `INSERT INTO suites (id, name, created_at, updated_at) VALUES (?, ?, ?, ?)`,
        suiteID,
        "test suite",
        now,
        now,
    ); err != nil {
        t.Fatalf("insert suite: %v", err)
    }

    if _, err := database.ExecContext(context.Background(),
        `INSERT INTO fixtures (id, suite_id, name, input, expected, created_at) VALUES (?, ?, ?, ?, ?, ?)`,
        fixtureID,
        suiteID,
        "fixture",
        "input",
        "expected",
        now,
    ); err != nil {
        t.Fatalf("insert fixture: %v", err)
    }

    artifactsDir := filepath.Join(tmp, "artifacts")
    manager := NewManager(database, artifactsDir)

    runID, err := manager.CreateRun(context.Background(), suiteID, map[string]any{"env": "test"}, "model-v1", "prompt-v1", "notes")
    if err != nil {
        t.Fatalf("CreateRun failed: %v", err)
    }
    if runID == "" {
        t.Fatalf("expected runID")
    }

    episodeID, err := manager.StoreEpisode(context.Background(), runID, fixtureID, "actual-output", 123, 456, 789, 0.95, true, "")
    if err != nil {
        t.Fatalf("StoreEpisode failed: %v", err)
    }
    if episodeID == "" {
        t.Fatalf("expected episodeID")
    }

    var artifactPath string
    if err := database.QueryRowContext(context.Background(),
        `SELECT artifact_path FROM episodes WHERE id = ?`,
        episodeID,
    ).Scan(&artifactPath); err != nil {
        t.Fatalf("fetch episode: %v", err)
    }

    if artifactPath == "" {
        t.Fatalf("expected artifact path")
    }
    if _, err := os.Stat(artifactPath); err != nil {
        t.Fatalf("artifact missing: %v", err)
    }
}
