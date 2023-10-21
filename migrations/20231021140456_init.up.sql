-- Add up migration script here
CREATE TABLE IF NOT EXISTS sharelinks (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    link varchar(255),
    redirect varchar(100),
    visited int
);


INSERT INTO sharelinks(link, redirect) VALUES ('share.link/1234', 'github.com/alilachguer')
