# Book API Service

## Overview

This project is a simple **Book API** that allows users to create, retrieve, update, and delete books. It is built using **Go**, **Goa framework**, and uses **MySQL** as the database. The project also includes unit tests using **GoMock** for service logic testing.

## Features

- **Create a Book**
- **Retrieve a Book by ID**
- **Update Book information**
- **Delete a Book**

---

## Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.16 or later)
- [MySQL](https://dev.mysql.com/downloads/mysql/)

---

## Project Setup

1. Environment Variables
Create a .env file in the project root directory to configure your MySQL connection. Here's an example:

DB_DSN="root:password@tcp(localhost:3306)/dbname"

Make sure the DB_DSN is correctly configured to point to your MySQL instance.

2. Running the application
To start the server: go run cmd/main.go

## Api Documentation

- **Base URL: http://localhost:8080**

Endpoints:

- **Create a book**
URL: http://localhost:8080/books
Method: POST
Request Body:
{
  "title": "Book Title",
  "author": "Author Name",
  "cover_url": "https://example.com/cover.jpg",
  "published_at": "2023-01-01"
}
Response Body:
{
  "id": "generated-book-id",
  "title": "Book Title",
  "author": "Author Name",
  "cover_url": "https://example.com/cover.jpg",
  "published_at": "2023-01-01"
}

- **Get a book by ID**
URL: http://localhost:8080/books/{id}
Method: GET
Response Body:
{
  "id": "generated-book-id",
  "title": "Book Title",
  "author": "Author Name",
  "cover_url": "https://example.com/cover.jpg",
  "published_at": "2023-01-01"
}

- **Update a book by ID**
URL: http://localhost:8080/books/{id}
Method: PUT
Request Body:
{
  "title": "Updated Book Title",
  "author": "Updated Author Name",
  "cover_url": "https://example.com/updated_cover.jpg",
  "published_at": "2024-01-01"
}
Response Body:
{
  "id": "book-id",
  "title": "Updated Book Title",
  "author": "Updated Author Name",
  "cover_url": "https://example.com/updated_cover.jpg",
  "published_at": "2024-01-01"

}

- **Delete a book by ID**
URL: http://localhost:8080/books/{id}
Method: DELETE


