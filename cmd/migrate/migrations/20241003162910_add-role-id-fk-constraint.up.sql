ALTER TABLE users
ADD CONSTRAINT fk_role_id
FOREIGN KEY (role_id) REFERENCES roles(id);