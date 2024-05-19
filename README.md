## REST_API Using GO for Event Management using Gin and SQLite

## Description
This project is a RESTful API for managing events, built with the Gin framework in Go and SQLite for data storage. It includes JWT-based authentication and authorization to secure the endpoints and ensure that only authorized users can modify event data.

## Features
- Create, retrieve, update, and delete events
- User authentication and authorization using JWT
- Event registration and cancellation for authenticated users
- Middleware for token validation and authorization

## Technologies Used
- Go (Gin Framework)
- SQLite
- JWT for authentication

## Installation
1. **Clone the repository:**
   ```sh
   git clone https://github.com/Milan-CS03/GO_REST.git
   cd REST-GO
   ```

2. **Install dependencies:**
   ```sh
   go mod tidy
   ```

3. **Set up the database:**
   Ensure SQLite is installed and configure the database as required.

4. **Run the application:**
   ```sh
   go run main.go
   ```

## Usage
1. **Register a new user:**
   POST `/signup` with JSON payload:
   ```json
   {
     "email": "example@abc.com",
     "password": "password"
   }
   ```

2. **Login to get JWT token:**
   POST `/login` with JSON payload:
   ```json
   {
     "username": "example@abc.com",
     "password": "password"
   }
   ```

3. **Use the token to access protected endpoints:**
   Add the token to the `Authorization` header:
   ```sh
   Authorization: Bearer your_jwt_token
   ```

## API Endpoints
- **Public Endpoints:**
  - POST `/signup` - Register a new user
  - POST `/login` - Login and obtain a JWT

- **Protected Endpoints (require JWT):**
  - GET `/events` - Get a list of all events
  - POST `/events` - Create a new event
  - PUT `/events/:id` - Update an event
  - DELETE `/events/:id` - Delete an event
  - POST `/events/:id/register` - Register for an event
  - DELETE `/events/:id/register` - Cancel registration for an event

## Authentication
This project uses JWT for authentication. After logging in, include the JWT in the `Authorization` header for protected endpoints.

## Middleware
- **JWT Middleware:** Validates the token and authorizes the user.
- **Error Handling Middleware:** Manages errors and provides meaningful responses.

