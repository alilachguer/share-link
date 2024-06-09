-- Add up migration script here
CREATE TABLE IF NOT EXISTS sharelinks (
    link varchar(255) UNIQUE,
    redirect varchar(255),
    visited integer
);


INSERT INTO sharelinks(link, redirect) VALUES ('share.link/1234', 'github.com/alilachguer')
