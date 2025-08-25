# 🚀 Project Setup Guide

This guide will help you set up a new project from the Full-Stack Boilerplate template.

## 📋 Prerequisites

Before creating a new project, ensure you have the following installed:

### Required Software
- **Go 1.25+** - [Download here](https://golang.org/dl/)
- **Node.js 18+** - [Download here](https://nodejs.org/)
- **Git** - [Download here](https://git-scm.com/)
- **Docker** (optional) - [Download here](https://www.docker.com/)

### Optional Software
- **PostgreSQL** - For local database development
- **Android Studio** - For Android mobile development
- **Xcode** - For iOS development (macOS only)

## 🆕 Creating a New Project

### Option 1: Using the Automated Scripts (Recommended)

#### For Unix/Linux/macOS:
```bash
# Make the script executable (first time only)
chmod +x scripts/create-project.sh

# Create a new project
./scripts/create-project.sh my-awesome-project
```

#### For Windows (Command Prompt):
```cmd
scripts\create-project.bat my-awesome-project
```

#### For Windows (PowerShell):
```powershell
.\scripts\create-project.ps1 my-awesome-project
```

### Option 2: Manual Setup

1. **Copy the boilerplate structure**
   ```bash
   cp -r /path/to/boilerplate /path/to/new-project
   cd /path/to/new-project
   ```

2. **Update project names manually**
   - Edit `apps/backend/go.mod`
   - Edit `apps/mobile/app.json`
   - Edit `apps/mobile/package.json`
   - Edit `README.md`

3. **Remove boilerplate-specific files**
   ```bash
   rm -rf .git
   rm -rf scripts/
   rm -rf tmp/
   rm -rf node_modules/
   ```

## ⚙️ Project Configuration

### Backend Configuration

1. **Update Go module name**
   ```go
   // In apps/backend/go.mod
   module github.com/yourusername/your-project-name/backend
   ```

2. **Set environment variables**
   ```bash
   cd apps/backend
   cp .env.example .env
   ```
   
   Edit `.env` with your database credentials:
   ```env
   BOILERPLATE_DB_DSN=postgres://user:password@localhost:5432/dbname
   BOILERPLATE_PORT=8080
   BOILERPLATE_ENV=development
   ```

3. **Database setup**
   ```bash
   # Start PostgreSQL (if using Docker)
   cd docker
   docker-compose up -d postgres
   
   # Run migrations
   cd ../..
   task migrations:up
   ```

### Mobile App Configuration

1. **Update app metadata**
   ```json
   // In apps/mobile/app.json
   {
     "name": "Your App Name",
     "slug": "your-app-slug"
   }
   ```

2. **Update package.json**
   ```json
   {
     "name": "your-app-name",
     "description": "Your App Description"
   }
   ```

## 🚀 First Run

### Backend
```bash
cd apps/backend

# Install Go dependencies
go mod tidy

# Run the application
task run
# or
go run ./cmd/go-boilerplate
```

### Mobile App
```bash
cd apps/mobile

# Install Node.js dependencies
npm install

# Start Expo development server
npm start
```

## 🔧 Customization Guide

### Backend Customization

#### 1. Models
- Edit `internal/models/` to define your data structures
- Update `internal/models/models.go` to register new models

#### 2. Routes
- Add new endpoints in `internal/routes/`
- Update `internal/routes/routes.go` to register new routes

#### 3. Services
- Implement business logic in `internal/services/`
- Follow the existing service pattern

#### 4. Controllers
- Handle HTTP requests in `internal/controllers/`
- Use the existing controller structure

### Mobile App Customization

#### 1. Screens
- Add new screens in the `app/` directory
- Follow Expo Router v5 file-based routing

#### 2. Components
- Create reusable components in `components/`
- Use NativeWind for styling

#### 3. State Management
- Use Zustand stores for global state
- Follow the existing store pattern

#### 4. API Integration
- Update API endpoints in your services
- Ensure they match your backend routes

## 🐳 Docker Development

### Backend with Docker
```bash
cd apps/backend

# Build and run with docker-compose
task docker:compose

# Build individual images
task docker:build    # Production
task docker:dev      # Development
```

### Database
```bash
cd apps/backend/docker
docker-compose up -d postgres
```

## 📱 Mobile Development

### Expo Development
1. Install Expo Go app on your device
2. Scan the QR code from `npm start`
3. Make changes and see live reload

### Platform-Specific Development
```bash
# Android
npm run android

# iOS
npm run ios

# Web
npm run web
```

## 🔍 Troubleshooting

### Common Issues

#### Backend Issues
- **Port already in use**: Change `BOILERPLATE_PORT` in `.env`
- **Database connection failed**: Check PostgreSQL is running and credentials are correct
- **Go module errors**: Run `go mod tidy` and `go mod download`

#### Mobile Issues
- **Metro bundler errors**: Clear cache with `npx expo start --clear`
- **Build failures**: Check Node.js version and clear `node_modules`
- **Device not detected**: Ensure USB debugging is enabled (Android) or device is trusted (iOS)

### Getting Help
1. Check the [documentation links](#documentation) in the main README
2. Search existing GitHub issues
3. Create a new issue with detailed error information

## 📚 Next Steps

After setting up your project:

1. **Customize the codebase** for your specific needs
2. **Add your business logic** to the backend services
3. **Design your mobile app UI** using NativeWind
4. **Set up CI/CD** for automated testing and deployment
5. **Add tests** for both backend and mobile code
6. **Configure production environments**

## 🎯 Best Practices

### Backend
- Follow Go naming conventions
- Use structured logging with Zerolog
- Implement proper error handling
- Add input validation for all endpoints
- Use database migrations for schema changes

### Mobile
- Follow React Native best practices
- Use TypeScript for type safety
- Implement proper error boundaries
- Optimize bundle size
- Test on multiple devices

### General
- Keep dependencies updated
- Use semantic versioning
- Write clear commit messages
- Document your API endpoints
- Set up automated testing

---

**Happy coding! 🚀**

For more information, refer to the main [README.md](../README.md) file.
