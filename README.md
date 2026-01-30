# mvr

A simple CLI tool to rename files using Go regex (RE2) patterns or plain string replacement.

## Goals

- Rename one or more files by matching a regex pattern and applying a replacement string.
- Support dry-run mode to preview renames without changing the filesystem.
- Support literal string replacement (no regex) via the `-x` flag.

## Installation

```bash
make build
```

This produces `bin/mvr`. Alternatively:

```bash
go build -o bin/mvr .
```

## Usage

```
mvr [options] <regex string> <regex replacement string> [files]*
```

### Options

| Flag | Description |
|------|-------------|
| `-d` | Dry run only: print renames without performing them |
| `-x` | Disable regex; use plain string replacement only |
| `-h` | Show help message and regex quick reference |

### Examples

Change file extension from `.txt` to `.md`:

```bash
mvr '\.txt$' '.md' *.txt
```

Remove a prefix from filenames:

```bash
mvr '^report_' '' report_jan.txt report_feb.txt
```

Capture and reorder: swap prefix and suffix (e.g. `foo-123` → `123-foo`):

```bash
mvr '^(.+)-(.+)$' '${2}-${1}' foo-123 bar-456
```

Literal replacement (no regex):

```bash
mvr -x "oldname" "newname" oldname.txt
```

Dry run to preview:

```bash
mvr -d '\.log$' '.bak' *.log
```

## Regex Reference (Go/RE2)

mvr uses Go’s [`regexp`](https://pkg.go.dev/regexp) package, which implements [RE2](https://github.com/google/re2/wiki/Syntax) syntax.

### Metacharacters

| Pattern | Meaning |
|---------|---------|
| `.` | Any single character |
| `^` | Start of string |
| `$` | End of string |
| `*` | Zero or more (greedy) |
| `+` | One or more (greedy) |
| `?` | Zero or one |
| `\|` | Alternation (OR) |

### Character classes

| Pattern | Meaning |
|---------|---------|
| `[abc]` | Any of a, b, c |
| `[^abc]` | Any character except a, b, c |
| `[a-z]` | Character range |
| `\d` | Digit `[0-9]` |
| `\D` | Non-digit |
| `\w` | Word character `[A-Za-z0-9_]` |
| `\W` | Non-word character |
| `\s` | Whitespace |
| `\S` | Non-whitespace |

### Groups and replacement

| Pattern | Meaning |
|---------|---------|
| `(...)` | Capturing group |
| `(?:...)` | Non-capturing group |
| `$1`, `$2`, ... | First, second, … captured group in replacement |
| `${1}`, `${2}`, ... | Same, with braces (e.g. for `$10`) |

### Escaping

Literal `.`, `*`, `+`, `?`, `[`, `]`, `(`, `)`, `{`, `}`, `|`, `\`, `^`, `$` must be escaped with `\` in the pattern (e.g. `\.` for a literal dot).

### Documentation

- [Go `regexp` package](https://pkg.go.dev/regexp) — API and overview
- [Go `regexp/syntax`](https://pkg.go.dev/regexp/syntax) — full RE2 syntax reference
- [RE2 syntax (wiki)](https://github.com/google/re2/wiki/Syntax) — RE2 pattern syntax
