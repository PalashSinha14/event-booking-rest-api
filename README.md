Event Booking REST API

Production-Ready Backend Service using Golang, Gin & PostgreSQL

Overview

This project is a production-ready RESTful backend service built using Golang and the Gin Web Framework. It enables users to register, authenticate, create events, and register for events.

The system is fully Dockerized and supports environment-based configuration, making it suitable for local development as well as cloud deployment (e.g., AWS EC2).

This project demonstrates backend system design fundamentals including authentication, database management, containerization, and service networking.

Architecture Overview

Client → Gin HTTP Server → Middleware (JWT Auth) → Service Layer → PostgreSQL Database

The application uses:

RESTful routing principles

JWT-based authentication

SQL database schema with relational integrity

Docker multi-container setup

Environment-based configuration switching

Retry logic for database readiness in containerized environments

Features
Authentication & Security

User signup and login

JWT-based authentication

Protected routes using middleware

Password hashing using bcrypt

Token validation for secured endpoints

Event Management

Create events

View all events

Delete events

Register for events

Prevent duplicate registrations

Database Layer

PostgreSQL relational database

Automatic table creation on startup

Connection pooling configuration

Retry mechanism for database readiness

Environment-based DB configuration (Local vs Docker)

DevOps & Deployment

Dockerized backend

Docker Compose multi-container setup

Service-to-service networking using container DNS

Volume persistence for PostgreSQL

Cloud deployment ready (EC2 compatible)

Tech Stack

Golang

Gin Web Framework

PostgreSQL

Docker

Docker Compose

JWT (JSON Web Tokens)

bcrypt password hashing

SQL

Project Structure
.
├── db/                # Database initialization and connection logic
├── models/            # Database models and SQL operations
├── routes/            # Route handlers
├── middlewares/       # JWT authentication middleware
├── utils/             # Helper utilities (token generation, hashing)
├── docker-compose.yml
├── Dockerfile
├── .env.local
├── .env.docker
├── main.go
├── go.mod
Environment Configuration

The project uses environment-based configuration.

Local Development (.env.local)
ENV=local
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=eventsdb
DB_SSLMODE=disable

Run locally:

go run main.go
Docker Environment (.env.docker)
ENV=docker
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=eventsdb
DB_SSLMODE=disable

Run with Docker:

docker compose up --build
API Endpoints
Authentication

POST /signup
Create new user account

POST /login
Authenticate user and receive JWT

Events

GET /events
Fetch all events

POST /events
Create new event (Protected)

DELETE /events/:id
Delete event (Protected)

Registration

POST /events/:id/register
Register logged-in user for event

Database Design
Users Table

id (Primary Key)

email (Unique)

password (Hashed)

Events Table

id (Primary Key)

name

description

location

dateTime

user_id (Foreign Key)

Registrations Table

id

event_id (Foreign Key)

user_id (Foreign Key)

Unique constraint to prevent duplicate registrations

Key Backend Concepts Implemented

REST API design

SQL schema modeling

Foreign key relationships

JWT authentication flow

Middleware-based authorization

Secure password storage

Docker container networking

Environment switching for configuration

Connection retry strategy

Connection pooling optimization

How It Works (Flow)

User signs up → Password hashed → Stored in DB

User logs in → JWT issued

User creates event (JWT required)

Other users can register for event

Duplicate registration is prevented at DB level

Why This Project Matters

This project demonstrates:

Backend system architecture understanding

Production-ready environment handling

Containerization knowledge

Real-world authentication flow

Database relationship modeling

Cloud deployment readiness

It simulates real backend engineering practices used in modern SaaS applications.

Deployment Readiness

The application is containerized and ready to be deployed on:

AWS EC2

Any Linux VM with Docker

Cloud container environments

It uses:

Multi-container orchestration

Service discovery via container DNS

Persistent PostgreSQL volumes

Future Improvements

Role-based access control
Pagination for event listing
Event search & filtering
Refresh tokens
CI/CD pipeline
Logging & monitoring integration
Swagger API documentation

Author

Palash
Backend Developer (Golang Focused)