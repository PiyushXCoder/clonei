<img width="282" height="69" alt="Group 10" src="https://github.com/user-attachments/assets/685c4ee8-c11c-4e41-9274-21b43b4c2d10" />

# iclone

A CLI tool to clone Git repositories and automatically install dependencies in a single shot.

## Installation

```bash
curl -fsSL https://raw.githubusercontent.com/soft4dev/iclone/refs/heads/main/scripts/install.sh | sh
```

## Usage

### Basic usage (auto-detect project type)

```bash
iclone <repository-url>
```

Example:

```bash
iclone https://github.com/username/my-project.git
```

### Specify project type manually

```bash
iclone -p <project-type> <repository-url>
```

Example:

```bash
iclone -p npm https://github.com/username/my-project.git
```

## Supported Project Types

- **Node.js**: npm, pnpm
- more to be added...

## How it works

1. Clones the specified Git repository
2. Detects the project type (or uses the specified type)
3. Automatically installs dependencies based on the project type

## License

See [LICENSE](LICENSE) for details.
