CREATE TABLE roles (
    id INT PRIMARY KEY,       -- Unique identifier for each role
    name VARCHAR(50) NOT NULL,   -- Role name (e.g., "Admin", "User")
    description TEXT             -- Optional description of the role
);