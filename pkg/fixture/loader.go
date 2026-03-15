package fixture

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Fixture struct {
	Name       string   `json:"name" yaml:"name"`
	Input      string   `json:"input" yaml:"input"`
	Expected   string   `json:"expected" yaml:"expected"`
	ScorerType string   `json:"scorer_type" yaml:"scorer_type"`
	Rubric     string   `json:"rubric,omitempty" yaml:"rubric,omitempty"`
	Tags       []string `json:"tags,omitempty" yaml:"tags,omitempty"`
	Difficulty string   `json:"difficulty,omitempty" yaml:"difficulty,omitempty"`
}

func (f *Fixture) Validate() error {
	var missing []string
	if strings.TrimSpace(f.Name) == "" {
		missing = append(missing, "name")
	}
	if strings.TrimSpace(f.Input) == "" {
		missing = append(missing, "input")
	}
	if strings.TrimSpace(f.Expected) == "" {
		missing = append(missing, "expected")
	}
	if strings.TrimSpace(f.ScorerType) == "" {
		missing = append(missing, "scorer_type")
	}
	if len(missing) > 0 {
		return fmt.Errorf("missing required fixture fields: %s", strings.Join(missing, ", "))
	}
	return nil
}

func LoadFixture(path string) (*Fixture, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var fixture Fixture
	switch strings.ToLower(filepath.Ext(path)) {
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, &fixture); err != nil {
			return nil, fmt.Errorf("failed to parse YAML fixture %s: %w", path, err)
		}
	case ".json":
		if err := json.Unmarshal(data, &fixture); err != nil {
			return nil, fmt.Errorf("failed to parse JSON fixture %s: %w", path, err)
		}
	default:
		return nil, fmt.Errorf("unsupported fixture file extension %s", path)
	}

	if err := fixture.Validate(); err != nil {
		return nil, err
	}
	return &fixture, nil
}
