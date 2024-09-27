CREATE TABLE membershipLocations (
    membership_id INT NOT NULL,           -- Foreign key to Memberships table
    location_id INT NOT NULL,             -- Foreign key to Locations table
    PRIMARY KEY (membership_id, location_id),
    FOREIGN KEY (membership_id) REFERENCES memberships(id),
    FOREIGN KEY (location_id) REFERENCES locations(id)
);