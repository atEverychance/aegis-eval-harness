# Aegis Eval Harness - Atomic Tasks

## Phase 1: Foundation (Core Infrastructure)

| # | Task | Description | Exit Criteria |
|---|------|-------------|---------------|
| 1.1 | **Project setup + Go module** | Initialize Go module, create directory structure per spec (`evals/suites/`, `evals/runs/artifacts/`, `evals/aegis.db`) | `go.mod` exists, dirs created |
| 1.2 | **SQLite schema** | Create tables: `suites`, `runs`, `episodes`, `baselines` | Schema compiles, `aegis.db` initializes |
| 1.3 | **CLI scaffolding** | Implement `eval` CLI with subcommands (init, run, score, compare, history, inspect) | `eval --help` outputs all commands |
| 1.4 | **Fixture loader** | YAML/JSON fixture parsing (input, expected, scorer, tags fields) | Loads sample fixture without error |

## Phase 2: Suite System

| # | Task | Description | Exit Criteria |
|---|------|-------------|---------------|
| 2.1 | **Suite manifest struct** | Define `Suite` struct with name, description, tags, categories | Struct passes `go vet` |
| 2.2 | **Suite loader** | Load suite.yaml from `suites/<name>/suite.yaml` | Loads sample suite correctly |
| 2.3 | **Suite CLI (init)** | Implement `eval init <name>` to create scaffold | Creates valid directory structure |

## Phase 3: Core Scorers (4 Independent Implementations)

| # | Task | Description | Exit Criteria |
|---|------|-------------|---------------|
| 3.1 | **Exact match scorer** | Compare output string to expected string exactly | Pass/fail for "hello" vs "hello" / "hello" vs "world" |
| 3.2 | **JSON scorer** | Compare output JSON to expected JSON (field-level) | Pass for `{"a":1}` vs `{"a":1}`, fail vs `{"a":2}` |
| 3.3 | **Numeric scorer** | Compare numeric output with tolerance (± threshold) | Pass 100±5 for 103, fail 100±5 for 110 |
| 3.4 | **Rubric scorer** | Parse rubric criteria from fixture, return score 0-1 | Returns score 0.0-1.0 for structured input |

## Phase 4: Input Adapter Contract

| # | Task | Description | Exit Criteria |
|---|------|-------------|---------------|
| 4.1 | **Adapter interface** | Define `InputAdapter` interface (any agent submits: input + expected + metadata) | Interface compiles, mock passes |
| 4.2 | **Adapter implementation** | Default adapter reads from fixture files | Adapter populates fixture from disk |

## Phase 5: Run & Episode Storage

| # | Task | Description | Exit Criteria |
|---|------|-------------|---------------|
| 5.1 | **Run creation** | Insert run record with system config, timestamp | Run ID returned, queryable |
| 5.2 | **Episode storage** | Store episode: output, tokens, latency, score breakdown | Episode persists to SQLite |
| 5.3 | **Artifact storage** | Save large outputs to `runs/artifacts/` | File written, path stored in DB |

## Phase 6: CLI Reporting & History

| # | Task | Description | Exit Criteria |
|---|------|-------------|---------------|
| 6.1 | **Score report** | `eval score <run>` outputs breakdown table (case, scorer, score, details) | Table renders correctly |
| 6.2 | **History view** | `eval history <suite>` shows runs over time with trendline data | Lists past runs |
| 6.3 | **Baseline comparison** | `eval compare <runA> <runB>` shows delta (score, cost, latency) | Delta calculated and displayed |
| 6.4 | **Episode inspection** | `eval inspect <episode>` dumps full episode details | JSON output of episode |

---

## Summary

- **Total: 18 atomic tasks**
- **Phases: 6**
- **Estimated: 90-270 min build time**
