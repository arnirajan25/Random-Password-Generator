# Random Password Generator

A lightweight, fast command-line random password generator written in Go.

## Features

- Generate secure random passwords with customizable length
- Optional inclusion of special characters
- Built-in password analysis showing character composition
- Simple CLI interface with intuitive flags

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/pswgnrtr.git
cd pswgnrtr

# Run directly
go run main.go --length 20 --special

# Or build and install
go build -o pswgnrtr main.go
./pswgnrtr --length 20 --special
```

## Usage

```bash
go run main.go [options]
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `--length <number>` | Length of the generated password | `16` |
| `--special` | Include special characters (`!@#$%^&*()_+-=[]{}|;:,.<>?`) | `false` |

### Examples

```bash
# Generate a 16-character password (default)
go run main.go --length 16

# Generate a 32-character password with special characters
go run main.go --length 32 --special

# Generate a simple 12-character alphanumeric password
go run main.go --length 12
```

## Sample Output

```
Generated Password (length 32):
----------------------------------------
K9#mP$vL2@xQ!wR5&nB8*eT1^hJ4
----------------------------------------
Lowercase: true | Uppercase: true | Digits: true | Special: true
```

## Character Sets

- Lowercase: `abcdefghijklmnopqrstuvwxyz`
- Uppercase: `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
- Digits: `0123456789`
- Special: `!@#$%^&*()_+-=[]{}|;:,.<>?`

## Requirements

- Go 1.16 or higher

## License

[MIT](LICENSE)

---

Feel free to customize the repository URL, add badges, or adjust the license section as needed.# Random Password Generator

A lightweight, fast command-line random password generator written in Go.

## Features

- Generate secure random passwords with customizable length
- Optional inclusion of special characters
- Built-in password analysis showing character composition
- Simple CLI interface with intuitive flags

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/pswgnrtr.git
cd pswgnrtr

# Run directly
go run main.go --length 20 --special

# Or build and install
go build -o pswgnrtr main.go
./pswgnrtr --length 20 --special
```

## Usage

```bash
go run main.go [options]
```

### Options

| Flag | Description | Default |
|------|-------------|---------|
| `--length <number>` | Length of the generated password | `16` |
| `--special` | Include special characters (`!@#$%^&*()_+-=[]{}|;:,.<>?`) | `false` |

### Examples

```bash
# Generate a 16-character password (default)
go run main.go --length 16

# Generate a 32-character password with special characters
go run main.go --length 32 --special

# Generate a simple 12-character alphanumeric password
go run main.go --length 12
```

## Sample Output

```
Generated Password (length 32):
----------------------------------------
K9#mP$vL2@xQ!wR5&nB8*eT1^hJ4
----------------------------------------
Lowercase: true | Uppercase: true | Digits: true | Special: true
```

## Character Sets

- Lowercase: `abcdefghijklmnopqrstuvwxyz`
- **Uppercase: `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
- Digits: `0123456789`
- Special: `!@#$%^&*()_+-=[]{}|;:,.<>?`

## Requirements

- Go 1.16 or higher

## License

[MIT](LICENSE)

---

Feel free to customize the repository URL, add badges, or adjust the license section as needed.
