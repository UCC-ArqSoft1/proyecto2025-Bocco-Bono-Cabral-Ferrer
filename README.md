# 🏃‍♂️ GymFlow - Gym Management System

[![Go Version](https://img.shields.io/badge/Go-1.24.1-blue.svg)](https://golang.org/)
[![React Version](https://img.shields.io/badge/React-19.1.0-61dafb.svg)](https://reactjs.org/)
[![MySQL](https://img.shields.io/badge/MySQL-9.3-orange.svg)](https://mysql.com/)
[![Docker](https://img.shields.io/badge/Docker-Ready-blue.svg)](https://docker.com/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A comprehensive gym management system built with modern web technologies, featuring activity scheduling, user enrollment management, and admin controls. This full-stack application demonstrates enterprise-level software engineering practices with clean architecture, JWT authentication, and containerized deployment.

## 📑 Table of Contents

- [Key Features](#-key-features)
- [Tech Stack](#-tech-stack)
- [Architecture Overview](#-architecture-overview)
- [Installation & Setup](#-installation--setup)
- [Usage Examples](#-usage-examples)
- [API Documentation](#-api-documentation)
- [Project Structure](#-project-structure)
- [Development](#-development)
- [Deployment](#-deployment)
- [Performance & Scalability](#-performance--scalability)
- [Challenges & Solutions](#-challenges--solutions)
- [Contributing](#-contributing)

## ✨ Key Features

### 🔐 **Authentication & Authorization**
- JWT-based authentication with automatic token expiration handling
- Role-based access control (Admin/User roles)
- Secure password handling and session management
- Protected routes with middleware validation

### 🏋️ **Activity Management**
- Complete CRUD operations for gym activities
- Advanced filtering and search capabilities
- Category-based organization (Cardio, Strength, Wellness, etc.)
- Instructor assignment and scheduling system
- Capacity management with real-time availability

### 📅 **Enrollment System**
- User enrollment in activities with capacity validation
- Real-time enrollment status tracking
- Automated enrollment date management
- Admin oversight of all enrollments
- Cancellation functionality with instant updates

### 👨‍💼 **Admin Dashboard**
- Comprehensive activity management interface
- User enrollment monitoring and management
- System analytics and reporting capabilities
- Bulk operations for efficient administration

### 🎨 **Modern UI/UX**
- Responsive design with mobile-first approach
- Real-time notifications using React Toastify
- Intuitive navigation with React Router
- Clean, professional interface design
- Loading states and error handling

## 🛠 Tech Stack

### **Backend**
- **Go 1.24.1** - High-performance server-side language
- **Gin Web Framework** - Fast HTTP web framework
- **GORM** - Feature-rich ORM for database operations
- **JWT-Go** - JSON Web Token implementation
- **MySQL Driver** - Database connectivity
- **Testify** - Testing framework with assertions
- **CORS Middleware** - Cross-origin resource sharing

### **Frontend**
- **React 19.1.0** - Modern component-based UI library
- **Vite 6.3.5** - Lightning-fast build tool and dev server
- **React Router Dom 7.6.1** - Client-side routing
- **Axios 1.9.0** - Promise-based HTTP client
- **React Toastify 11.0.5** - Elegant notification system
- **JWT Decode 4.0.0** - JWT token parsing
- **React Icons 5.5.0** - Icon library

### **Database & Infrastructure**
- **MySQL 9.3** - Reliable relational database
- **Docker** - Containerization platform
- **Docker Compose** - Multi-container orchestration

### **Development Tools**
- **ESLint** - JavaScript/TypeScript linting
- **TypeScript Support** - Type definitions included
- **Hot Module Replacement** - Fast development experience

## 🏗 Architecture Overview

The application follows a **clean architecture pattern** with clear separation of concerns:

### Backend Architecture
```
├── Controllers/     # HTTP request handlers
├── Services/        # Business logic layer
├── Clients/         # Data access layer
├── DAO/            # Database entities
├── Domain/         # Business domain models
├── Middleware/     # Authentication & validation
└── Database/       # Connection & migration
```

### Design Patterns Implemented
- **Repository Pattern** - Data access abstraction
- **Dependency Injection** - Loose coupling between layers
- **Service Layer Pattern** - Business logic encapsulation
- **DTO Pattern** - Data transfer object mapping
- **Factory Pattern** - Database connection management

### Security Features
- **JWT Authentication** with secure token validation
- **CORS Configuration** for cross-origin requests
- **SQL Injection Prevention** through ORM
- **Password Security** with proper hashing
- **Route Protection** with middleware guards

## 🚀 Installation & Setup

### Prerequisites
- Docker & Docker Compose installed
- Git for version control
- (Optional) Go 1.24.1+ and Node.js 22+ for local development

### Quick Start with Docker

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd proyecto2025-Bocco-Bono-Cabral-Ferrer
   ```

2. **Start the application stack**
   ```bash
   docker-compose up -d
   ```

3. **Initialize database with user types**
   ```sql
   # Connect to MySQL (port 3307) and run:
   INSERT INTO db_gym.user_types (id, name) VALUES (1, 'admin'), (2, 'usuario');
   ```

4. **Access the application**
   - Frontend: http://localhost:5173
   - Backend API: http://localhost:8080
   - MySQL: localhost:3307

### Local Development Setup

#### Backend Setup
```bash
cd backend
go mod tidy
go run main.go
```

#### Frontend Setup
```bash
cd frontend_gym
npm install
npm run dev
```

### Environment Configuration

Create a `.env` file in the backend directory:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=root
DB_NAME=db_gym
JWT_SECRET=your-secret-key
```

## 💡 Usage Examples

### User Registration & Login
```javascript
// Registration
const userData = {
  first_name: "John",
  last_name: "Doe",
  email: "john@example.com",
  password: "securepass123",
  birth_date: "1990-01-01",
  sex: "M"
};

const response = await axios.post('/users/register', userData);
```

### Activity Management
```javascript
// Create Activity (Admin only)
const activityData = {
  name: "Morning Yoga",
  description: "Relaxing yoga session for beginners",
  capacity: 20,
  category: "Wellness",
  profesor: "Jane Smith",
  image_url: "https://example.com/yoga.jpg",
  schedules: [
    {
      day: "Monday",
      start_time: "08:00",
      end_time: "09:00"
    }
  ]
};

const response = await axios.post('/activities', activityData, {
  headers: { Authorization: `Bearer ${token}` }
});
```

### Enrollment Process
```javascript
// Enroll in Activity
const enrollmentData = {
  activity_id: 1
};

const response = await axios.post('/enrollment', enrollmentData, {
  headers: { Authorization: `Bearer ${token}` }
});
```

## 📋 API Documentation

### Authentication Endpoints
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/users/register` | Register new user | No |
| POST | `/users/login` | User authentication | No |

### Activity Endpoints
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| GET | `/activities` | Get all activities | No |
| GET | `/activities/:id` | Get specific activity | No |
| GET | `/activities/search?keyword=` | Search activities | No |
| POST | `/activities` | Create activity | Admin |
| PUT | `/activities/:id` | Update activity | Admin |
| DELETE | `/activities/:id` | Delete activity | Admin |

### Enrollment Endpoints
| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/enrollment` | Enroll in activity | User |
| DELETE | `/enrollment` | Cancel enrollment | User |
| GET | `/enrollment/check` | Check enrollment status | User |
| GET | `/enrollment/capacity` | Get activity capacity | User |
| GET | `/enrollments` | Get user enrollments | User |

### Request/Response Examples

#### Activity Response
```json
{
  "id": 1,
  "name": "Morning Yoga",
  "description": "Relaxing yoga session",
  "capacity": 20,
  "category": "Wellness",
  "profesor": "Jane Smith",
  "image_url": "https://example.com/yoga.jpg",
  "schedules": [
    {
      "id": 1,
      "activity_id": 1,
      "day": "Monday",
      "start_time": "08:00",
      "end_time": "09:00"
    }
  ]
}
```

## 📁 Project Structure

```
proyecto2025-Bocco-Bono-Cabral-Ferrer/
├── 📁 backend/                    # Go backend application
│   ├── 📁 clients/               # Data access layer
│   │   ├── activityclient/       # Activity repository
│   │   ├── enrollmentClient/     # Enrollment repository
│   │   └── userClients/          # User repository
│   ├── 📁 controllers/           # HTTP handlers
│   │   ├── activity/             # Activity endpoints
│   │   ├── enrollment/           # Enrollment endpoints
│   │   └── user/                 # User endpoints
│   ├── 📁 dao/                   # Database entities
│   ├── 📁 domain/                # Business models
│   ├── 📁 services/              # Business logic
│   ├── 📁 middleware/            # Auth & validation
│   ├── 📁 database/              # DB connection
│   ├── 📁 mocks/                 # Test mocks
│   ├── 🐳 Dockerfile             # Backend container
│   ├── 📄 go.mod                 # Go dependencies
│   ├── ⚙️ .env                  # Environment config
│   └── 🧪 *_test.go             # Unit tests
├── 📁 frontend_gym/              # React frontend application
│   ├── 📁 src/
│   │   ├── 📁 components/        # Reusable UI components
│   │   │   ├── header.jsx        # Navigation header
│   │   │   ├── footer.jsx        # Page footer
│   │   │   ├── layout.jsx        # App layout wrapper
│   │   │   ├── MyEnrollments.jsx # User enrollments
│   │   │   └── SearchActivities.jsx # Activity search
│   │   ├── 📁 pages/            # Route components
│   │   │   ├── Activities.jsx    # Activity listing
│   │   │   ├── activitiesAdmin.jsx # Admin dashboard
│   │   │   ├── ActivityDetail.jsx # Activity details
│   │   │   ├── home.jsx          # Landing page
│   │   │   ├── login.jsx         # Authentication
│   │   │   ├── register.jsx      # User registration
│   │   │   └── main.jsx          # App entry point
│   │   ├── 📁 hooks/            # Custom React hooks
│   │   │   ├── useAuth.jsx       # Authentication logic
│   │   │   ├── useAuthValidation.jsx # Auth validation
│   │   │   └── useEnrollment.jsx # Enrollment management
│   │   ├── 📁 Styles/           # CSS stylesheets
│   │   └── 📁 assets/           # Static assets
│   ├── 🐳 Dockerfile            # Frontend container
│   ├── 📄 package.json          # Node dependencies
│   ├── ⚙️ vite.config.js       # Vite configuration
│   └── 🧹 eslint.config.js     # Linting rules
├── 🐳 docker-compose.yml        # Multi-container setup
├── 📁 mysql_data/               # Database persistence
├── 📄 package.json              # Root dependencies
└── 📚 README.md                 # This documentation
```

## 🔧 Development

### Running Tests

#### Backend Tests
```bash
cd backend
go test ./... -v
```

The backend includes comprehensive unit tests for:
- Activity controller operations (CRUD)
- Mock service implementations
- HTTP endpoint validation
- Request/response handling

### Code Quality

#### Frontend Linting
```bash
cd frontend_gym
npm run lint
```

#### Build Process
```bash
# Frontend production build
cd frontend_gym
npm run build

# Backend binary compilation
cd backend
go build -o server
```

### Development Workflow

1. **Backend Development**
   ```bash
   cd backend
   go run main.go  # Starts on :8080
   ```

2. **Frontend Development**
   ```bash
   cd frontend_gym
   npm run dev     # Starts on :5173
   ```

3. **Database Management**
   - Auto-migration on startup
   - Manual SQL scripts for initial data
   - Docker volume persistence

### Hot Reload & Development Features
- **Vite HMR** - Instant frontend updates
- **Go Air** - Backend hot reload (optional)
- **Docker Watch** - Container development mode

## 🚀 Deployment

### Production Docker Deployment

1. **Build production images**
   ```bash
   docker-compose build --no-cache
   ```

2. **Deploy with environment config**
   ```bash
   # Set production environment variables
   export DB_PASSWORD=secure_production_password
   export JWT_SECRET=production_jwt_secret
   
   docker-compose up -d
   ```

3. **Database initialization**
   ```bash
   docker exec -it <mysql-container> mysql -u root -p
   # Run initial SQL commands
   ```

### Production Considerations

- **Environment Variables** - Secure secret management
- **Database Backups** - Regular MySQL dumps
- **SSL/TLS** - HTTPS configuration for production
- **Load Balancing** - Multiple backend instances
- **Monitoring** - Application health checks
- **Logging** - Structured logging with levels

### Cloud Deployment Options

- **AWS ECS** - Container orchestration
- **Google Cloud Run** - Serverless containers
- **Digital Ocean App Platform** - Simple deployment
- **Kubernetes** - Enterprise orchestration

## ⚡ Performance & Scalability

### Backend Optimizations

- **Connection Pooling** - Efficient database connections
- **GORM Optimizations** - Eager loading for related data
- **Middleware Caching** - Response caching where appropriate
- **Gin Performance** - Minimal overhead HTTP framework

### Frontend Optimizations

- **Vite Bundle Optimization** - Tree shaking and code splitting
- **React Lazy Loading** - Component-based loading
- **Asset Optimization** - Image compression and caching
- **Service Worker** - Offline capability (future enhancement)

### Database Performance

- **Indexing Strategy** - Optimal query performance
- **Query Optimization** - Efficient JOIN operations
- **Connection Management** - Pool configuration
- **Migration Scripts** - Version-controlled schema changes

### Scalability Considerations

- **Horizontal Scaling** - Multiple backend instances
- **Database Sharding** - Future data distribution
- **CDN Integration** - Static asset delivery
- **Caching Layer** - Redis for session management

## 🧩 Challenges & Solutions

### 🔐 **JWT Token Management**
**Challenge**: Secure token storage and automatic refresh handling
**Solution**: Implemented comprehensive token validation with automatic logout on expiration, secure localStorage management, and HTTP interceptors for seamless user experience.

### 🔄 **Real-time Enrollment Updates**
**Challenge**: Maintaining accurate capacity counts with concurrent enrollments
**Solution**: Database-level constraints with transaction handling, real-time capacity validation, and optimistic locking to prevent overbooking.

### 🏗 **Clean Architecture Implementation**
**Challenge**: Maintaining separation of concerns across multiple layers
**Solution**: Implemented repository pattern with dependency injection, clear DTO/DAO separation, and service layer abstraction for maintainable, testable code.

### 🐳 **Docker Development Workflow**
**Challenge**: Efficient development experience with containerized services
**Solution**: Docker Compose with volume mounting, environment-specific configurations, and database persistence across container restarts.

### 🧪 **Comprehensive Testing Strategy**
**Challenge**: Testing HTTP handlers with mock dependencies
**Solution**: Implemented mock service layer with Testify framework, comprehensive controller testing, and separation of business logic for unit test coverage.

### 📱 **Responsive UI Development**
**Challenge**: Creating intuitive admin and user interfaces
**Solution**: Component-based architecture with custom hooks, efficient state management, and responsive design principles with modern CSS.

## 🤝 Contributing

We welcome contributions! Please follow these guidelines:

1. **Fork the repository**
2. **Create a feature branch** (`git checkout -b feature/amazing-feature`)
3. **Commit changes** (`git commit -m 'Add amazing feature'`)
4. **Push to branch** (`git push origin feature/amazing-feature`)
5. **Open a Pull Request**

### Development Standards

- **Go**: Follow effective Go conventions and gofmt
- **React**: Use functional components with hooks
- **Testing**: Maintain test coverage for new features
- **Documentation**: Update README for significant changes
- **Git**: Use conventional commit messages

### Code Review Process

- All changes require pull request review
- Automated tests must pass
- Code must follow established patterns
- Documentation updates required for new features

---

**Built with ❤️ by the GymFlow Team**

*This project demonstrates modern full-stack development practices, clean architecture principles, and production-ready deployment strategies suitable for enterprise-level applications.*