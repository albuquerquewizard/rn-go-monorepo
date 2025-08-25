#!/bin/bash

# Full-Stack Boilerplate Project Creator
# This script creates a new project from the boilerplate template

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_status() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check if project name is provided
if [ $# -eq 0 ]; then
    print_error "Project name is required!"
    echo "Usage: $0 <project-name>"
    echo "Example: $0 my-awesome-project"
    exit 1
fi

PROJECT_NAME="$1"
PROJECT_DIR="../$PROJECT_NAME"
CURRENT_DIR=$(pwd)

# Validate project name
if [[ ! "$PROJECT_NAME" =~ ^[a-zA-Z0-9_-]+$ ]]; then
    print_error "Invalid project name! Use only letters, numbers, hyphens, and underscores."
    exit 1
fi

# Check if project directory already exists
if [ -d "$PROJECT_DIR" ]; then
    print_error "Project directory '$PROJECT_DIR' already exists!"
    echo "Please choose a different project name or remove the existing directory."
    exit 1
fi

print_status "Creating new project: $PROJECT_NAME"
print_status "Target directory: $PROJECT_DIR"

# Create project directory
mkdir -p "$PROJECT_DIR"

# Copy the entire monorepo structure
print_status "Copying boilerplate files..."
cp -r . "$PROJECT_DIR/"

# Remove unnecessary files from the new project
cd "$PROJECT_DIR"
rm -rf .git
rm -rf scripts/create-project.sh
rm -rf tmp/
rm -rf node_modules/

# Update Go module name
if [ -f "apps/backend/go.mod" ]; then
    print_status "Updating Go module name..."
    sed -i.bak "s|github.com/albuquerquewizard/monorepo/backend|github.com/$(whoami)/$PROJECT_NAME/backend|g" "apps/backend/go.mod"
    rm "apps/backend/go.mod.bak"
fi

# Update mobile app configuration
if [ -f "apps/mobile/app.json" ]; then
    print_status "Updating mobile app configuration..."
    # Update app name
    sed -i.bak "s/\"name\": \"mobile\"/\"name\": \"$PROJECT_NAME\"/g" "apps/mobile/app.json"
    # Update slug
    sed -i.bak "s/\"slug\": \"mobile\"/\"slug\": \"$PROJECT_NAME\"/g" "apps/mobile/app.json"
    rm "apps/mobile/app.json.bak"
fi

if [ -f "apps/mobile/package.json" ]; then
    print_status "Updating package.json..."
    # Update package name
    sed -i.bak "s/\"name\": \"mobile\"/\"name\": \"$PROJECT_NAME\"/g" "apps/mobile/package.json"
    # Update description
    sed -i.bak "s/\"description\": \"Mobile React Native Boilerplate\"/\"description\": \"$PROJECT_NAME - Mobile App\"/g" "apps/mobile/package.json"
    rm "apps/mobile/package.json.bak"
fi

# Update README
if [ -f "README.md" ]; then
    print_status "Updating README..."
    sed -i.bak "s/Full-Stack Boilerplate Monorepo/$PROJECT_NAME - Full-Stack Project/g" "README.md"
    sed -i.bak "s/A production-ready monorepo template containing/A production-ready project containing/g" "README.md"
    rm "README.md.bak"
fi

# Create .env template for backend
if [ -f "apps/backend/.env.example" ]; then
    print_status "Creating .env template..."
    cp "apps/backend/.env.example" "apps/backend/.env"
    print_warning "Please update apps/backend/.env with your database credentials!"
fi

# Create .gitignore if it doesn't exist
if [ ! -f ".gitignore" ]; then
    print_status "Creating .gitignore..."
    cat > .gitignore << 'EOF'
# Dependencies
node_modules/
vendor/

# Environment files
.env
.env.local
.env.*.local

# Build outputs
dist/
build/
*.exe
*.dll
*.so
*.dylib

# IDE files
.vscode/
.idea/
*.swp
*.swo
*~

# OS files
.DS_Store
Thumbs.db

# Logs
*.log
logs/

# Temporary files
tmp/
temp/

# Database
*.db
*.sqlite

# Go
*.test
*.out

# Mobile
.expo/
*.jks
*.p8
*.p12
*.key
*.mobileprovision
*.orig.*
web-build/

# Docker
.dockerignore
EOF
fi

# Initialize git repository
print_status "Initializing git repository..."
git init
git add .
git commit -m "Initial commit: $PROJECT_NAME project created from boilerplate"

print_success "Project '$PROJECT_NAME' created successfully!"
echo ""
echo "Next steps:"
echo "1. cd $PROJECT_DIR"
echo "2. Update database configuration in apps/backend/.env"
echo "3. Customize the codebase for your specific needs"
echo "4. Start development:"
echo "   - Backend: cd apps/backend && task run"
echo "   - Mobile: cd apps/mobile && npm start"
echo ""
echo "Happy coding! 🚀"
