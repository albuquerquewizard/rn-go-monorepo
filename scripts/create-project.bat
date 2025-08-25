@echo off
setlocal enabledelayedexpansion

REM Full-Stack Boilerplate Project Creator for Windows
REM This script creates a new project from the boilerplate template

REM Check if project name is provided
if "%~1"=="" (
    echo [ERROR] Project name is required!
    echo Usage: %0 ^<project-name^>
    echo Example: %0 my-awesome-project
    exit /b 1
)

set PROJECT_NAME=%~1
set PROJECT_DIR=..\%PROJECT_NAME%
set CURRENT_DIR=%CD%

REM Validate project name (basic check for Windows)
echo %PROJECT_NAME%| findstr /r "^[a-zA-Z0-9_-]*$" >nul
if errorlevel 1 (
    echo [ERROR] Invalid project name! Use only letters, numbers, hyphens, and underscores.
    exit /b 1
)

REM Check if project directory already exists
if exist "%PROJECT_DIR%" (
    echo [ERROR] Project directory '%PROJECT_DIR%' already exists!
    echo Please choose a different project name or remove the existing directory.
    exit /b 1
)

echo [INFO] Creating new project: %PROJECT_NAME%
echo [INFO] Target directory: %PROJECT_DIR%

REM Create project directory
mkdir "%PROJECT_DIR%"

REM Copy the entire monorepo structure
echo [INFO] Copying boilerplate files...
xcopy . "%PROJECT_DIR%" /E /I /H /Y >nul

REM Remove unnecessary files from the new project
cd /d "%PROJECT_DIR%"
if exist .git rmdir /s /q .git
if exist scripts\create-project.sh del /q scripts\create-project.sh
if exist scripts\create-project.bat del /q scripts\create-project.bat
if exist tmp rmdir /s /q tmp
if exist node_modules rmdir /s /q node_modules

REM Update Go module name
if exist "apps\backend\go.mod" (
    echo [INFO] Updating Go module name...
    powershell -Command "(Get-Content 'apps\backend\go.mod') -replace 'github\.com/albuquerquewizard/monorepo/backend', 'github.com/%USERNAME%/%PROJECT_NAME%/backend' | Set-Content 'apps\backend\go.mod'"
)

REM Update mobile app configuration
if exist "apps\mobile\app.json" (
    echo [INFO] Updating mobile app configuration...
    REM Update app name
    powershell -Command "(Get-Content 'apps\mobile\app.json') -replace '\"name\": \"mobile\"', '\"name\": \"%PROJECT_NAME%\"' | Set-Content 'apps\mobile\app.json'"
    REM Update slug
    powershell -Command "(Get-Content 'apps\mobile\app.json') -replace '\"slug\": \"mobile\"', '\"slug\": \"%PROJECT_NAME%\"' | Set-Content 'apps\mobile\app.json'"
)

if exist "apps\mobile\package.json" (
    echo [INFO] Updating package.json...
    REM Update package name
    powershell -Command "(Get-Content 'apps\mobile\package.json') -replace '\"name\": \"mobile\"', '\"name\": \"%PROJECT_NAME%\"' | Set-Content 'apps\mobile\package.json'"
    REM Update description
    powershell -Command "(Get-Content 'apps\mobile\package.json') -replace '\"description\": \"Mobile React Native Boilerplate\"', '\"description\": \"%PROJECT_NAME% - Mobile App\"' | Set-Content 'apps\mobile\package.json'"
)

REM Update README
if exist "README.md" (
    echo [INFO] Updating README...
    powershell -Command "(Get-Content 'README.md') -replace 'Full-Stack Boilerplate Monorepo', '%PROJECT_NAME% - Full-Stack Project' | Set-Content 'README.md'"
    powershell -Command "(Get-Content 'README.md') -replace 'A production-ready monorepo template containing', 'A production-ready project containing' | Set-Content 'README.md'"
)

REM Create .env template for backend
if exist "apps\backend\.env.example" (
    echo [INFO] Creating .env template...
    copy "apps\backend\.env.example" "apps\backend\.env" >nul
    echo [WARNING] Please update apps\backend\.env with your database credentials!
)

REM Create .gitignore if it doesn't exist
if not exist ".gitignore" (
    echo [INFO] Creating .gitignore...
    (
        echo # Dependencies
        echo node_modules/
        echo vendor/
        echo.
        echo # Environment files
        echo .env
        echo .env.local
        echo .env.*.local
        echo.
        echo # Build outputs
        echo dist/
        echo build/
        echo *.exe
        echo *.dll
        echo *.so
        echo *.dylib
        echo.
        echo # IDE files
        echo .vscode/
        echo .idea/
        echo *.swp
        echo *.swo
        echo *~
        echo.
        echo # OS files
        echo .DS_Store
        echo Thumbs.db
        echo.
        echo # Logs
        echo *.log
        echo logs/
        echo.
        echo # Temporary files
        echo tmp/
        echo temp/
        echo.
        echo # Database
        echo *.db
        echo *.sqlite
        echo.
        echo # Go
        echo *.test
        echo *.out
        echo.
        echo # Mobile
        echo .expo/
        echo *.jks
        echo *.p8
        echo *.p12
        echo *.key
        echo *.mobileprovision
        echo *.orig.*
        echo web-build/
        echo.
        echo # Docker
        echo .dockerignore
    ) > .gitignore
)

REM Initialize git repository
echo [INFO] Initializing git repository...
git init >nul 2>&1
if errorlevel 1 (
    echo [WARNING] Git not found or failed to initialize. Please initialize manually.
) else (
    git add . >nul 2>&1
    git commit -m "Initial commit: %PROJECT_NAME% project created from boilerplate" >nul 2>&1
)

echo [SUCCESS] Project '%PROJECT_NAME%' created successfully!
echo.
echo Next steps:
echo 1. cd %PROJECT_DIR%
echo 2. Update database configuration in apps\backend\.env
echo 3. Customize the codebase for your specific needs
echo 4. Start development:
echo    - Backend: cd apps\backend ^&^& task run
echo    - Mobile: cd apps\mobile ^&^& npm start
echo.
echo Happy coding! 🚀

cd /d "%CURRENT_DIR%"
