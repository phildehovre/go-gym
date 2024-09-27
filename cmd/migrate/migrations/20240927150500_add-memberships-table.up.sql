CREATE TABLE
    IF NOT EXISTS memberships (
        `id` INT AUTO_INCREMENT PRIMARY KEY,
        `user_id` INT UNSIGNED NOT NULL, -- Foreign key to Users table
        `membership_type` VARCHAR(50) NOT NULL, -- 'Single' or 'Multi'
        `status` ENUM ('Active', 'Expired', 'Canceled') DEFAULT 'Active',
        `start_date` DATE NOT NULL,
        `end_date` DATE NOT NULL,
        `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(id)
    );