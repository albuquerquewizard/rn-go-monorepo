# 📖 Script Usage Guide

This guide shows you how to use the project creation scripts to quickly set up new projects from this boilerplate.

## 🚀 Quick Start

### For Windows Users

#### Using PowerShell (Recommended)
```powershell
# Navigate to your boilerplate directory
cd C:\path\to\your\boilerplate

# Run the PowerShell script
.\scripts\create-project.ps1 my-new-project
```

#### Using Command Prompt
```cmd
# Navigate to your boilerplate directory
cd C:\path\to\your\boilerplate

# Run the batch script
scripts\create-project.bat my-new-project
```

### For Unix/Linux/macOS Users

```bash
# Navigate to your boilerplate directory
cd /path/to/your/boilerplate

# Make the script executable (first time only)
chmod +x scripts/create-project.sh

# Run the script
./scripts/create-project.sh my-new-project
```

## 📝 What the Scripts Do

The project creation scripts will:

1. ✅ **Create a new directory** with your project name
2. ✅ **Copy all boilerplate files** to the new project
3. ✅ **Update project names** in configuration files
4. ✅ **Remove boilerplate-specific files** (scripts, .git, etc.)
5. ✅ **Initialize a new git repository**
6. ✅ **Create a .gitignore file**
7. ✅ **Set up environment templates**

## 🔧 After Running the Script

Once your new project is created:

1. **Navigate to your new project**
   ```bash
   cd ../my-new-project
   ```

2. **Configure the backend**
   ```bash
   cd apps/backend
   # Edit .env file with your database credentials
   go mod tidy
   ```

3. **Configure the mobile app**
   ```bash
   cd ../mobile
   npm install
   ```

4. **Start development**
   ```bash
   # Backend
   cd ../backend
   task run
   
   # Mobile (in another terminal)
   cd ../mobile
   npm start
   ```

## 🎯 Example Usage

```bash
# Create a todo app
.\scripts\create-project.ps1 todo-app

# Create a chat application
.\scripts\create-project.ps1 chat-app

# Create an e-commerce platform
.\scripts\create-project.ps1 ecommerce-platform
```

## ⚠️ Important Notes

- **Project names** should only contain letters, numbers, hyphens, and underscores
- **Don't run the script** from inside the new project directory
- **Always run from** the boilerplate root directory
- **Backup your work** before running the script (just in case)

## 🆘 Troubleshooting

### Script won't run
- Ensure you're in the boilerplate root directory
- Check that the script files exist in the `scripts/` folder
- On Windows, try running PowerShell as Administrator

### Permission denied (Unix/Linux/macOS)
```bash
chmod +x scripts/create-project.sh
```

### Project already exists
- Choose a different project name, or
- Remove the existing directory first

## 📚 Next Steps

After creating your project, refer to:
- [PROJECT_SETUP.md](PROJECT_SETUP.md) - Detailed setup guide
- [../README.md](../README.md) - Main boilerplate documentation

---

**Happy project creation! 🚀**
