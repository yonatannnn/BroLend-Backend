# BroLend-Backend Documentation

## Table of Contents
- [Overview](#overview)
- [User API](#user-api)
- [Debt API](#debt-api)
- [Authentication](#authentication)
- [Deployment (Railway)](#deployment-railway)
- [Environment Variables](#environment-variables)

---

## Overview
This backend provides user management and debt tracking APIs using Go, Gin, and MongoDB.

---

## User API

### 1. Register
- **POST** `/register`
- **Payload:**
  ```json
  {
    "user_id": "string",
    "first_name": "string",
    "last_name": "string",
    "username": "string",
    "password": "string"
  }
  ```
- **Response:**
  ```json
  {
    "message": "User created successfully",
    "userID": "string",
    "token": "string"
  }
  ```

### 2. Login
- **POST** `/login`
- **Payload:**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```
- **Response:**
  ```json
  {
    "message": "Login successful",
    "user": "string",
    "token": "string"
  }
  ```

### 3. Update User
- **PUT** `/user`
- **Auth:** Yes (JWT)
- **Payload:**
  ```json
  {
    "user_id": "string",
    "first_name": "string",
    "last_name": "string",
    "username": "string",
    "password": "string"
  }
  ```
- **Response:**
  ```json
  { "message": "User updated successfully" }
  ```

### 4. Find User by Username
- **GET** `/user/:username`
- **Auth:** Yes (JWT)
- **Response:** User object

### 5. Delete User
- **DELETE** `/user/:username`
- **Auth:** Yes (JWT)
- **Response:**
  ```json
  { "message": "User deleted successfully" }
  ```

---

## Debt API

### 1. Create Debt
- **POST** `/debt`
- **Auth:** Yes (JWT)
- **Payload:**
  ```json
  {
    "lender_id": "string",
    "amount": 100.0,
    "currency": "USD"
  }
  ```
- **Response:**
  ```json
  {
    "message": "Debt created",
    "id": "string"
  }
  ```

### 2. Lender Accepts Debt
- **POST** `/debt/:id/accept`
- **Auth:** Yes (JWT, must be lender)
- **Payload:** _None_
- **Response:**
  ```json
  { "message": "Debt accepted" }
  ```

### 3. Lender Rejects Debt
- **POST** `/debt/:id/reject`
- **Auth:** Yes (JWT, must be lender)
- **Payload:** _None_
- **Response:**
  ```json
  { "message": "Debt rejected" }
  ```

### 4. Borrower Requests Paid Approval
- **POST** `/debt/:id/request-paid`
- **Auth:** Yes (JWT, must be borrower)
- **Payload:** _None_
- **Response:**
  ```json
  { "message": "Paid approval requested" }
  ```

### 5. Lender Approves Payment
- **POST** `/debt/:id/approve-payment`
- **Auth:** Yes (JWT, must be lender)
- **Payload:** _None_
- **Response:**
  ```json
  { "message": "Payment approved" }
  ```

### 6. Lender Rejects Payment Request
- **POST** `/debt/:id/reject-payment`
- **Auth:** Yes (JWT, must be lender)
- **Payload:** _None_
- **Response:**
  ```json
  { "message": "Payment request rejected" }
  ```

### 7. Net Money Amounts for User
- **GET** `/debt/net`
- **Auth:** Yes (JWT)
- **Payload:** _None_
- **Response:**
  ```json
  {
    "USD": 100.0,
    "ETB": -50.0,
    "USDT": 0.0
  }
  ```

### 8. User Debt History
- **GET** `/debt/history`
- **Auth:** Yes (JWT)
- **Payload:** _None_
- **Response:** Array of all debt objects for the user (all statuses).

### 9. Active Incoming Debts (as Lender)
- **GET** `/debt/active-incoming`
- **Auth:** Yes (JWT)
- **Payload:** _None_
- **Response:** Array of active debt objects where you are the lender.

### 10. Active Outgoing Debts (as Borrower)
- **GET** `/debt/active-outgoing`
- **Auth:** Yes (JWT)
- **Payload:** _None_
- **Response:** Array of active debt objects where you are the borrower.

### 11. Incoming Requests for Debt (as Lender)
- **GET** `/debt/incoming-requests`
- **Auth:** Yes (JWT)
- **Payload:** _None_
- **Response:** Array of pending debt requests where you are the lender.

---

## Authentication
- All endpoints except `/register` and `/login` require a valid JWT in the `Authorization: Bearer <token>` header.
- The user's ID is extracted from the JWT and used for borrower/lender identification.

---

## Deployment (Railway)

### 1. Prepare
- Push your code to GitHub.
- Add a `.env.example` file with all required environment variables.

### 2. Create Project
- Go to [railway.app](https://railway.app/) and create a new project from your GitHub repo.

### 3. Set Environment Variables
- Add `MONGO_URI`, `MONGO_DB`, `JWT_SECRET`, `PORT` in the Railway dashboard.

### 4. Build & Start Commands
- **Build Command:** `go build -o app .`
- **Start Command:** `./app`

### 5. Deploy
- Click "Deploy" and Railway will build and run your app.
- Use the provided public URL to access your API.

---

## Environment Variables
- `MONGO_URI`: MongoDB connection string
- `MONGO_DB`: Database name
- `JWT_SECRET`: Secret for signing JWTs
- `PORT`: Port to run the server (default: 3000)

---

For more details or troubleshooting, see the comments in the code or ask your team!

# BroLend Backend

A Go-based backend application for a lending platform built with Clean Architecture principles.

## Tech Stack

- **Language**: Go 1.24.2
- **Web Framework**: Gin (Gin-Gonic)
- **Database**: MongoDB with official MongoDB Go driver
- **Authentication**: JWT (JSON Web Tokens)
- **Password Hashing**: bcrypt (via golang.org/x/crypto)
- **Architecture**: Clean Architecture (Repository Pattern)

## Project Structure

```
├── controller/          # HTTP handlers
│   └── user_controller.go
├── domain/             # Business entities and interfaces
│   ├── user.go
│   ├── debt.go
│   ├── payment.go
│   ├── net_debt.go
│   ├── request.go
│   ├── repositories.go
│   └── usecases.go
├── infrastructure/     # External services
│   ├── jwt_Service.go
│   └── password_service.go
├── repository/         # Data access layer
│   └── user_repository.go
├── router/            # HTTP routing
│   └── router.go
├── usecase/           # Business logic layer
│   └── user_usecase.go
├── utils/             # Utility functions
├── main.go           # Application entry point
├── go.mod            # Go module file
└── README.md         # This file
```

## API Endpoints

### User Management

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| POST | `/api/users/register` | Register a new user | `{"username": "string", "password": "string", "first_name": "string", "last_name": "string"}` |
| POST | `/api/users/login/:username/:password` | Login user | Path parameters |
| GET | `/api/users/:username` | Get user by username | Path parameter |
| PUT | `/api/users/update` | Update user information | `{"id": "string", "username": "string", "password": "string", "first_name": "string", "last_name": "string"}` |
| DELETE | `/api/users/:username` | Delete user | Path parameter (ObjectID) |

## Setup Instructions

### Prerequisites

1. Go 1.24.2 or later
2. MongoDB running locally on port 27017

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd BroLend-Backend
```

2. Install dependencies:
```bash
go mod tidy
```

3. Start MongoDB (if not already running):
```bash
# On Ubuntu/Debian
sudo systemctl start mongod

# Or using Docker
docker run -d -p 27017:27017 --name mongodb mongo:latest
```

4. Build and run the application:
```bash
go build -o brolend .
./brolend
```

The server will start on port 8080.

### Environment Configuration

The application currently uses default configurations:
- MongoDB URI: `mongodb://localhost:27017`
- Database: `brolend`
- Collection: `users`
- JWT Secret: `your-secret-key-here` (should be changed in production)

## Architecture Overview

This project follows Clean Architecture principles:

1. **Domain Layer**: Contains business entities (`User`, `Debt`, etc.) and interfaces for repositories and use cases
2. **Repository Layer**: Implements data access logic using MongoDB
3. **Use Case Layer**: Contains business logic and orchestrates between repositories and external services
4. **Controller Layer**: Handles HTTP requests and responses
5. **Infrastructure Layer**: Provides external services like JWT authentication and password hashing
6. **Router Layer**: Defines HTTP routes and middleware

## Features

- User registration and authentication
- JWT-based authentication
- Secure password hashing with bcrypt
- MongoDB integration
- RESTful API design
- Clean architecture implementation

## Future Enhancements

- Add debt management endpoints
- Implement payment processing
- Add request management
- Add middleware for authentication
- Add input validation
- Add logging and monitoring
- Add unit and integration tests
