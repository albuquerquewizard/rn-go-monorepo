@echo off
title Full-Stack Boilerplate Project Creator

echo.
echo ========================================
echo    Full-Stack Boilerplate Creator
echo ========================================
echo.

if "%~1"=="" (
    echo Usage: create.bat ^<project-name^>
    echo.
    echo Example: create.bat my-awesome-project
    echo.
    echo Available scripts:
    echo   - create-project.bat    (Command Prompt)
    echo   - create-project.ps1    (PowerShell - Recommended)
    echo   - create-project.sh     (Bash - WSL/Git Bash)
    echo.
    pause
    exit /b 1
)

echo Creating project: %~1
echo.
echo Choose your preferred script:
echo.
echo 1. PowerShell (Recommended for Windows)
echo 2. Command Prompt
echo 3. Bash (WSL/Git Bash)
echo.
set /p choice="Enter your choice (1-3): "

if "%choice%"=="1" (
    echo.
    echo Running PowerShell script...
    powershell -ExecutionPolicy Bypass -File "%~dp0create-project.ps1" "%~1"
) else if "%choice%"=="2" (
    echo.
    echo Running Command Prompt script...
    call "%~dp0create-project.bat" "%~1"
) else if "%choice%"=="3" (
    echo.
    echo Running Bash script...
    bash "%~dp0create-project.sh" "%~1"
) else (
    echo.
    echo Invalid choice. Please run the script again.
    pause
    exit /b 1
)

echo.
echo Script execution completed.
pause
