# Deki

Deki is a small interactive shell written in Go. It reads commands in raw terminal mode, supports a few built-in commands, runs external programs found on `PATH`, and includes tab completion for common commands.

## Features

- Interactive prompt with character-by-character input handling
- Built-in commands: `cd`, `pwd`, `echo`, `type`, and `exit`
- External command execution through the operating system `PATH`
- Output redirection with `>` and `>>`
- Simple quoting and escaping support in the command parser
- Tab completion for built-ins and executable names discovered from `PATH`

## Running

This project is a Go module and the shell entry point lives in `app/main.go`.

```bash
go run ./app
```

If you want to build a binary instead:

```bash
go build -o deki ./app
```

Then run it with:

```bash
./deki
```

## Usage Notes

- `cd` accepts regular paths and also handles `~` for the home directory.
- `pwd` prints the current working directory.
- `echo` prints its arguments joined by spaces.
- `type` reports whether a name is a shell builtin or an executable on `PATH`.
- External commands are executed directly and inherit standard input, output, and error streams.

Examples:

```bash
pwd
echo hello world
type echo ls
cd ..
ls
```

## Project Structure

- `app/main.go` - interactive shell loop, parsing, redirection, and command dispatch
- `internal/parser` - command parsing with quote and escape handling
- `internal/builtins` - implementations of shell built-ins
- `internal/exec` - external command lookup and execution helpers
- `internal/autocomplete/trie` - autocomplete trie and longest-prefix matching



## Requirements

- Go 1.25 or newer
- A terminal that supports raw mode input for the interactive shell