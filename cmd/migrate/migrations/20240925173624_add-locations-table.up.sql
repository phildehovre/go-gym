CREATE TABLE
    IF NOT EXISTS locations (
        `id` INT AUTO_INCREMENT PRIMARY KEY,
        `name` VARCHAR(100) NOT NULL,
        `address` VARCHAR(255) NOT NULL,
        `city` VARCHAR(100) NOT NULL,
        `state` VARCHAR(100) NOT NULL,
        `postal_code` VARCHAR(20) NOT NULL,
        `country` VARCHAR(100) NOT NULL,
        `phone_number` VARCHAR(20),
        `email` VARCHAR(100),
        `capacity` INT,
        `operating_hours` VARCHAR(100),
        `is_active` BOOLEAN DEFAULT TRUE,
        `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
        UNIQUE(name, address, city)
    );