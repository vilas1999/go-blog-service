# Go REST API 

This project is a simple REST API service built using Go, GORM, Fiber, and JWT for authentication. It allows users to sign up, log in, and create posts. The project was developed with the primary intention of learning how to build RESTful APIs in Go, handle user authentication with JSON Web Tokens (JWT), and manage database relationships using GORM.

## Features

- **User Registration**: Users can sign up by providing their email, name, and password.
- **User Authentication**: Users can log in with their credentials to obtain a JWT token.
- **JWT Protected Routes**: Only authenticated users (with a valid JWT token) can create and manage posts.
- **Post Creation**: Authenticated users can create posts and assign content to them.

## Tech Stack

- **Go**: Main programming language.
- **GORM**: ORM library for interacting with the database.
- **Go Fiber**: Web framework for building the API.
- **JWT**: JSON Web Token for securing routes.
- **PostgreSQL**: Database for storing users and posts.

## Endpoints

### Authentication Endpoints
- **POST /user/signup**: Register a new user.
    - Payload:
      ```json
      {
        "name": "John Doe",
        "email": "john@example.com",
        "password": "password123"
      }
      ```
- **POST /user/login**: Log in and receive a JWT token.
    - Payload:
      ```json
      {
        "email": "john@example.com",
        "password": "password123"
      }
      ```

### Posts Endpoints (Protected with JWT)
- **POST /posts**: Create a new post (Authenticated users only).
    - Header: 
      ```json
      {
        "Authorization": "Bearer <JWT token>"
      }
      ```
    - Payload:
      ```json
      {
        "title": "My First Post",
        "content": "This is the content of my first post."
      }
      ```

- **GET /posts**: Fetch all posts (Authenticated users only).
    - Header: 
      ```json
      {
        "Authorization": "Bearer <JWT token>"
      }
      ```

## Future Improvements

* Add features like updating and deleting posts.
* Implement password hashing and validation.
* Add unit tests and integration tests.
* Improve error handling and input validation.

## Lessons Learned

This project served as a great way to understand:

* Learning go programming language
* How to structure a Go project for REST API development.
* Implementing JWT for authentication and securing API endpoints.
* Using GORM for ORM, migrations, and establishing foreign key relationships between models.
* Utilizing Go Fiber for creating a lightweight and efficient web server.
