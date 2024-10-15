-- Create the games table
CREATE TABLE IF NOT EXISTS games
(
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(255) NOT NULL,
    admin_id INTEGER NOT NULL -- Reference to user id, no FK constraint
);

-- Insert data into games table using concrete admin ids
INSERT INTO games (name, admin_id)
VALUES ('Origins', 1), -- Alice's ID is 1
       ('Classic', 2), -- Bob's ID is 2
       ('Very Cool Game', 3); -- Charlie's ID is 3
