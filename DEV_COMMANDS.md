# 🚀 **MONOREPO DEVELOPMENT COMMANDS**

*Quick reference for all development commands in your Go + React Native monorepo*

---

## 📍 **ROOT LEVEL COMMANDS** (Run from monorepo root)

### **🚀 Start Development Environment**
```bash
# Start both backend and mobile simultaneously
pnpm turbo run start

# Start in development mode
pnpm turbo run dev

# Start with force (ignore cache)
pnpm turbo run start --force
```

### **📱 Mobile App Commands**
```bash
# Start Expo development server
pnpm turbo run start

# Run on specific platforms
pnpm turbo run android    # Android device/emulator
pnpm turbo run ios        # iOS simulator
pnpm turbo run web        # Web browser

# Lint mobile code
pnpm turbo run lint
```

### **🔧 Backend Commands**
```bash
# Start Go backend with hot reload
pnpm turbo run start

# Run Go tests
pnpm turbo run test

# Lint Go code
pnpm turbo run lint

# Format and tidy Go code
pnpm turbo run tidy
```

### **🐳 Docker Commands**
```bash
# Start PostgreSQL database
pnpm turbo run docker:compose

# Stop database
pnpm turbo run docker:compose:down

# Build production Docker image
pnpm turbo run docker:build

# Build development Docker image
pnpm turbo run docker:dev
```

### **🗄️ Database Commands**
```bash
# Create new migration
pnpm turbo run migrations:new name=migration_name

# Run all migrations
pnpm turbo run migrations:up
```

### **🧹 Utility Commands**
```bash
# Reset project (both packages)
pnpm turbo run reset-project

# Install all dependencies
pnpm install

# Check available commands
pnpm turbo run --dry-run
```

---

## 📁 **BACKEND COMMANDS** (Run from `apps/backend/`)

### **🚀 Development**
```bash
# Start with hot reload (Air)
task run

# Start without hot reload
go run ./cmd/go-boilerplate

# Build binary
go build ./cmd/go-boilerplate

# Run tests
go test ./...
```

### **🔧 Task Commands**
```bash
# List all available tasks
task --list-all

# Format Go code
task tidy

# Docker operations
task docker:build
task docker:dev
task docker:compose
task docker:compose:down

# Database migrations
task migrations:new name=migration_name
task migrations:up
```

### **📦 Go Module Management**
```bash
# Install dependencies
go mod tidy

# Verify dependencies
go mod verify

# Download dependencies
go mod download
```

### **🔍 Code Quality**
```bash
# Run linter
golangci-lint run

# Format code
go fmt ./...

# Vet code
go vet ./...
```

---

## 📱 **MOBILE COMMANDS** (Run from `apps/mobile/`)

### **🚀 Development**
```bash
# Start Expo development server
npm start
# or
expo start

# Start with specific platform
npm run android
npm run ios
npm run web
```

### **📱 Platform Specific**
```bash
# Android
expo run:android
expo start --android

# iOS
expo run:ios
expo start --ios

# Web
expo start --web
expo export:web
```

### **🔧 Development Tools**
```bash
# Lint code
npm run lint
expo lint

# Reset project
npm run reset-project

# Install dependencies
npm install
```

---

## 🌐 **ENVIRONMENT SETUP**

### **Backend Environment**
```bash
cd apps/backend

# Copy environment template
cp .env.example .env

# Edit with your database credentials
# .env should contain:
APP_PORT=8080
APP_ENV=development
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=go_boilerplate
```

### **Database Setup**
```bash
# Start PostgreSQL
cd apps/backend
task docker:compose

# Run migrations
task migrations:up
```

---

## 🚀 **QUICK START WORKFLOW**

### **1. First Time Setup**
```bash
# Install dependencies
pnpm install

# Setup backend environment
cd apps/backend
cp .env.example .env
# Edit .env with your database credentials

# Start database
task docker:compose

# Run migrations
task migrations:up

# Return to root
cd ../..
```

### **2. Daily Development**
```bash
# Start everything at once
pnpm turbo run start

# This starts:
# 🚀 Backend on http://localhost:8080
# 📱 Expo on http://localhost:19000
# 🗄️ PostgreSQL on localhost:5432
```

### **3. Development Workflow**
```bash
# Terminal 1: Backend development
cd apps/backend
task run

# Terminal 2: Mobile development  
cd apps/mobile
npm start

# Or use single command from root:
pnpm turbo run start
```

---

## 🔧 **TROUBLESHOOTING**

### **Common Issues**
```bash
# Clear Turborepo cache
pnpm turbo run --force

# Reinstall dependencies
pnpm install

# Reset Go modules
cd apps/backend
go mod tidy
go mod download

# Clear Expo cache
cd apps/mobile
npx expo start --clear
```

### **Port Conflicts**
```bash
# Backend port conflict
# Edit apps/backend/.env: APP_PORT=8081

# Expo port conflict
# Edit apps/mobile/package.json: "start": "expo start --port 19001"
```

---

## 📚 **USEFUL ALIASES** (Add to your shell profile)

```bash
# Quick navigation
alias mb="cd apps/backend"
alias mm="cd apps/mobile"
alias mr="cd /path/to/monorepo"

# Quick commands
alias dev="pnpm turbo run start"
alias devb="cd apps/backend && task run"
alias devm="cd apps/mobile && npm start"
alias db="cd apps/backend && task docker:compose"
```

---

## 🎯 **PRO TIPS**

1. **Use `pnpm turbo run start`** from root to start everything
2. **Keep database running** with `task docker:compose` in backend
3. **Use `--force` flag** when you need to bypass Turborepo cache
4. **Run `pnpm turbo run --dry-run`** to see what will execute
5. **Use `task --list-all`** in backend to see all available tasks

---

*💡 **Remember**: Most commands can be run from the root using `pnpm turbo run <command>` for convenience!*
