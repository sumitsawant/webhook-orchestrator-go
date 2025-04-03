#!/bin/bash

# === USER CONFIGURATION ===
GITHUB_USERNAME="your-username"             # Replace with your GitHub username
REPO_NAME="webhook-orchestrator-go"         # Must match the GitHub repo you created
COMMIT_MSG="Initial commit: Webhook Orchestrator in Go"

# === SCRIPT START ===
echo "⚙️  Setting up Git repository..."

# Exit on error
set -e

# Initialize Git repo
git init

# Add remote
git remote add origin https://github.com/$GITHUB_USERNAME/$REPO_NAME.git

# Stage and commit
git add .
git commit -m "$COMMIT_MSG"

# Set default branch
git branch -M main

# Push to GitHub
echo "🚀 Pushing to GitHub..."
git push -u origin main

echo "✅ Code pushed to: https://github.com/$GITHUB_USERNAME/$REPO_NAME"

