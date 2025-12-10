#!/bin/bash

echo "Refreshing all Go files for 01-edu checker..."

# Remove all Go files from Git index (but keep them on disk)
git rm --cached $(git ls-files '*.go')

# Re-add all files
git add .

# Create a new commit
git commit -m "Full refresh of Go files for 01-edu checker"

echo "Done! All Go files refreshed."
