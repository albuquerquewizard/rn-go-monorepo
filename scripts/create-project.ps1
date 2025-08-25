# Full-Stack Boilerplate Project Creator for Windows PowerShell
# This script creates a new project from the boilerplate template

param(
    [Parameter(Mandatory=$true)]
    [string]$ProjectName
)

# Colors for output
$Red = "Red"
$Green = "Green"
$Yellow = "Yellow"
$Blue = "Blue"

# Function to print colored output
function Write-Status {
    param([string]$Message)
    Write-Host "[INFO] $Message" -ForegroundColor $Blue
}

function Write-Success {
    param([string]$Message)
    Write-Host "[SUCCESS] $Message" -ForegroundColor $Green
}

function Write-Warning {
    param([string]$Message)
    Write-Host "[WARNING] $Message" -ForegroundColor $Yellow
}

function Write-Error {
    param([string]$Message)
    Write-Host "[ERROR] $Message" -ForegroundColor $Red
}

# Validate project name
if ($ProjectName -notmatch '^[a-zA-Z0-9_-]+$') {
    Write-Error "Invalid project name! Use only letters, numbers, hyphens, and underscores."
    exit 1
}

$ProjectDir = Join-Path (Split-Path (Get-Location) -Parent) $ProjectName
$CurrentDir = Get-Location

# Check if project directory already exists
if (Test-Path $ProjectDir) {
    Write-Error "Project directory '$ProjectDir' already exists!"
    Write-Host "Please choose a different project name or remove the existing directory."
    exit 1
}

Write-Status "Creating new project: $ProjectName"
Write-Status "Target directory: $ProjectDir"

# Create project directory
New-Item -ItemType Directory -Path $ProjectDir -Force | Out-Null

# Copy the entire monorepo structure
Write-Status "Copying boilerplate files..."
Copy-Item -Path ".\*" -Destination $ProjectDir -Recurse -Force

# Remove unnecessary files from the new project
Set-Location $ProjectDir

# Remove git directory and other unnecessary files
if (Test-Path ".git") { Remove-Item ".git" -Recurse -Force }
if (Test-Path "scripts\create-project.sh") { Remove-Item "scripts\create-project.sh" -Force }
if (Test-Path "scripts\create-project.bat") { Remove-Item "scripts\create-project.bat" -Force }
if (Test-Path "scripts\create-project.ps1") { Remove-Item "scripts\create-project.ps1" -Force }
if (Test-Path "tmp") { Remove-Item "tmp" -Recurse -Force }
if (Test-Path "node_modules") { Remove-Item "node_modules" -Recurse -Force }

# Update Go module name
if (Test-Path "apps\backend\go.mod") {
    Write-Status "Updating Go module name..."
    $goModContent = Get-Content "apps\backend\go.mod" -Raw
    $goModContent = $goModContent -replace 'github\.com/albuquerquewizard/monorepo/backend', "github.com/$env:USERNAME/$ProjectName/backend"
    Set-Content "apps\backend\go.mod" $goModContent -NoNewline
}

# Update mobile app configuration
if (Test-Path "apps\mobile\app.json") {
    Write-Status "Updating mobile app configuration..."
    $appJsonContent = Get-Content "apps\mobile\app.json" -Raw
    
    # Update app name
    $appJsonContent = $appJsonContent -replace '"name": "mobile"', "`"name`": `"$ProjectName`""
    # Update slug
    $appJsonContent = $appJsonContent -replace '"slug": "mobile"', "`"slug`": `"$ProjectName`""
    
    Set-Content "apps\mobile\app.json" $appJsonContent -NoNewline
}

if (Test-Path "apps\mobile\package.json") {
    Write-Status "Updating package.json..."
    $packageJsonContent = Get-Content "apps\mobile\package.json" -Raw
    
    # Update package name
    $packageJsonContent = $packageJsonContent -replace '"name": "mobile"', "`"name`": `"$ProjectName`""
    # Update description
    $packageJsonContent = $packageJsonContent -replace '"description": "Mobile React Native Boilerplate"', "`"description`": `"$ProjectName - Mobile App`""
    
    Set-Content "apps\mobile\package.json" $packageJsonContent -NoNewline
}

# Update README
if (Test-Path "README.md") {
    Write-Status "Updating README..."
    $readmeContent = Get-Content "README.md" -Raw
    $readmeContent = $readmeContent -replace 'Full-Stack Boilerplate Monorepo', "$ProjectName - Full-Stack Project"
    $readmeContent = $readmeContent -replace 'A production-ready monorepo template containing', 'A production-ready project containing'
    Set-Content "README.md" $readmeContent -NoNewline
}

# Create .env template for backend
if (Test-Path "apps\backend\.env.example") {
    Write-Status "Creating .env template..."
    Copy-Item "apps\backend\.env.example" "apps\backend\.env"
    Write-Warning "Please update apps\backend\.env with your database credentials!"
}

# Create .gitignore if it doesn't exist
if (-not (Test-Path ".gitignore")) {
    Write-Status "Creating .gitignore..."
    @"
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
"@ | Set-Content ".gitignore"
}

# Initialize git repository
Write-Status "Initializing git repository..."
try {
    git init | Out-Null
    git add . | Out-Null
    git commit -m "Initial commit: $ProjectName project created from boilerplate" | Out-Null
} catch {
    Write-Warning "Git not found or failed to initialize. Please initialize manually."
}

Write-Success "Project '$ProjectName' created successfully!"
Write-Host ""
Write-Host "Next steps:"
Write-Host "1. cd $ProjectDir"
Write-Host "2. Update database configuration in apps\backend\.env"
Write-Host "3. Customize the codebase for your specific needs"
Write-Host "4. Start development:"
Write-Host "   - Backend: cd apps\backend && task run"
Write-Host "   - Mobile: cd apps\mobile && npm start"
Write-Host ""
Write-Host "Happy coding! 🚀"

# Return to original directory
Set-Location $CurrentDir
