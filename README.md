# API Name

## Description

A brief description of what your API does. Explain the main use case, and why someone would want to use it.

## Features

- List key features of the API.
- Example: User management, gym membership management, etc.

## Prerequisites

- List the necessary software, libraries, and versions required to run the API. For example:
  - `Go v1.XX`
  - `PostgreSQL vX.X`

## Installation

To get started with this API, follow the instructions below:

1. Clone the repository:
   ```bash
   git clone https://github.com/username/repository.git
   cd repository
   ```
2. Install dependencies:
   ```bash
   go get ./...
   ```
3. Set up environment variables:

   - `DB_HOST`: Database host
   - `DB_USER`: Database username
   - `DB_PASS`: Database password

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

If your API requires authentication, explain how users should authenticate (e.g., using JWT tokens).

1. Obtain a token by sending a POST request to `/auth/login`:

   ```bash
   curl -X POST https://yourapi.com/auth/login -d '{
       "username": "user",
       "password": "pass"
   }'
   ```

   The response will include a JWT token to be used in subsequent requests.

2. Pass the token in the `Authorization` header for requests:
   ```bash
   curl -H "Authorization: Bearer <token>" https://yourapi.com/protected-endpoint
   ```

### Endpoints

#### `GET /users`

- Description: Retrieve a list of all users.
- Response:
  ```json
  [
    {
      "id": 1,
      "username": "exampleuser",
      "email": "user@example.com"
    }
  ]
  ```

#### `POST /users`

- Description: Add a new user.
- Request:
  ```json
  {
    "username": "newuser",
    "email": "newuser@example.com",
    "password": "password123"
  }
  ```
- Response:
  ```json
  {
    "id": 2,
    "username": "newuser",
    "email": "newuser@example.com"
  }
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

Here are some common errors users may encounter and how to address them:

- `400 Bad Request`: Ensure all required fields are included.
- `401 Unauthorized`: Ensure a valid token is passed in the `Authorization` header.

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
