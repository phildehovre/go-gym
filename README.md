# Gym Membership & Location Management API

## Description
This API allows users to manage gym memberships, gym locations, employees, and other related operations. The API provides functionality for registering new members, selecting gyms, cancelling memberships, and managing gym employees and locations.

## Features
- User registration and authentication.
- Gym selection and membership management.
- Membership cancellation.
- Management of gym locations and employees.

## Prerequisites
- `Go v1.XX`
- Database (e.g., `PostgreSQL vX.X`)
- JWT library for authentication (e.g., `github.com/dgrijalva/jwt-go`)

## Installation
To get started with this API, follow the instructions below:

1. Clone the repository:
    ```bash
    git clone https://github.com/username/gym-management-api.git
    cd gym-management-api
    ```
2. Install dependencies:
    ```bash
    go get ./...
    ```
3. Set up environment variables:
    - `DB_HOST`: Database host (e.g., `localhost`)
    - `DB_USER`: Database username
    - `DB_PASS`: Database password
    - `JWT_SECRET`: Secret key for JWT authentication

4. Run migrations (if applicable):
    ```bash
    go run migrations.go
    ```

5. Start the server:
    ```bash
    go run main.go
    ```

## Usage

### Authentication
This API uses JWT (JSON Web Token) for authentication. To interact with protected endpoints, you need to obtain a token and pass it in the `Authorization` header.

1. Obtain a token by sending a POST request to `/auth/login`:
    ```bash
    curl -X POST https://yourapi.com/auth/login -d '{
        "email": "user@example.com",
        "password": "password123"
    }'
    ```
    The response will include a JWT token:
    ```json
    {
      "token": "your.jwt.token"
    }
    ```

2. Pass the token in the `Authorization` header for subsequent requests:
    ```bash
    curl -H "Authorization: Bearer <token>" https://yourapi.com/members
    ```

### Endpoints

#### `POST /members/register`
- Description: Register a new member.
- Request:
    ```json
    {
      "name": "John Doe",
      "email": "johndoe@example.com",
      "password": "password123",
      "location_id": 1
    }
    ```
- Response:
    ```json
    {
      "id": 1,
      "name": "John Doe",
      "email": "johndoe@example.com",
      "location_id": 1
    }
    ```

#### `GET /members`
- Description: Retrieve a list of all members.
- Response:
    ```json
    [
      {
        "id": 1,
        "name": "John Doe",
        "email": "johndoe@example.com",
        "location_id": 1
      }
    ]
    ```

#### `POST /members/cancel`
- Description: Cancel a membership.
- Request:
    ```json
    {
      "member_id": 1
    }
    ```
- Response:
    ```json
    {
      "message": "Membership cancelled for member_id: 1"
    }
    ```

#### `GET /locations`
- Description: Retrieve a list of gym locations.
- Response:
    ```json
    [
      {
        "id": 1,
        "name": "Downtown Gym",
        "address": "123 Main St."
      }
    ]
    ```

### Example Error Responses

- 400 Bad Request:
    ```json
    {
      "error": "Invalid data format"
    }
    ```

- 401 Unauthorized:
    ```json
    {
      "error": "Invalid token"
    }
    ```

## Error Handling
Common errors users may encounter and how to address them:

- `400 Bad Request`: Make sure all required fields are included and formatted correctly.
- `401 Unauthorized`: Ensure a valid JWT token is passed in the `Authorization` header.

## Contributing
If you'd like to contribute to this API, please follow these steps:

1. Fork the repository.
2. Create a new feature branch: `git checkout -b feature-name`.
3. Commit your changes: `git commit -m 'Add new feature'`.
4. Push to the branch: `git push origin feature-name`.
5. Submit a pull request.

## License
This project is licensed under the MIT License. See the `LICENSE` file for more details.

## Contact
For any questions or support, please contact:
- Name: Philippe De Hovre
- Email: your-email@example.com
