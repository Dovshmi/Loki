# Loki The Custom Script Header Generator

This Bash script is designed to create custom script headers with Figlet ASCII art. It allows users to add a personalized touch to their scripts by generating headers that include script name, author name, description, and Figlet ASCII art. 

## Prerequisites

Before using this script, ensure that Figlet is installed on your system.
Depending on your package manager, you can install Figlet using the following commands:
```bash
sudo apt-get install figlet
```
```bash
sudo dnf install figlet
```
```bash
sudo  pacman -S figlet
```

## Usage

### Interactive Mode
By Using Interactive Mode, You can provide the necessary details for generating the script header.
Like old scripts, you can use the interactive mode to add details to the script header.

```bash
./loki.sh -t
```
Or you could use the -- TUI mode:
```bash
./loki.sh --tui
```
### Command-line Arguments
You can also use command-line arguments to provide the necessary details for generating the script header. Available options are:

- `-a, --author`: Specify the author name.
- `-o, --output`: Specify the output file name.
- `-d, --description`: Provide a description for the script.
- `-l, --language`: Specify the script language (e.g., bash, python).
- `-f, --font`: Optionally specify the Figlet font to use.
- `-t, --tui`: Enable textual user interface mode.

Example:

```bash
./loki.sh -a "John Doe" -o "test.sh" -d "A simple script" -l bash
```

### Examples

#### Example 1: Creating a Bash Script Header
```bash
./loki.sh -a "Jane Smith" -o "my_script.sh" -d "A utility script for file manipulation" -l bash
```

#### Example 2: Generating a Python Script Header
```bash
./loki.sh -a "John Doe" -o "python_script.py" -d "A script to analyze data" -l python
```

#### Example 3: Customizing Font and Using TUI Mode
```bash
./loki.sh -t
```

## Use Cases

- **Personal Scripts**: Customize your personal scripts with unique headers, making them more enjoyable to use and share.
- **Project Scripts**: Enhance the appearance of scripts used in your projects, adding branding and professionalism.
- **Educational Materials**: Create visually appealing script headers for tutorials, guides, or educational materials.
- **Script Libraries**: Incorporate custom headers into script libraries to provide a consistent look and feel.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgements

- This script was inspired by the need for easily creating custom script headers.
- Figlet ASCII art adds a creative touch to scripts, enhancing their appearance and personalization.
