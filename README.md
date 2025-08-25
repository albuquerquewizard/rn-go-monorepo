# 🚀 Full-Stack Boilerplate Monorepo

A production-ready monorepo template containing a **Go backend** with Fiber framework and a **React Native mobile app** built with Expo and NativeWind. Perfect for rapid prototyping and practice projects.

## 🏗️ Project Structure

```
monorepo/
├── apps/
│   ├── backend/         # Go backend with Fiber
│   │   ├── cmd/        # Application entry points
│   │   ├── internal/   # Private application code
│   │   ├── docker/     # Docker configurations
│   │   ├── scripts/    # Backend utility scripts
│   │   ├── .air.toml   # Hot reload configuration
│   │   ├── .golangci.yml # Go linting configuration
│   │   └── Taskfile.yml # Task runner for Go
│   └── mobile/         # React Native app with Expo
│       ├── app/        # Expo Router app directory
│       ├── components/ # Reusable components
│       ├── types/      # TypeScript definitions
│       ├── assets/     # Images, fonts, and static files
│       ├── android/    # Native Android code
│       └── .vscode/    # VS Code configuration
├── packages/            # Shared packages (future)
└── scripts/            # Project creation scripts
```

## 🚀 Quick Start

### Prerequisites
- **Go 1.22+** for backend
- **Node.js 18+** for mobile app
- **Docker** (optional, for containerized development)
- **PostgreSQL** database

### 1. Clone and Setup
```bash
git clone <your-repo-url>
cd monorepo
```

### 2. Backend Setup
```bash
cd apps/backend

# Install dependencies
go mod tidy

# Set environment variables
cp .env.example .env  # Create from template if exists

# Run the application
task run
# or
go run ./cmd/go-boilerplate
```

### 3. Mobile App Setup
```bash
cd apps/mobile

# Install dependencies
pnpm install

# Start development server
pnpm start
```

## 🔧 Backend (Go + Fiber)

### Features
- **Fiber v2** - Fast HTTP framework
- **GORM** - ORM with PostgreSQL support
- **Zerolog** - Structured logging
- **Viper** - Configuration management
- **Validator** - Input validation
- **Task** - Task runner for common operations
- **Air** - Hot reload for development
- **golangci-lint** - Code quality and linting

### Tech Stack
- Go 1.22+
- Fiber v2 (HTTP framework)
- GORM (ORM)
- PostgreSQL (database)
- Zerolog (logging)
- Viper (config)
- Air (hot reload)
- golangci-lint (linting)

### Backend Commands
```bash
cd apps/backend

# Run application
task run

# Build Docker image
task docker:build

# Start with docker-compose
task docker:compose

# Database migrations (using tern)
task migrations:new name=create_users_table
task migrations:up

# Code formatting and linting
task tidy
```

### Environment Variables
Create a `.env` file in `apps/backend/`:
```env
BOILERPLATE_DB_DSN=postgres://user:password@localhost:5432/dbname
BOILERPLATE_PORT=8080
BOILERPLATE_ENV=development
```

## 📱 Mobile App (React Native + Expo)

### Features
- **Expo SDK 53** with React Native 0.79.6
- **NativeWind v4** for Tailwind CSS styling
- **TypeScript** for type safety
- **Expo Router v5** for file-based routing
- **React Hook Form** with Zod validation
- **Zustand** for state management
- **MMKV Storage** for fast key-value storage
- **Expo Vector Icons** and **Lucide React Native** for icons
- **Expo Blur** and **Expo Haptics** for enhanced UX

### Tech Stack
- React Native 0.79.6
- Expo SDK 53
- NativeWind (Tailwind CSS)
- TypeScript
- React Hook Form + Zod
- Zustand
- MMKV Storage
- Expo Vector Icons
- Lucide React Native
- Expo Blur & Haptics
- React Native Reanimated
- React Native Gesture Handler

### Mobile Commands
```bash
cd apps/mobile

# Start development server
pnpm start

# Run on specific platform
pnpm run android
pnpm run ios
pnpm run web

# Lint and format
pnpm run lint
```

## 🆕 Creating New Projects

### Using the Boilerplate Script
```bash
# From the monorepo root
./scripts/create-project.sh my-new-project

# This will create a new directory with your project
cd my-new-project
```

### Manual Setup
1. Copy the entire monorepo structure
2. Update project names in:
   - `apps/backend/go.mod`
   - `apps/mobile/app.json`
   - `apps/mobile/package.json`
3. Update database names and configurations
4. Customize the codebase for your specific needs

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
The backend includes PostgreSQL setup with docker-compose:
```bash
cd apps/backend/docker
docker-compose up -d postgres
```

## 📚 Development Workflow

### Backend Development
1. **Start database**: `task docker:compose`
2. **Run migrations**: `task migrations:up`
3. **Start server**: `task run` (with Air hot reload)
4. **Make changes** and see live reload

### Mobile Development
1. **Start Expo server**: `pnpm start`
2. **Use Expo Go app** on your device
3. **Make changes** and see live reload

## 🛠️ Available Scripts

### Backend (Task)
- `task run` - Run the application with hot reload
- `task docker:build` - Build production Docker image
- `task docker:dev` - Build development Docker image
- `task docker:compose` - Start services with docker-compose
- `task migrations:new name=X` - Create new migration using tern
- `task migrations:up` - Apply migrations using tern
- `task tidy` - Format code and tidy dependencies

### Mobile (pnpm)
- `pnpm start` - Start Expo development server
- `pnpm run android` - Run on Android
- `pnpm run ios` - Run on iOS
- `pnpm run web` - Run in web browser
- `pnpm run lint` - Run ESLint

## 🔄 Project Customization

### Backend Customization
1. **Update module name** in `go.mod`
2. **Modify models** in `internal/models/`
3. **Add new routes** in `internal/routes/`
4. **Create new services** in `internal/services/`
5. **Update database schema** with migrations
6. **Configure linting** in `.golangci.yml`
7. **Adjust hot reload** in `.air.toml`

### Mobile Customization
1. **Update app name** in `app.json`
2. **Modify screens** in `app/` directory
3. **Add new components** in `components/`
4. **Update types** in `types/`
5. **Customize styling** in `tailwind.config.js`
6. **Add assets** in `assets/` directory

## 📖 Documentation

- [Go Documentation](https://golang.org/doc/)
- [Fiber Documentation](https://docs.gofiber.io/)
- [GORM Documentation](https://gorm.io/docs/)
- [Expo Documentation](https://docs.expo.dev/)
- [React Native Documentation](https://reactnative.dev/)
- [NativeWind Documentation](https://www.nativewind.dev/)
- [Air Documentation](https://github.com/cosmtrek/air)
- [golangci-lint Documentation](https://golangci-lint.run/)
- [Tern Migration Tool](https://github.com/jackc/tern)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

If you encounter any issues:
1. Check the documentation links above
2. Search existing GitHub issues
3. Create a new issue with detailed information

---

**Built with ❤️ using Go, Fiber, React Native, and Expo**
