DELETE FROM users;

DELETE FROM tokens;

DELETE FROM characters;

DELETE FROM user_characters;

DELETE FROM character_probability;

INSERT INTO users 
(id, name, created_at, updated_at, deleted_at) 
VALUES 
('572146c4-b7c2-439a-9c76-2c0c7d40687c', 'testName', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);

INSERT INTO tokens
(user_id, token, created_at)
VALUES
('572146c4-b7c2-439a-9c76-2c0c7d40687c', '?+6p,Jr_!O', CURRENT_TIMESTAMP);

INSERT INTO characters
(id, name, created_at, updated_at)
VALUES
(1, 'S', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'A', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(3, 'B', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(4, 'C', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO character_probability
(id, probability, character_id, created_at, updated_at, deleted_at)
VALUES
(1, 1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(2, 9, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(3, 20, 3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL),
(4, 70, 4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);