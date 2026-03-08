package fzf

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/silee-tools/jg/internal/entry"
)

// Run launches fzf with the given entries and optional query.
// Returns the selected path or empty string if cancelled.
func Run(entries []entry.Entry, query string) (string, error) {
	fzfPath, err := exec.LookPath("fzf")
	if err != nil {
		return "", fmt.Errorf("fzf not found. Install it: brew install fzf")
	}

	args := []string{
		"--height=40%",
		"--reverse",
		"--no-sort",
		"--select-1",
		"--header=Git Repos",
		"--preview", `git -C {} log --oneline -5 2>/dev/null; echo; echo "branch: $(git -C {} branch --show-current 2>/dev/null)"; echo; git -C {} status --short 2>/dev/null | head -10`,
	}
	if query != "" {
		args = append(args, "--query", query)
	}

	cmd := exec.Command(fzfPath, args...)
	cmd.Stderr = os.Stderr

	var input strings.Builder
	for _, e := range entries {
		fmt.Fprintln(&input, e.Path)
	}
	cmd.Stdin = strings.NewReader(input.String())

	out, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			// fzf exit 1 = no match, exit 130 = cancelled
			if exitErr.ExitCode() == 1 || exitErr.ExitCode() == 130 {
				return "", nil
			}
		}
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
