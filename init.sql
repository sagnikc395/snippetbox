CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE snippetbox;
-- Create snippets table
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);
-- Add indexes
CREATE INDEX idx_snippets_created ON snippets(created);
-- add some dummy records 
INSERT INTO snippets(title, content, created, expires)
VALUES(
        'An old silent pond',
        'An old silent pond...\nA frog jumps into the pond,\nsplash!Silence again.',
        UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
    );
INSERT INTO snippets(title, content, created, expires)
VALUES(
        'Over the wintry forest',
        'Over the wintry \nforest, winds howl in rage\n with no leaves to blow.\n\n',
        UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
    );
INSERT INTO snippets(title, content, created, expires)
VALUES (
        'First autumn morning',
        "First autumn morning\nthe mirror I state into\nshowsmy father's face.\n\n",
        UTC_TIMESTAMP(),
        DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
    );