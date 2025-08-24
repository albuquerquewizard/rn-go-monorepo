# 🚀 Mobile React Native Boilerplate + Backend (Coming Soon)

A modern, production-ready monorepo containing a React Native mobile app boilerplate built with Expo and NativeWind, with a Go backend in development.

## 📱 Mobile App (React Native + Expo)

### Features
- **Expo SDK 53** with React Native 0.79.6
- **NativeWind v4** for Tailwind CSS styling
- **TypeScript** for type safety
- **Expo Router v5** for file-based routing
- **React Hook Form** with Zod validation
- **Zustand** for state management
- **AsyncStorage** for local data persistence
- **MMKV** for high-performance key-value storage
- **React Native Reanimated** for smooth animations
- **Lucide React Native** for beautiful icons
- **Axios** for HTTP requests
- **ESLint + Prettier** for code quality

### Tech Stack
- React Native 0.79.6
- Expo SDK 53
- NativeWind (Tailwind CSS)
- TypeScript
- React Hook Form + Zod
- Zustand
- MMKV Storage

### Getting Started

1. **Install dependencies**
   ```bash
   cd apps/mobile
   npm install
   ```

2. **Start the development server**
   ```bash
   npm start
   # or
   npx expo start
   ```

3. **Run on device/emulator**
   ```bash
   # Android
   npm run android
   
   # iOS
   npm run ios
   
   # Web
   npm run web
   ```

### Project Structure
```
apps/mobile/
├── app/                 # Expo Router app directory
├── components/          # Reusable components
├── types/              # TypeScript type definitions
├── assets/             # Images, fonts, etc.
├── global.css          # Global Tailwind styles
└── package.json        # Dependencies and scripts
```

## 🔧 Backend (Go - In Development)
- Coming soon with Go backend
- RESTful API design
- Database integration
- Authentication system

## 🏗️ Monorepo Structure
```
monorepo/
├── apps/
│   ├── mobile/         # React Native app
│   └── backend/        # Go backend (coming soon)
├── packages/            # Shared packages (future)
└── README.md           # This file
```

## 🚀 Quick Start

1. **Clone the repository**
   ```bash
   git clone <your-repo-url>
   cd monorepo
   ```

2. **Set up mobile app**
   ```bash
   cd apps/mobile
   npm install
   npm start
   ```

3. **Follow Expo instructions** to run on your preferred platform

## 📱 Mobile App Screenshots

The mobile app includes:
- Modern UI with NativeWind styling
- Responsive design
- Type-safe development
- File-based routing with Expo Router

## 🛠️ Development

### Prerequisites
- Node.js 18+ 
- npm or yarn
- Expo CLI
- Android Studio (for Android development)
- Xcode (for iOS development, macOS only)

### Available Scripts
- `npm start` - Start Expo development server
- `npm run android` - Run on Android device/emulator
- `npm run ios` - Run on iOS device/simulator
- `npm run web` - Run in web browser
- `npm run lint` - Run ESLint
- `npm run reset-project` - Reset to blank project

## 📚 Documentation

- [Expo Documentation](https://docs.expo.dev/)
- [React Native Documentation](https://reactnative.dev/)
- [NativeWind Documentation](https://www.nativewind.dev/)
- [Expo Router Documentation](https://expo.github.io/router/)

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

If you encounter any issues or have questions:
1. Check the [Expo documentation](https://docs.expo.dev/)
2. Search existing [GitHub issues](https://github.com/yourusername/yourrepo/issues)
3. Create a new issue with detailed information

---

**Built with ❤️ using Expo, React Native, and NativeWind**
