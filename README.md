# ASCII Art Generator (Go)

A command-line program written in Go that converts text into its ASCII-art graphical representation using predefined banner templates.

---

## Features

- Supports:
  - Uppercase letters
  - Lowercase letters
  - Numbers
  - Special characters
  - Spaces
  - `\n` for multi-line input
- Multiple banner styles:
  - `standard.txt`
  - `shadow.txt`
  - `thinkertoy.txt`
- Modular structure using Go packages
- Uses only Go standard library

---

## Project Structure

ascii-art/
├── main.go
├── go.mod
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
├── funcs/
│ ├── helper.go
│ ├── printer.go
│ └── splitendl.go
└── .gitignore


---

## Build

go build -o ascii-art


---

## Usage

### Default banner (standard)

./ascii-art "Hello"


### With specific banner

./ascii-art "Hello" shadow.txt

./ascii-art "Hello" thinkertoy.txt


If no banner is specified, `standard.txt` is used by default.

---

## Multi-line Input

./ascii-art "Hello\n\nWorld"


To verify correct newline handling:

./ascii-art "Hello\n\nWorld" | cat -e


---

## How It Works

- Reads banner file
- Normalizes line endings (`\r\n` → `\n`)
- Converts `\n` from user input into real newline characters
- Splits input into lines
- Maps each character using:

startLine = (ASCII - 32) * 9


Each character has a height of 8 lines in the banner file.

---

## Requirements

- Go 1.20+
- No external dependencies

---

## License

Educational project.
