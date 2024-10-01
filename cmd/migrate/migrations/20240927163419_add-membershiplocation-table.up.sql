CREATE TABLE membershipLocations (
    membership_id INT NOT NULL,          
    location_id INT NOT NULL,            
    PRIMARY KEY (membership_id, location_id),
    FOREIGN KEY (membership_id) REFERENCES memberships(id),
    FOREIGN KEY (location_id) REFERENCES locations(id)
);