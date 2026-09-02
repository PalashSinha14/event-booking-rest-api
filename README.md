# 🎟️ Event Booking REST API

**A production-ready backend service for creating, browsing, and registering for events — built with Go, Gin, and PostgreSQL.**

![Go]
![Gin]
![PostgreSQL]
![Docker]
![JWT]

**🔗 Live Demo:** [event-booking-rest-api.onrender.com](https://event-booking-rest-api.onrender.com)
*(Deployed on Render, built and run from this project's Docker image)*

---

## Overview

This project is a production-ready RESTful backend service built using Golang and the Gin Web Framework. It enables users to register, authenticate, create events, and register for events.

The system is fully Dockerized and supports environment-based configuration, making it suitable for local development as well as cloud deployment (e.g., AWS EC2). It demonstrates backend system design fundamentals including authentication, database management, containerization, and service networking.

## Architecture

```
Client → Gin HTTP Server → Middleware (JWT Auth) → Service Layer → PostgreSQL Database
```

The application uses:
- RESTful routing principles
- JWT-based authentication
- SQL database schema with relational integrity
- Docker multi-container setup
- Environment-based configuration switching
- Retry logic for database readiness in containerized environments

## Features

### 🔐 Authentication & Security
- User signup and login
- JWT-based authentication
- Protected routes using middleware
- Password hashing using bcrypt
- Token validation for secured endpoints

### 📅 Event Management
- Create events
- View all events
- Delete events
- Register for events
- Prevent duplicate registrations
- View events you created (**My Events**), with the option to delete them
- View events you've registered for (**My Registrations**), with the option to cancel a registration

### 🖥️ Web Interface
- Server-rendered HTML pages (Go `html/template` via Gin) for signup, login, dashboard, event browsing, and event creation
- Shared navigation bar (reusable Go template partial) across all logged-in pages
- Dashboard with quick-access action cards linking to every core feature

### 🗄️ Database Layer
- PostgreSQL relational database
- Automatic table creation on startup
- Connection pooling configuration
- Retry mechanism for database readiness
- Environment-based DB configuration (Local vs Docker)

### 🚀 DevOps & Deployment
- Dockerized backend
- Docker Compose multi-container setup
- Service-to-service networking using container DNS
- Volume persistence for PostgreSQL
- Cloud deployment ready (EC2 compatible)
- Live deployment on Render, built and run from the project's Docker image

## Tech Stack

| Category | Technology |
|---|---|
| Language | Go |
| Web Framework | Gin |
| Database | PostgreSQL |
| Containerization | Docker, Docker Compose |
| Authentication | JWT (JSON Web Tokens), bcrypt |
| Query Layer | Raw SQL (`database/sql`) |

## Project Structure

```
.
├── db/                # Database initialization and connection logic
├── models/            # Database models and SQL operations
├── routes/            # Route handlers
├── middlewares/       # JWT authentication middleware
├── utils/             # Helper utilities (token generation, hashing)
├── frontend/          # Server-rendered HTML pages, shared navbar partial, CSS and JS assets
├── docker-compose.yml
├── Dockerfile
├── .env.local
├── .env.docker
├── main.go
└── go.mod
```

## Environment Configuration

The project uses environment-based configuration.

### Local Development (`.env.local`)

```
ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=eventsdb
DB_SSLMODE=disable
```

**Run locally:**
```bash
go run main.go
```

### Docker Environment (`.env.docker`)

```
ENV=docker
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=eventsdb
DB_SSLMODE=disable
```

**Run with Docker:**
```bash
docker compose up --build
```

## API Endpoints

### Authentication
| Method | Endpoint | Description |
|---|---|---|
| POST | `/signup` | Create new user account |
| POST | `/login` | Authenticate user and receive JWT |

### Events
| Method | Endpoint | Description | Auth |
|---|---|---|---|
| GET | `/events` | Fetch all events | — |
| GET | `/events/:id` | Fetch single event | — |
| POST | `/events` | Create new event | 🔒 Protected |
| PUT | `/events/:id` | Update event | 🔒 Protected |
| DELETE | `/events/:id` | Delete event | 🔒 Protected |

### Registration
| Method | Endpoint | Description | Auth |
|---|---|---|---|
| POST | `/events/:id/register` | Register logged-in user for event | 🔒 Protected |
| DELETE | `/events/:id/register` | Cancel registration for logged-in user | 🔒 Protected |

## Database Design

**Users Table**
- `id` (Primary Key)
- `email` (Unique)
- `password` (Hashed)

**Events Table**
- `id` (Primary Key)
- `name`
- `description`
- `location`
- `dateTime`
- `user_id` (Foreign Key)

**Registrations Table**
- `id`
- `event_id` (Foreign Key)
- `user_id` (Foreign Key)
- Unique constraint to prevent duplicate registrations

## Key Backend Concepts Implemented

- REST API design
- SQL schema modeling
- Foreign key relationships
- JWT authentication flow
- Middleware-based authorization
- Secure password storage
- Docker container networking
- Environment switching for configuration
- Connection retry strategy
- Connection pooling optimization

## How It Works (Flow)

1. User signs up → password hashed → stored in DB
2. User logs in → JWT issued
3. User creates event (JWT required)
4. Other users can register for the event
5. Duplicate registration is prevented at the DB level

## Why This Project Matters

This project demonstrates:
- Backend system architecture understanding
- Production-ready environment handling
- Containerization knowledge
- Real-world authentication flow
- Database relationship modeling
- Cloud deployment readiness

It simulates real backend engineering practices used in modern SaaS applications.

## Deployment Readiness

The application is **currently deployed on Render**, built and run from the project's Docker image: **[event-booking-rest-api.onrender.com](https://event-booking-rest-api.onrender.com)**

It is also containerized and ready to be deployed on:
- AWS EC2
- Any Linux VM with Docker
- Cloud container environments

It uses:
- Multi-container orchestration
- Service discovery via container DNS
- Persistent PostgreSQL volumes

## Future Improvements

- [ ] Role-based access control
- [ ] Pagination for event listing
- [ ] Event search & filtering
- [ ] Refresh tokens
- [ ] CI/CD pipeline
- [ ] Logging & monitoring integration
- [ ] Swagger API documentation

## Author

**Palash Sinha**
Backend Developer (Golang Focused)
