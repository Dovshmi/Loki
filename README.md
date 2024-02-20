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
sudo pacman -S figlet
```

## Usage

### Interactive Mode
By Using Interactive Mode, You can provide the necessary details for generating the script header.
Like old scripts, you can use the interactive mode to add details to the script header.

```bash
./loki.sh -T
```
Or you could use the -- TUI mode:
```bash
./loki.sh --tui
```
![2024-02-20_11-08-15](https://github.com/Dovshmi/Loki/assets/70490582/df76da56-6a0b-4a15-aa9d-616882fd23d2)

### Command-line Arguments
You can also use command-line arguments to provide the necessary details for generating the script header. Available options are:

- `-a, --author`: Specify the author name.
- `-o, --output`: Specify the output file name.
- `-d, --description`: Provide a description for the script.
- `-l, --language`: Specify the script language (e.g., bash, python).
- `-f, --font`: Optionally specify the Figlet font to use.
- `-t, --time`: Optionally specify the time of day to use in the header.
- `-T, --tui`: Enable textual user interface mode.
- `-h, --help`: Display help information.

![2024-02-20_11-04-54](https://github.com/Dovshmi/Loki/assets/70490582/5abd2d54-0800-4518-b9c5-25cb5c5ae60b)

### Examples

#### Example 1: Creating a Bash Script Header Using .sh Extension 
```bash
./loki.sh -a "Jane Smith" -o "my_script.sh" -d "A utility script for file manipulation"
```
![2024-02-20_11-06-22](https://github.com/Dovshmi/Loki/assets/70490582/32a511cc-b466-461e-abe8-b705d998b5e3)

#### Example 2: Generating a Python Script Header Using -l Flag 
```bash
./loki.sh -a "John Doe" -o "my_script" -d "A script to analyze data" -l python
```
![2024-02-20_11-07-16](https://github.com/Dovshmi/Loki/assets/70490582/d49359dc-cda7-4f82-be9e-994d4aeeea0d)

#### Example 3: Generating a Go Script Header Using -f Flag For figlet Font With -t Flag For Date
```bash
./loki.sh -a "John Doe" -o "my_script.go" -d "simple script" -f "mini" -t
```

## Use Cases

- **Personal Scripts**: Customize your personal scripts with unique headers, making them more enjoyable to use and share.
- **Project Scripts**: Enhance the appearance of scripts used in your projects, adding branding and professionalism.
- **Educational Materials**: Create visually appealing script headers for tutorials, guides, or educational materials.
- **Script Libraries**: Incorporate custom headers into script libraries to provide a consistent look and feel.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
Made By Dovshmi
## Acknowledgements

- This script was inspired by the need for easily creating custom script headers.
- Figlet ASCII art adds a creative touch to scripts, enhancing their appearance and personalization.
