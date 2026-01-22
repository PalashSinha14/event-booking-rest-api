# Event Booking REST API (Go)

A RESTful Event Booking API built using **Golang**, **Gin**, and **SQLite**.

## Features
- User Signup & Login (JWT Authentication)
- Create, Read, Update, Delete Events
- Secure password hashing
- RESTful API conventions
- SQLite database

## Tech Stack
- Go (Golang)
- Gin Web Framework
- SQLite db
- JWT Authentication

## API Endpoints

### Auth
- POST `/signup`
- POST `/login`

### Events
- GET `/events`
- GET `/events/:id`
- POST `/events`
- PUT `/events/:id`
- DELETE `/events/:id`

## How to Run

```bash
go mod tidy
go run main.go
