#!/bin/bash

# === USER CONFIGURATION ===
GITHUB_USERNAME="sumitsawant"                     # ✅ Change if needed
REPO_NAME="webhook-orchestrator-go"               # ✅ Must match GitHub repo name
COMMIT_MSG="Initial commit: Webhook Orchestrator in Go"
REMOTE_URL="https://github.com/$GITHUB_USERNAME/$REPO_NAME.git"

# === SCRIPT START ===
echo "⚙️  Setting up Git repository..."
set -e

# Initialize git if needed
if [ ! -d ".git" ]; then
  git init
else
  echo "✅ Git repo already initialized."
fi

# Handle origin
if git remote get-url origin >/dev/null 2>&1; then
  echo "🔁 Remote 'origin' exists. Updating URL if needed..."
  git remote set-url origin "$REMOTE_URL"
else
  echo "🔗 Adding remote 'origin'..."
  git remote add origin "$REMOTE_URL"
fi

# Stage and commit
echo "📦 Staging changes..."
git add .

if git diff --cached --quiet; then
  echo "⚠️  Nothing new to commit."
else
  git commit -m "$COMMIT_MSG"
fi

# Set branch
git branch -M main

# Pull remote if needed (avoids non-fast-forward push)
echo "🔄 Syncing with remote 'main' branch..."
if git ls-remote --heads origin main | grep main >/dev/null; then
  git pull --rebase origin main || {
    echo "❌ Rebase failed. Resolve conflicts and try again."
    exit 1
  }
fi

# Push to GitHub
echo "🚀 Pushing to GitHub..."
git push -u origin main

echo "✅ Code pushed successfully: $REMOTE_URL"
