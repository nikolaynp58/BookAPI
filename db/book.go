package db

import (
	gen "bookAPI/gen/book"
	"context"
)

type Book struct {
	ID          string `db:"id" json:"id"`
	Title       string `db:"title" json:"title"`
	Author      string `db:"author" json:"author"`
	CoverURL    string `db:"cover_url" json:"cover_url"`
	PublishedAt string `db:"published_at" json:"published_at"`
}

// CreateBook inserts a new book into the database.
func (r *BookStorage) CreateBook(ctx context.Context, book *gen.Book) error {
	query := "INSERT INTO books (id, title, author, cover_url, published_at) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, book.ID, book.Title, book.Author, book.CoverURL, book.PublishedAt)
	return err
}

// GetBookByID retrieves a book from the database by ID.
func (r *BookStorage) GetBookByID(ctx context.Context, id string) (*gen.Book, error) {

	var book Book
	query := "SELECT id, title, author, cover_url, published_at FROM books WHERE id = ?"
	err := r.db.GetContext(ctx, &book, query, id)
	if err != nil {
		return nil, err // Return nil if not found or an error occurs
	}

	return &gen.Book{
		ID:          book.ID,
		Title:       book.Title,
		Author:      book.Author,
		CoverURL:    book.CoverURL,
		PublishedAt: book.PublishedAt,
	}, nil
}

// UpdateBook updates an existing book in the database.
func (r *BookStorage) UpdateBook(ctx context.Context, book *gen.Book) error {
	query := "UPDATE books SET title = ?, author = ?, cover_url = ?, published_at = ? WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, book.Title, book.Author, book.CoverURL, book.PublishedAt, book.ID)
	return err
}

// DeleteBook removes a book from the database by ID.
func (r *BookStorage) DeleteBook(ctx context.Context, id string) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
