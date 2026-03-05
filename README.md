# jg

[한국어 (Korean)](docs/README_ko.md)

A frecency-based CLI for quickly jumping to Git repositories.

Ranks your Git repositories by frecency (frequency + recency) and lets you quickly select and navigate to them via fzf.

## Tech Stack

- Go
- fzf (external dependency)

## Features

- **Frecency-based ranking**: Scoring that combines visit frequency and recency
- **Automatic collection**: Automatically records Git repository visits via `chpwd` hook
- **fzf preview**: Shows branch, recent commits, and dirty status in preview
- **Cleanup**: Automatically removes entries for deleted directories

## Getting Started

```bash
# Build
mise run build

# Install (~/.local/bin/jg)
mise run install
```

Add the shell wrapper to your `.zshrc` to use. See source code for detailed setup instructions.
