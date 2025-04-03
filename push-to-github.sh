#!/bin/bash

# === USER CONFIGURATION ===
GITHUB_USERNAME="sumitsawant"               # ✅ Replace if needed
REPO_NAME="webhook-orchestrator-go"         # ✅ Match your GitHub repo name
COMMIT_MSG="Initial commit: Webhook Orchestrator in Go"
REMOTE_URL="https://github.com/$GITHUB_USERNAME/$REPO_NAME.git"

# === SCRIPT START ===
echo "⚙️  Setting up Git repository..."

set -e  # Exit on error

# Initialize git if not already
if [ ! -d ".git" ]; then
  git init
else
  echo "✅ Git repo already initialized."
fi

# Check if remote origin exists
if git remote get-url origin >/dev/null 2>&1; then
  echo "🔁 Remote 'origin' already exists. Updating URL..."
  git remote set-url origin "$REMOTE_URL"
else
  echo "🔗 Adding remote 'origin'..."
  git remote add origin "$REMOTE_URL"
fi

# Stage and commit changes
echo "📦 Adding files..."
git add .

if git diff --cached --quiet; then
  echo "⚠️  Nothing new to commit."
else
  git commit -m "$COMMIT_MSG"
fi

# Push to main branch
git branch -M main
echo "🚀 Pushing to GitHub..."
git push -u origin main

echo "✅ Code pushed to: $REMOTE_URL"
