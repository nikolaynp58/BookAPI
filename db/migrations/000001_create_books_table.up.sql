-- migrations/000001_create_books_table.up.sql
CREATE TABLE books (
    id VARCHAR(15) PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    cover_url VARCHAR(255) NOT NULL,
    published_at VARCHAR(25) NOT NULL
);
