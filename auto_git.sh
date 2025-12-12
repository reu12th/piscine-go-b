b#!/bin/bash

# Script to auto add, commit, and push each changed file separately
# Default commit message = filename without extension
# If user inputs 'y' → use default
# If user types anything else → use that as custom message
# Allows skipping files

# Get list of changed or untracked files
files=$(git status --porcelain | awk '{print $2}')

if [ -z "$files" ]; then
    echo "No changes to commit."
    exit 0
fi

for file in $files; do
    echo "Process file: $file ? (y = yes, n = skip)"
    read -p "> " ans

    if [ "$ans" != "y" ]; then
        echo "Skipping $file"
        continue
    fi

    # Extract filename without extension (default commit message)
    filename=$(basename "$file")
    default_msg="${filename%.*}"

    echo "Press 'y' to use default commit message: '$default_msg'"
    echo "Or type your custom commit message:"
    read -p "> " msg_input

    if [ "$msg_input" = "y" ]; then
        commit_msg="$default_msg"
    else
        commit_msg="$msg_input"
    fi

    git add "$file"
    git commit -m "$commit_msg"
done

git push
echo "All done!"
