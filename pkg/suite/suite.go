package suite

import (
    "fmt"
    "os"
    "path/filepath"

    "gopkg.in/yaml.v3"
)

// Suite is the manifest that describes an eval suite.
type Suite struct {
    Name        string   `yaml:"name"`
    Description string   `yaml:"description,omitempty"`
    Tags        []string `yaml:"tags,omitempty"`
    Categories  []string `yaml:"categories,omitempty"`
    Fixtures    []string `yaml:"fixtures,omitempty"`
}

// LoadSuite reads suites/<name>/suite.yaml and unmarshals the manifest.
func LoadSuite(name string) (*Suite, error) {
    manifestPath := filepath.Join("suites", name, "suite.yaml")
    data, err := os.ReadFile(manifestPath)
    if err != nil {
        return nil, fmt.Errorf("read suite manifest %s: %w", manifestPath, err)
    }

    var suite Suite
    if err := yaml.Unmarshal(data, &suite); err != nil {
        return nil, fmt.Errorf("parse suite manifest %s: %w", manifestPath, err)
    }

    if suite.Name == "" {
        suite.Name = name
    }

    return &suite, nil
}
