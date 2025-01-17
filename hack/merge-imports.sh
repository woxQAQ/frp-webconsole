#!/bin/bash

# Check if a file is provided
if [ -z "$1" ]; then
    echo "Usage: $0 <filename>"
    exit 1
fi

# Temporary file for processing
temp_file=$(mktemp)

# Process the file
awk '
BEGIN { in_import_block = 0 }

# Detect the start of an import block
/^import \(/ { 
    in_import_block = 1
    print
    next
}

# Detect the end of an import block
/^\)/ && in_import_block { 
    in_import_block = 0
    print
    next
}

# If inside an import block, remove extra newlines
in_import_block {
    if ($0 != "") {
        print
    }
    next
}

# Print non-import lines
{ print }
' "$1" > "$temp_file"

# Replace the original file with the processed file
mv "$temp_file" "$1"