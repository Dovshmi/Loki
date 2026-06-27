# 🐍 Loki — Custom Script Header Generator

<div align="center">
  <img src="https://img.shields.io/badge/Bash-Script-4EAA25?style=for-the-badge&logo=gnubash&logoColor=white" alt="Bash" />
  <img src="https://img.shields.io/badge/Figlet-ASCII_Art-7C3AED?style=for-the-badge" alt="Figlet" />
  <img src="https://img.shields.io/badge/Mode-CLI_%2B_TUI-0F766E?style=for-the-badge" alt="CLI and TUI" />
  <img src="https://img.shields.io/badge/Output-Multi_Language-2563EB?style=for-the-badge" alt="Multi-language output" />
  <img src="https://img.shields.io/badge/License-Not_Specified-6B7280?style=for-the-badge" alt="License Not Specified" />
</div>

<div align="center">
  <p><strong>A Bash utility that generates clean starter script headers with Figlet ASCII art, author details, descriptions, optional dates, and language-aware shebangs.</strong></p>
</div>

---

## Overview

**Loki** is a small command-line and text-interface tool for creating polished script headers quickly. It asks for, or accepts through flags, the script name, author name, description, language, Figlet font, and optional timestamp. Loki then creates a starter file with a matching extension, language-aware shebang where relevant, Figlet banner, description line, and executable permission.

It is useful when you create many scripts and want each file to start with a consistent, recognizable, and more personal header.

---

## Demo

### TUI Mode

Use the interactive terminal flow when you want Loki to ask for each value step by step.

```bash
./bashscript/loki.sh --tui
```

<div align="center">
  <img src="https://github.com/Dovshmi/Loki/assets/70490582/df76da56-6a0b-4a15-aa9d-616882fd23d2" alt="Loki TUI mode demo" width="760" />
</div>

### CLI Help and Options

Loki can also run fully from command-line arguments.

<div align="center">
  <img src="https://github.com/Dovshmi/Loki/assets/70490582/5abd2d54-0800-4518-b9c5-25cb5c5ae60b" alt="Loki command-line options" width="760" />
</div>

### Generated Header Examples

<div align="center">
  <img src="https://github.com/Dovshmi/Loki/assets/70490582/32a511cc-b466-461e-abe8-b705d998b5e3" alt="Generated Bash header example" width="760" />
</div>

<div align="center">
  <img src="https://github.com/Dovshmi/Loki/assets/70490582/d49359dc-cda7-4f82-be9e-994d4aeeea0d" alt="Generated Python header example" width="760" />
</div>

---

## Features

- **Interactive TUI mode** for guided header creation.
- **Command-line mode** for fast repeatable generation.
- **Figlet ASCII art** for script name and author branding.
- **Language-aware output** with extension and shebang handling.
- **Optional custom Figlet font** with a default fallback.
- **Optional date/time banner** in the generated header.
- **Automatic executable permission** for the generated file.
- **Extension inference** when the output file name already includes a supported extension.

---

## Supported Languages

Loki currently maps these language names/extensions:

| Language input | Output extension | Shebang / behavior |
| :--- | :--- | :--- |
| `bash`, `sh` | `.sh` | `#!/bin/bash` |
| `python`, `py` | `.py` | `#!/usr/bin/env python` |
| `perl`, `pl` | `.pl` | `#!/usr/bin/perl` |
| `ruby`, `rb` | `.rb` | `#!/usr/bin/env ruby` |
| `php` | `.php` | `#!/usr/bin/env php` |
| `javascript`, `js` | `.js` | `#!/usr/bin/env node` |
| `c` | `.c` | no shebang |
| `cpp`, `c++` | `.cpp` | no runtime shebang |
| `java` | `.java` | no runtime shebang |
| `golang`, `go` | `.go` | no runtime shebang |

---

## Requirements

Required:

- Bash
- Figlet
- Standard Unix tools such as `getopt`, `sed`, `awk`, `date`, and `chmod`

