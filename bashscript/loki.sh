#!/bin/bash

# Check if Figlet is installed
if ! command -v figlet &>/dev/null; then
	echo "Figlet is not installed. Please install Figlet first."
	exit 1
fi

# Function to handle language cases
handle_language() {
	case "$1" in
	bash | sh)
		echo "sh"
		echo "#!/bin/bash"
		;;
	python | py)
		echo "py"
		echo "#!/usr/bin/env python"
		;;
	perl | pl)
		echo "pl"
		echo "#!/usr/bin/perl"
		;;
	ruby | rb)
		echo "rb"
		echo "#!/usr/bin/env ruby"
		;;
	php)
		echo "php"
		echo "#!/usr/bin/env php"
		;;
	javascript | js)
		echo "js"
		echo "#!/usr/bin/env node"
		;;
	c | c)
		echo "c"
		echo "" # No shebang for C
		;;
	cpp | c++)
		echo "cpp"
		echo "#" # No shebang for C++
		;;
	java)
		echo "java"
		echo "#" # No shebang for Java
		;;
	golang | go)
		echo "go"
		echo "#" # No shebang for Go
		;;
	*)
		echo "false"
		exit 1
		;;
	esac
}

fig_use() {
	script_name=$1
	author_name="$2"
	description="$3"
	script_lang=$4
	fig_font="$5"
	time_line="$6"
	#echo "$script_name $author_name $description $script_lang $fig_font $time_line"
	file_extension=$(handle_language "$script_lang" | head -n 1)
	shebang=$(handle_language "$script_lang" | tail -n 1)

	echo "$shebang" >"$script_name.$file_extension"
	figlet -f "$fig_font" "$script_name" | sed 's/^/# /' >>"$script_name.$file_extension"
	echo "By  .$author_name" | figlet -f "$fig_font" | sed 's/^/# /' >>"$script_name.$file_extension"
	if [[ "$time_line" == "true" ]]; then
		current_date=$(date +%c | awk '{printf "%s %s %s %s",$1,$2,$3,$4}' | figlet -f "$fig_font" | sed 's/^/# /')
		echo "$current_date" >>"$script_name.$file_extension"
	fi
	echo "# Description: $description" >>"$script_name.$file_extension"
	chmod +x "$script_name.$file_extension"
	echo "Custom script header created successfully in $script_name.$file_extension!"

}

# Function to create custom script header
create_script_header() {
	clear
	echo "Creating Custom Script Header..."
	read -p "Enter your script name: " script_name
	read -p "Enter your name or alias: " author_name
	read -p "Enter script description: " description
	read -p "Enter script language (bash, python, perl, ruby, php, javascript, c, cpp, java, go): " script_lang
	read -p "Enter Figlet font (optional, press Enter for default 'small'): " fig_font
	read -p "Enter add date (y/n) (optional, press Enter for default 'y'): " time_line
	case "$time_line" in
	y | Y | yes | Yes | YES)
		time_line="true"
		;;
	n | N | no | No | NO)
		time_line="false"
		;;
	*)
		time_line="true"
		;;
	esac
	fig_font=${fig_font:-small} # Set default font to small if empty
	fig_use "$script_name" "$author_name" "$description" "$script_lang" "$fig_font" "$time_line"
}
# Function to display usage information
display_usage() {
	echo "Usage: $0 [OPTIONS]"
	echo "Options:"
	echo "  -a, --author      Specify the author name."
	echo "  -o, --output      Specify the output file name."
	echo "  -d, --description Specify the script description."
	echo "  -l, --language    Specify the script language."
	echo "  -f, --font        Specify the Figlet font."
	echo "  -t, --time        Add the current date to the header."
	echo "  -T, --tui         Use TUI (Text-based User Interface) to interactively provide script details."
	echo "  -h, --help        Display this help message."
	exit 0
}

# Main function
main() {
	ARGS=$(getopt --options a:o:d:l:f:htT -a -l "author:,output:,description:,language:,font:,time,help,tui" -- "$@")
	eval set --"$ARGS"
	#echo $ARGS
	Check=$(echo $ARGS | awk '{print $NF}')
	if [[ ! "$Check" == "--" ]]; then
		echo "Invalid command. Exiting..."
		exit 1
	fi

	author="false"
	output="false"
	description="false"
	language="false"
	font="false"
	tui="false"
	timeline="false"
	handle_help="false"
	while true; do
		case "$1" in
		-a | --author)
			author=$2
			shift 2
			;;
		-o | --output)
			output=$2
			dot=$(echo "$output" | tr -s "." " ")
			checkoutput=$(echo "$dot" | awk '{print $2}')
			if [[ -n "$checkoutput" ]]; then
				language=$(handle_language "$checkoutput" | head -n 1)
				output=$(echo "$dot" | awk '{print $1}')
			fi
			shift 2
			;;
		-d | --description)
			description=$2
			shift 2
			;;
		-l | --language)
			if [[ "$language" == "false" ]]; then
				language=$(handle_language "$2" | head -n 1)
			else
				echo "$language Already set by the file name"
				exit 2
			fi
			shift 2
			;;
		-f | --font)
			font=$2
			shift 2
			;;
		-t | --time)
			timeline="true"
			shift
			;;
		-h | --help)
			handle_help="true"
			display_usage
			exit 1
			;;
		-T | --tui)
			tui="true"
			create_script_header
			exit 1
			;;
		--)
			break
			;;
		*)
			echo "Invalid option: $1"
			display_usage
			exit 1
			;;
		esac
	done
	#echo "tui: $tui, author: $author, description: $description, language: $language, output: $output, font: $font"

	if [[ "$author" == "false" && "$tui" == "false" ]]; then
		echo "No author name specified. Exiting..."
		display_usage
		exit 1
	fi
	if [[ "$output" == "false" && "$tui" == "false" ]]; then
		echo "No output file specified. Exiting..."
		display_usage
		exit 1
	fi
	if [[ ! "$output" == "false" && "$language" == "false" ]]; then
		echo "No language specified. Exiting..."
		display_usage
	fi

	if [[ "$description" == "false" ]]; then
		description="you could write a small poem"
	fi
	if [[ "$font" == "false" ]]; then
		font="small"
	fi
	fig_use "$output" "$author" "$description" "$language" "$font" "$timeline"

}

# Execute main function
main "$@"
