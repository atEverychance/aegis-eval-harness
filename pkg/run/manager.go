package run

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Manager struct {
	db           *sql.DB
	artifactsDir string
}

func NewManager(db *sql.DB, artifactsDir string) *Manager {
	if artifactsDir == "" {
		artifactsDir = filepath.Join("evals", "runs", "artifacts")
	}
	return &Manager{db: db, artifactsDir: artifactsDir}
}

func (m *Manager) CreateRun(ctx context.Context, suiteID string, systemConfig any, model, promptVersion, notes string) (string, error) {
	runID := uuid.NewString()
	configJSON, err := marshalSystemConfig(systemConfig)
	if err != nil {
		return "", err
	}
	if _, err := m.db.ExecContext(ctx,
		`INSERT INTO runs (id, suite_id, system_config, model, prompt_version, timestamp, notes) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		runID,
		suiteID,
		configJSON,
		model,
		promptVersion,
		time.Now().UTC().Format(time.RFC3339),
		notes,
	); err != nil {
		return "", fmt.Errorf("insert run: %w", err)
	}
	return runID, nil
}

func marshalSystemConfig(cfg any) (string, error) {
	if cfg == nil {
		return "", nil
	}
	switch v := cfg.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	default:
		bytes, err := json.Marshal(v)
		if err != nil {
			return "", fmt.Errorf("marshal system config: %w", err)
		}
		return string(bytes), nil
	}
}

func (m *Manager) StoreEpisode(ctx context.Context, runID, fixtureID, actualOutput string, tokensUsed, latencyMs, costCents int, score float64, passed bool, errorMsg string) (string, error) {
	episodeID := uuid.NewString()
	artifactPath, err := m.writeArtifact(runID, episodeID, actualOutput)
	if err != nil {
		return "", err
	}
	passedInt := 0
	if passed {
		passedInt = 1
	}
	if _, err := m.db.ExecContext(ctx,
		`INSERT INTO episodes (id, run_id, fixture_id, actual_output, tokens_used, latency_ms, cost_cents, score, passed, error_msg, artifact_path) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		episodeID,
		runID,
		fixtureID,
		actualOutput,
		tokensUsed,
		latencyMs,
		costCents,
		score,
		passedInt,
		errorMsg,
		artifactPath,
	); err != nil {
		return "", fmt.Errorf("insert episode: %w", err)
	}
	return episodeID, nil
}

func (m *Manager) writeArtifact(runID, episodeID, output string) (string, error) {
	if m.artifactsDir == "" {
		return "", fmt.Errorf("artifacts directory is not configured")
	}
	rootDir := filepath.Join(m.artifactsDir, runID)
	if err := os.MkdirAll(rootDir, 0o755); err != nil {
		return "", fmt.Errorf("ensure artifact dir: %w", err)
	}
	path := filepath.Join(rootDir, fmt.Sprintf("%s.out", episodeID))
	if err := os.WriteFile(path, []byte(output), 0o644); err != nil {
		return "", fmt.Errorf("write artifact: %w", err)
	}
	return path, nil
}
