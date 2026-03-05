# jg

[한국어 (Korean)](docs/README_ko.md)

A frecency-based CLI for quickly jumping to Git repositories.

Ranks your Git repositories by frecency (frequency + recency) and lets you quickly select and navigate to them via fzf.

## Installation

```bash
brew install silee-tools/tap/jg
```

This automatically installs `fzf` as a dependency.

## Shell Setup

Add to your `~/.zshrc`:

```zsh
eval "$(jg init zsh)"
```

Or for Bash, add to `~/.bashrc`:

```bash
eval "$(jg init bash)"
```

## Usage

```bash
jg              # Interactive jump with fzf
jg <query>      # Jump with pre-filtered query
jg -l           # List all tracked repos with scores
jg --clean      # Remove stale entries
jg --remove .   # Remove current directory from tracking
```

Once shell integration is set up, repositories are automatically tracked as you `cd` into them.

## Features

- **Frecency-based ranking**: Scoring that combines visit frequency and recency
- **Automatic collection**: Automatically records Git repository visits via shell hook
- **fzf preview**: Shows branch, recent commits, and dirty status in preview
- **Cleanup**: Automatically removes entries for deleted directories
- **Multi-shell support**: Works with both Zsh and Bash

## Development

```bash
mise run build      # Build
mise run test       # Run tests
mise run install    # Install to ~/.local/bin/jg
```
