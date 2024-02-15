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
		echo "" # No shebang for C++
		;;
	java)
		echo "java"
		echo "" # No shebang for Java
		;;
	golang | go)
		echo "go"
		echo "" # No shebang for Go
		;;
	*)
		echo "Unsupported language. Exiting..."
		exit 1
		;;
	esac
}

fig_use() {
	script_name=$1
	author_name="$2"
	description="$3"
	script_lang=$4
	echo "$script_name $author_name $description $script_lang"
	current_date=$(date +%c | awk '{printf "%s %s %s %s",$1,$2,$3,$4}' | figlet -f small | sed 's/^/# /')

	file_extension=$(handle_language "$script_lang" | head -n 1)
	shebang=$(handle_language "$script_lang" | tail -n 1)

	echo "$shebang" >"$script_name.$file_extension"
	figlet -f small "$script_name" | sed 's/^/# /' >>"$script_name.$file_extension"
	echo "By  .$author_name" | figlet -f small | sed 's/^/# /' >>"$script_name.$file_extension"
	echo "$current_date" >>"$script_name.$file_extension"
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
	fig_use $script_name "$author_name" "$description" $script_lang
}

# Main function
main() {
	ARGS=$(getopt --options a:o:d:l:t -a -l "author:,output:,description:,language:,tui" -- "$@")
	eval set --"$ARGS"
	echo $ARGS
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
				language=$(handle_language "$checkoutput" | head -n 1)
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
				language=$(handle_language "$2" | head -n 1)
			else
				echo "$language Already set by the file name"
				exit 2
			fi
			shift 2
			;;
		-t | --tui)
			tui="true"
			create_script_header
			exit 1
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
}

# Execute main function
main "$@"
