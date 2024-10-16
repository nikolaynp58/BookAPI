package service

import (
	"bookAPI/db"
	gen "bookAPI/gen/book"
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type BookService struct {
	db db.Storage
}

func NewBookService(db db.Storage) *BookService {
	return &BookService{db: db}
}

func (s *BookService) Create(ctx context.Context, req *gen.BookPayload) (res *gen.Book, err error) {
	if req == nil {
		return nil, gen.MakeValidationError(fmt.Errorf("request payload is nil"))
	}

	id, err := generateRandomID(15)
	if err != nil {
		return nil, gen.MakeInternalError(fmt.Errorf("failed to generate random ID: %w", err))
	}
	fmt.Println("Generated ID:", id)

	book := &gen.Book{
		ID:          id,
		Title:       req.Title,
		Author:      req.Author,
		CoverURL:    req.CoverURL,
		PublishedAt: req.PublishedAt,
	}

	err = s.db.CreateBook(ctx, book)
	if err != nil {
		return book, gen.MakeInternalError(fmt.Errorf("failed to create book: %w", err))
	}

	return book, nil
}

func (s *BookService) Show(ctx context.Context, req *gen.ShowPayload) (res *gen.Book, err error) {
	if req == nil {
		return nil, gen.MakeValidationError(fmt.Errorf("request payload is nil"))
	}

	book, err := s.db.GetBookByID(ctx, req.ID)
	if err == sql.ErrNoRows {
		return nil, gen.MakeNotFound(fmt.Errorf("book with ID %s not found", req.ID))
	} else if err != nil {
		return nil, gen.MakeInternalError(fmt.Errorf("failed to get book: %w", err))
	}

	return book, nil
}

func (s *BookService) Update(ctx context.Context, req *gen.UpdatePayload) (res *gen.Book, err error) {

	if req == nil {
		return nil, gen.MakeValidationError(fmt.Errorf("request payload is nil"))
	}

	book := &gen.Book{
		ID:          req.ID,
		Title:       req.Book.Title,
		Author:      req.Book.Author,
		CoverURL:    req.Book.CoverURL,
		PublishedAt: req.Book.PublishedAt,
	}

	err = s.db.UpdateBook(ctx, book)
	if err == sql.ErrNoRows {
		return nil, gen.MakeNotFound(fmt.Errorf("book with ID %s not found", req.ID))
	} else if err != nil {
		return nil, gen.MakeInternalError(fmt.Errorf("failed to update book: %w", err))
	}

	return book, nil
}

func (s *BookService) Delete(ctx context.Context, req *gen.DeletePayload) (err error) {
	if req == nil {
		return gen.MakeValidationError(fmt.Errorf("request payload is nil"))
	}

	err = s.db.DeleteBook(ctx, req.ID)
	if err == sql.ErrNoRows {
		return gen.MakeNotFound(fmt.Errorf("book with ID %s not found", req.ID))
	} else if err != nil {
		return gen.MakeInternalError(fmt.Errorf("failed to delete book: %w", err))
	}

	return nil
}

func generateRandomID(length int) (string, error) {
	// Create a byte slice to hold the random bytes
	bytes := make([]byte, length)

	// Fill the byte slice with random bytes
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	// Convert random bytes to characters from the letters set
	for i := 0; i < length; i++ {
		bytes[i] = letters[int(bytes[i])%len(letters)]
	}

	return string(bytes), nil
}
