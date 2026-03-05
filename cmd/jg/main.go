package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/silee-tools/jg/internal/entry"
	"github.com/silee-tools/jg/internal/frecency"
	"github.com/silee-tools/jg/internal/fzf"
	"github.com/silee-tools/jg/internal/shell"
)

var (
	addPath    = flag.String("add", "", "Add/update entry for path")
	removePath = flag.String("remove", "", "Remove entry for path")
	listFlag   = flag.Bool("l", false, "List all repos with frecency scores")
	cleanFlag  = flag.Bool("clean", false, "Remove entries for non-existent directories")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: jg [options] [query...]\n\n")
		fmt.Fprintf(os.Stderr, "A frecency-based CLI for quickly jumping to Git repositories.\n\n")
		fmt.Fprintf(os.Stderr, "Commands:\n")
		fmt.Fprintf(os.Stderr, "  jg init <shell>    Output shell integration code (zsh, bash)\n")
		fmt.Fprintf(os.Stderr, "  jg [query...]      Interactive jump with fzf\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	// Handle "init" subcommand before flag.Parse
	if len(os.Args) >= 2 && os.Args[1] == "init" {
		runInit(os.Args[2:])
		return
	}

	flag.Parse()

	switch {
	case *addPath != "":
		runAdd(*addPath)
	case *removePath != "":
		runRemove(*removePath)
	case *listFlag:
		runList()
	case *cleanFlag:
		runClean()
	default:
		runJump(flag.Args())
	}
}

func runInit(args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: jg init <shell>  (zsh, bash)")
		os.Exit(1)
	}
	code, err := shell.Init(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(code)
}

func runAdd(path string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return // silent fail, called from hook
	}

	// Check if path is inside a git repo
	cmd := exec.Command("git", "-C", absPath, "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		return // not a git repo, silently ignore
	}

	repoRoot := strings.TrimSpace(string(out))
	if err := entry.AddOrUpdate(repoRoot); err != nil {
		return // silent fail
	}
}

func runRemove(path string) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	removed, err := entry.Remove(absPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if removed {
		fmt.Fprintf(os.Stderr, "Removed: %s\n", absPath)
	} else {
		fmt.Fprintf(os.Stderr, "Not found: %s\n", absPath)
	}
}

func runList() {
	entries, err := entry.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if len(entries) == 0 {
		fmt.Fprintln(os.Stderr, "No entries. cd into git repos to start tracking.")
		return
	}

	sorted := frecency.Sort(entries)
	now := time.Now().Unix()
	for _, e := range sorted {
		score := frecency.Score(e.Rank, e.Timestamp, now)
		fmt.Printf("%8.1f  %4.0f  %s\n", score, e.Rank, e.Path)
	}
}

func runClean() {
	removed, err := entry.Clean()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Cleaned %d stale entries.\n", removed)
}

func runJump(queryArgs []string) {
	entries, err := entry.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Auto-clean stale entries
	var valid []entry.Entry
	for _, e := range entries {
		info, statErr := os.Stat(e.Path)
		if statErr == nil && info.IsDir() {
			valid = append(valid, e)
		}
	}
	if len(valid) != len(entries) {
		_ = entry.Save(valid)
	}

	if len(valid) == 0 {
		fmt.Fprintln(os.Stderr, "No entries. cd into git repos to start tracking.")
		os.Exit(0)
	}

	sorted := frecency.Sort(valid)
	query := strings.Join(queryArgs, " ")

	selected, err := fzf.Run(sorted, query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if selected == "" {
		os.Exit(1) // cancelled, shell wrapper won't cd
	}

	fmt.Println(selected)
}
