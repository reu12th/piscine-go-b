# Script to auto add, commit, and push each changed file separately
# Works on Windows PowerShell

# Get changed or untracked files
$files = git status --porcelain | ForEach-Object {
    ($_ -split "\s+")[1]
}

if (!$files) {
    Write-Host "No changes to commit."
    exit
}

foreach ($file in $files) {
    Write-Host "Process file: $file ? (y = yes, n = skip)"
    $ans = Read-Host ">"

    if ($ans -ne "y") {
        Write-Host "Skipping $file"
        continue
    }

    # Extract filename without extension (default commit message)
    $filename = Split-Path $file -Leaf
    $default_msg = [System.IO.Path]::GetFileNameWithoutExtension($filename)

    Write-Host "Press 'y' to use default commit message: '$default_msg'"
    Write-Host "Or type your custom commit message:"
    $msg_input = Read-Host ">"

    if ($msg_input -eq "y") {
        $commit_msg = $default_msg
    } else {
        $commit_msg = $msg_input
    }

    git add "$file"
    git commit -m "$commit_msg"
}

git push
Write-Host "All done!"
