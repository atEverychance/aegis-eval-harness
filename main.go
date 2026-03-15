package main

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "eval",
    Short: "Aegis Eval Harness CLI",
    Long:  "A CLI for running, inspecting, and comparing eval suites.",
    Run: func(cmd *cobra.Command, args []string) {
        _ = cmd.Help()
    },
}

var suitesRoot = "suites"

func init() {
    rootCmd.AddCommand(initCmd)
    rootCmd.AddCommand(runCmd)
    rootCmd.AddCommand(scoreCmd)
    rootCmd.AddCommand(compareCmd)
    rootCmd.AddCommand(historyCmd)
    rootCmd.AddCommand(inspectCmd)
}

var initCmd = &cobra.Command{
    Use:   "init <name>",
    Short: "Create a new eval suite scaffold",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        suiteDir := filepath.Join(suitesRoot, name)
        fixturesDir := filepath.Join(suiteDir, "fixtures")
        if err := os.MkdirAll(fixturesDir, 0o755); err != nil {
            fmt.Fprintf(os.Stderr, "failed to create suite directories: %v\n", err)
            os.Exit(1)
        }

        suiteFile := filepath.Join(suiteDir, "suite.yaml")
        if _, err := os.Stat(suiteFile); os.IsNotExist(err) {
            content := fmt.Sprintf("name: %s\n# description: Optional description\ntags: []\ncategories: []\nfixtures:\n# - fixtures/001.yaml\n", name)
            if err := os.WriteFile(suiteFile, []byte(content), 0o644); err != nil {
                fmt.Fprintf(os.Stderr, "failed to write suite manifest: %v\n", err)
                os.Exit(1)
            }
        }

        gitkeep := filepath.Join(fixturesDir, ".gitkeep")
        if _, err := os.Stat(gitkeep); os.IsNotExist(err) {
            if err := os.WriteFile(gitkeep, []byte(""), 0o644); err != nil {
                fmt.Fprintf(os.Stderr, "failed to write fixtures placeholder: %v\n", err)
                os.Exit(1)
            }
        }

        fmt.Printf("Initialized eval suite %s at %s\n", name, suiteDir)
    },
}

var runCmd = &cobra.Command{
    Use:   "run <suite>",
    Short: "Run an eval suite",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Running %s...\n", args[0])
    },
}

var scoreCmd = &cobra.Command{
    Use:   "score <run>",
    Short: "Show score breakdown for a run",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("Scoring run %s...\n", args[0])
    },
}

var compareCmd = &cobra.Command{
    Use:   "compare <runA> <runB>",
    Short: "Compare two runs",
    Args:  cobra.ExactArgs(2),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Comparing...")
    },
}

var historyCmd = &cobra.Command{
    Use:   "history <suite>",
    Short: "Show run history for a suite",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Printf("History for %s...\n", args[0])
    },
}

var inspectCmd = &cobra.Command{
    Use:   "inspect <episode>",
    Short: "Show episode details",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Inspecting...")
    },
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
