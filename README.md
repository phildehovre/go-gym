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
    go run cmd/migrate/main.go
    ```

5. Start the server:
    ```bash
    go run cmd/main.go
    ```

## Usage

### Authentication
This API uses JWT (JSON Web Token) for authentication. To interact with protected endpoints, you need to obtain a token and pass it in the `Authorization` header.

1. Obtain a token by sending a POST request to `/auth/login`:
    ```bash
    curl -X POST http://localhost:<port>/api/v1/register -d '{
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
    curl -H "Authorization: <token> http://localhost:<port>/api/v1/login
    ```

### Endpoints

IMPORTANT: All endpoints should be prefixed with `/api/v1`

## API Endpoints

### Authentication
This API uses JWT for authentication. Admin routes require an admin role, enforced via `auth.AdminMiddleware`.

### Membership Service

#### `POST /membership`
- **Description**: Create a new membership for the authenticated user.
- **Request**:
    ```json
    {
      "MembershipType": "Gold",
      "Status": "Active",
      "StartDate": "2024-01-01",
      "EndDate": "2024-12-31",
      "LocationIDS": [1, 2, 3]
    }
    ```
- **Response**:
    ```json
    {
      "status": "Membership created"
    }
    ```

#### `GET /membership`
- **Description**: Retrieve the authenticated user's membership details.
- **Response**:
    ```json
    {
      "UserID": 1,
      "MembershipType": "Gold",
      "Status": "Active",
      "StartDate": "2024-01-01",
      "EndDate": "2024-12-31",
      "Locations": ["Location1", "Location2"]
    }
    ```

#### `PATCH /membership`
- **Description**: Update membership details for the authenticated user.
- **Request**:
    ```json
    {
      "MembershipType": "Platinum",
      "Status": "Active",
      "StartDate": "2024-01-01",
      "EndDate": "2025-12-31"
    }
    ```
- **Response**:
    ```json
    {
      "status": "Membership updated"
    }
    ```

#### `DELETE /membership`
- **Description**: Deactivate the authenticated user's membership.
- **Response**:
    ```json
    {
      "status": "Membership deactivated"
    }
    ```

#### `GET /membership/locations`
- **Description**: Retrieve available gym locations for memberships.
- **Response**:
    ```json
    [
      {
        "id": 1,
        "name": "Downtown Gym",
        "address": "123 Main St."
      }
    ]
    ```

#### `PATCH /membership/renew`
- **Description**: Renew the authenticated user's membership.
- **Response**:
    ```json
    {
      "status": "Membership renewed"
    }
    ```

---

### User Service

#### `POST /login`
- **Description**: Log in a user and retrieve a JWT token.
- **Request**:
    ```json
    {
      "email": "user@example.com",
      "password": "password123"
    }
    ```
- **Response**:
    ```json
    {
      "token": "your.jwt.token"
    }
    ```

#### `POST /register`
- **Description**: Register a new user.
- **Request**:
    ```json
    {
      "FirstName": "John",
      "LastName": "Doe",
      "Email": "johndoe@example.com",
      "Password": "password123"
    }
    ```
- **Response**:
    ```json
    {
      "status": "User registered"
    }
    ```

#### `GET /users`
- **Description**: Retrieve a list of all users (requires admin access).
- **Response**:
    ```json
    [
      {
        "id": 1,
        "FirstName": "John",
        "LastName": "Doe",
        "Email": "johndoe@example.com"
      }
    ]
    ```

#### `PATCH /users`
- **Description**: Update the authenticated user's details.
- **Request**:
    ```json
    {
      "FirstName": "Jane",
      "LastName": "Doe"
    }
    ```
- **Response**:
    ```json
    {
      "status": "User updated"
    }
    ```

---

### Admin Service

#### `GET /admin/users`
- **Description**: Retrieve a list of all users (requires admin access).
- **Response**:
    ```json
    [
      {
        "id": 1,
        "FirstName": "John",
        "LastName": "Doe",
        "Email": "johndoe@example.com"
      }
    ]
    ```

#### `GET /admin/memberships`
- **Description**: Retrieve all memberships in the system (requires admin access).
- **Response**:
    ```json
    [
      {
        "UserID": 1,
        "MembershipType": "Gold",
        "Status": "Active",
        "StartDate": "2024-01-01",
        "EndDate": "2024-12-31"
      }
    ]
    ```

---

### Locations Service

#### `POST /location`
- **Description**: Create a new location.
- **Request**:
    ```json
    {
      "Name": "Central Gym",
      "Address": "123 Fitness Ave",
      "City": "Metropolis",
      "State": "State",
      "PostalCode": "12345",
      "Country": "Country",
      "PhoneNumber": "123-456-7890",
      "Email": "centralgym@example.com",
      "Capacity": 100,
      "OperatingHours": "9 AM - 9 PM",
      "IsActive": true
    }
    ```
- **Response**:
    ```json
    {
      "status": "Location created"
    }
    ```

#### `GET /location`
- **Description**: Retrieve all available locations.
- **Response**:
    ```json
    [
      {
        "id": 1,
        "Name": "Central Gym",
        "Address": "123 Fitness Ave",
        "City": "Metropolis",
        "State": "State",
        "PostalCode": "12345",
        "Country": "Country",
        "PhoneNumber": "123-456-7890",
        "Email": "centralgym@example.com",
        "Capacity": 100,
        "OperatingHours": "9 AM - 9 PM",
        "IsActive": true
      }
    ]
    ```

#### `GET /location/{id}`
- **Description**: Retrieve details for a specific location by ID.
- **Response**:
    ```json
    {
      "id": 1,
      "Name": "Central Gym",
      "Address": "123 Fitness Ave",
      "City": "Metropolis",
      "State": "State",
      "PostalCode": "12345",
      "Country": "Country",
      "PhoneNumber": "123-456-7890",
      "Email": "centralgym@example.com",
      "Capacity": 100,
      "OperatingHours": "9 AM - 9 PM",
      "IsActive": true
    }
    ```

---

### Common Error Responses

- **400 Bad Request**:
    ```json
    {
      "error": "Invalid data format"
    }
    ```

- **401 Unauthorized**:
    ```json
    {
      "error": "Invalid token"
    }
    ```

- **403 Forbidden**:
    ```json
    {
      "error": "You do not have permission to access this resource"
    }
    ```

- **404 Not Found**:
    ```json
    {
      "error": "Resource not found"
    }
    ```


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
- Email: me@philippedehovre.com
