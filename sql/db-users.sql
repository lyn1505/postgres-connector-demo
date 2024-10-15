-- Create the users table
CREATE TABLE IF NOT EXISTS "users"
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Insert data into users table with explicit ids
INSERT INTO "users" (id, name)
VALUES (1, 'Alice'),
       (2, 'Bob'),
       (3, 'Charlie');

-- Reset the sequence to start from the max id in the users table
SELECT setval(pg_get_serial_sequence('users', 'id'), MAX(id)) FROM users;
