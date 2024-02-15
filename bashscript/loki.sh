#!/bin/bash

# Check if Figlet is installed
if ! command -v figlet &>/dev/null; then
	echo "Figlet is not installed. Please install Figlet first."
	exit 1
fi

# Function to create custom script header
create_script_header() {
	clear
	echo "Creating Custom Script Header..."
	read -p "Enter your script name: " script_name
	read -p "Enter your name or alias: " author_name
	read -p "Enter script description: " description
	read -p "Enter script language (bash, python, perl, ruby, php, javascript, c, cpp, java, go): " script_lang
	current_date=$(date +%c | awk '{printf "%s %s %s %s",$1,$2,$3,$4}' | figlet -f small | sed 's/^/# /')
	case "$script_lang" in
	bash)
		file_extension="sh"
		shebang="#!/bin/bash"
		;;
	python)
		file_extension="py"
		shebang="#!/usr/bin/env python"
		;;
	perl)
		file_extension="pl"
		shebang="#!/usr/bin/perl"
		;;
	ruby)
		file_extension="rb"
		shebang="#!/usr/bin/env ruby"
		;;
	php)
		file_extension="php"
		shebang="#!/usr/bin/env php"
		;;
	javascript)
		file_extension="js"
		shebang="#!/usr/bin/env node"
		;;
	c)
		file_extension="c"
		shebang=""
		;;
	cpp)
		file_extension="cpp"
		shebang=""
		;;
	java)
		file_extension="java"
		shebang=""
		;;
	go)
		file_extension="go"
		shebang=""
		;;
	*)
		echo "Unsupported language. Exiting..."
		exit 1
		;;
	esac
	echo "$shebang" >"$script_name.$file_extension"
	figlet -f small "$script_name" | sed 's/^/# /' >>"$script_name.$file_extension"
	echo "By  .$author_name" | figlet -f small | sed 's/^/# /' >>"$script_name.$file_extension"
	echo "$current_date" >>"$script_name.$file_extension"
	echo "# Description: $description" >>"$script_name.$file_extension"
	chmod +x "$script_name.$file_extension"
	echo "Custom script header created successfully in $script_name.$file_extension!"
}

# Main function
main() {
	#	echo "$@"
	ARGS=$(getopt --options a:o:d:l:t -a -l "author:,output:,description:,language:,tui" -- "$@")
	eval set --"$ARGS"
	echo "$ARGS"
	Check=$(echo $ARGS | awk '{print $NF}')
	if [[ ! "$Check" == "--" ]]; then
		echo "Invalid command. Exiting..."
		exit 1
	fi
	author="false"
	output="false"
	description="false"
	language="false"
	tui="false"
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
				language="$checkoutput"
				echo "$language Loaded"
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
				language=$2
			else
				echo "$language Already set by the file name"
				exit 2
			fi
			shift 2
			;;
		-t | --tui)
			tui="true"
			create_script_header
			break
			;;
		--)
			break
			;;
		*)
			echo "Invalid option: $1"
			exit 1
			;;
		esac
	done
	echo "tui: $tui, author: $author, description: $description, language: $language, output: $output"
	#create_script_header
}

# Execute main function
main "$@"
