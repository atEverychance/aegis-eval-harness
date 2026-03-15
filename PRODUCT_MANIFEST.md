# Aegis Eval Harness - Product Manifest

**Status:** Approved for implementation  
**Source:** Athena (bigBrain)  
**Date:** 2026-03-15

---

## Core Concept

A general-purpose eval harness for agentic systems that scores runs against ground truth fixtures, tracks token/latency costs, and profiles performance over time.

**Plain-English pitch:** "Pytest for agents, but with ground truth, scoring, cost accounting, and replayable episodes."

---

## The Gap

| Current (Artemis) | Missing Layer |
|---|---|
| Did it run without error? | Was it correct? |
| Pass/fail | How correct? |
| | What did it cost? |
| | How long did it take? |
| | Is it improving over time? |

---

## Feel

- **Clinical, not cute** — instrumentation, not vibes
- **Fast feedback** — run evals, get signal immediately
- **Trustworthy** — obvious scoring rules, inspectable fixtures
- **Local-first** — fixtures on disk, SQLite for runs, easy to diff
- **Boring in the best way** — stable, scriptable, composable

---

## Personas

1. **Agent Builder** — validate changes without blowing up cost/latency
2. **Evaluator/QA** — turn subjective quality into repeatable suites
3. **Product/Research Lead** — compare variants over time
4. **RL Tinkerer** — treat runs as episodes for training loops
5. **Cross-Pantheon Operator** — neutral harness usable everywhere

---

## Core Objects

- **Eval Suite** — named collection of test cases
- **Fixture** — ground-truth-backed test (input, expected, scorer, tags)
- **Run** — one execution against a system config
- **Episode** — per-case artifact (output, tokens, latency, score breakdown)
- **Scorer** — pluggable comparison (exact, JSON, semantic, numeric, rubric)
- **Profile/Baseline** — saved reference for comparison

---

## Storage (v1)

```
evals/
  suites/
    <suite-name>/
      suite.yaml
      fixtures/
        001.yaml
        002.yaml
  runs/
    artifacts/
  aegis.db
```

- Fixtures: YAML/JSON on disk (git-versionable)
- Runs/Episodes: SQLite
- Artifacts: filesystem

---

## CLI Commands (v1)

```
eval init <name>
eval run <suite>
eval score <run>
eval compare <runA> <runB>
eval history <suite>
eval inspect <episode>
```

Output: table (humans), JSON (automation), markdown (sharing)

---

## Scoring Model

### A. Correctness
- exact match, regex, JSON field-level, schema, numeric tolerance, semantic/rubric

### B. Cost
- prompt/completion/total tokens, estimated $

### C. Performance
- wall-clock latency, first-token latency, tool-call count, step count

### D. Quality Composite (optional)
- correctness + cost + latency with configurable weights

---

## Success Criteria

Teams can answer with evidence:
1. Did the new version improve correctness?
2. At acceptable cost?
3. Did latency regress?
4. Which fixtures got better/worse?
5. Is performance trending up/down?

---

## v1 Scope (MVP)

### In Scope
- ✅ Fixture system (YAML/JSON)
- ✅ Suite manifests with tags/categories
- ✅ Input adapter contract (any agent can submit)
- ✅ 4 scorers: exact, JSON, rubric, numeric
- ✅ SQLite run/episode history
- ✅ CLI reporting with breakdown
- ✅ Baseline comparison
- ✅ Run history + trendlines
- ✅ Pantheon-portable (no godmode dependency)

### Out of Scope
- ❌ Browser dashboard
- ❌ Multimodal scoring
- ❌ Automatic reward training
- ❌ Cloud infra
- ❌ CI/CD (future)

---

## v2 Evolution

- Regression alerting
- CI/CD integration (when we have it)
- Richer scorer plugins
- Pairwise comparison
- Failure clustering
- Trace-aware scoring
- Benchmark packs

---

## v3 Evolution

- Episode replay
- Reward signal generation
- Policy optimization hooks
- Auto-fixture generation from failures
- RL/training data curation
- Multimodal eval support

---

## GitHub Integration

- Repo: `atEverychance/aegis-eval-harness`
- Issues for tracking
- Use godmode project boards
- No labels — assignee-based tracking only
