CREATE TABLE IF NOT EXISTS suites (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT,
    tags TEXT,
    categories TEXT,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS fixtures (
    id TEXT PRIMARY KEY,
    suite_id TEXT NOT NULL REFERENCES suites(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    input TEXT NOT NULL,
    expected TEXT,
    scorer_type TEXT,
    rubric TEXT,
    tags TEXT,
    difficulty INTEGER,
    created_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS runs (
    id TEXT PRIMARY KEY,
    suite_id TEXT NOT NULL REFERENCES suites(id) ON DELETE CASCADE,
    system_config TEXT,
    model TEXT,
    prompt_version TEXT,
    timestamp TEXT NOT NULL,
    notes TEXT
);

CREATE TABLE IF NOT EXISTS episodes (
    id TEXT PRIMARY KEY,
    run_id TEXT NOT NULL REFERENCES runs(id) ON DELETE CASCADE,
    fixture_id TEXT NOT NULL REFERENCES fixtures(id) ON DELETE CASCADE,
    actual_output TEXT,
    tokens_used INTEGER,
    latency_ms INTEGER,
    cost_cents INTEGER,
    score REAL,
    passed INTEGER NOT NULL CHECK (passed IN (0,1)),
    error_msg TEXT,
    artifact_path TEXT
);

CREATE TABLE IF NOT EXISTS baselines (
    id TEXT PRIMARY KEY,
    suite_id TEXT NOT NULL REFERENCES suites(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    run_id TEXT REFERENCES runs(id),
    created_at TEXT NOT NULL
);