Install Figlet with your package manager:

```bash
# Debian / Ubuntu
sudo apt-get install figlet

# Fedora
sudo dnf install figlet

# Arch Linux
sudo pacman -S figlet
```

---

## Project Structure

```text
Loki/
├── bashscript/
│   └── loki.sh        # Main Loki script
└── README.md          # Project documentation
```

---

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/Dovshmi/Loki.git
cd Loki
```

### 2. Make the script executable

```bash
chmod +x bashscript/loki.sh
```

### 3. Show help

```bash
./bashscript/loki.sh --help
```

---

## Usage

### Interactive Mode

```bash
./bashscript/loki.sh -T
```

or:

```bash
./bashscript/loki.sh --tui
```

### Command-Line Mode

Create a Bash script header:

```bash
./bashscript/loki.sh \
  -a "Rony Shmidov" \
  -o "backup" \
  -d "Backup local project files" \
  -l bash
```

Create a Python script by inferring the language from the output extension:

```bash
./bashscript/loki.sh \
  -a "Rony Shmidov" \
  -o "analyze.py" \
  -d "Analyze project data"
```

Create a Go starter file with a custom Figlet font and date banner:

```bash
./bashscript/loki.sh \
  -a "Rony Shmidov" \
  -o "tool.go" \
  -d "Small Go utility" \
  -f mini \
  -t
```

---

## Options

| Option | Description |
| :--- | :--- |
| `-a`, `--author` | Author name or alias to place in the header. |
| `-o`, `--output` | Output file name. Can include a supported extension. |
| `-d`, `--description` | Description line for the generated script. |
| `-l`, `--language` | Script language when the output extension is not provided. |
| `-f`, `--font` | Figlet font to use. Defaults to `small`. |
| `-t`, `--time` | Add the current date/time banner to the header. |
| `-T`, `--tui` | Start the interactive text user interface. |
| `-h`, `--help` | Show help output. |

---

## Optional Global Shortcut

To run Loki from anywhere, create a symbolic link:

```bash
sudo ln -s "$(pwd)/bashscript/loki.sh" /usr/local/bin/loki
```

Then use:

```bash
loki --help
loki --tui
```

---

## Output Behavior

When Loki runs successfully, it creates a new script file in the current directory and makes it executable.

Example:

```bash
./bashscript/loki.sh -a "Rony Shmidov" -o "hello.sh" -d "Simple hello script"
```

Expected output file:

```text
hello.sh
```

The generated file includes:

- the relevant shebang, where applicable;
- Figlet art for the script name;
- Figlet art for the author line;
- an optional date/time banner;
- a `Description:` line;
- executable permissions.

---

## Use Cases

- Personal script templates.
- Project automation files.
- Educational examples and tutorials.
- Script libraries that need a consistent header style.
- Quick branded starter files for Bash, Python, Go, JavaScript, and other small tools.

---

## Known Limitations

- Existing output files may be overwritten because there is no overwrite confirmation flow yet.
- Figlet must be installed before running Loki.
- Long option parsing depends on `getopt`, so behavior may vary slightly between Linux and macOS environments.
- There is no formal installer script yet.
- No formal `LICENSE` file is currently included in the repository.

---

## Roadmap Ideas

- Add overwrite confirmation before replacing an existing file.
- Add validation for unavailable Figlet fonts.
- Add a safer install script for the global `loki` command.
- Add examples for every supported language.
- Add ShellCheck and automated smoke tests.
- Add a formal open-source license file.

---

## License

No formal `LICENSE` file is currently included in this repository. Add a license file before encouraging reuse, modification, or redistribution by others.

---

## Acknowledgements

- Built around Figlet ASCII art for fast visual script branding.
- Created to make starter scripts feel more personal and easier to identify.

---

<div align="center">
  Built by <strong>Dovshmi</strong><br />
  GitHub: <a href="https://github.com/Dovshmi">@Dovshmi</a>
</div>
