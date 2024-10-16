package service_test

import (
	"bookAPI/gen/book"
	"bookAPI/mocks"
	"bookAPI/service"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBookService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock for the BookStorageInterface
	mockStorage := mocks.NewMockStorage(ctrl)

	// Define the input payload
	req := &book.BookPayload{
		Title:       "Test Title",
		Author:      "Test Author",
		CoverURL:    "https://example.com/cover.jpg",
		PublishedAt: "2023-01-01",
	}

	// Expect the CreateBook method to be called with any context and book, and return nil error
	mockStorage.EXPECT().CreateBook(gomock.Any(), gomock.Any()).Return(nil)

	// Create the service with the mock storage
	bookService := service.NewBookService(mockStorage)

	// Call the Create method
	result, err := bookService.Create(context.Background(), req)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, req.Title, result.Title)
	assert.Equal(t, req.Author, result.Author)
	assert.Equal(t, req.CoverURL, result.CoverURL)
	assert.Equal(t, req.PublishedAt, result.PublishedAt)
}

func TestBookService_Show(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock for the BookStorageInterface
	mockStorage := mocks.NewMockStorage(ctrl)

	// Define the input payload
	req := &book.ShowPayload{
		ID: "test-book-id",
	}

	// Define the expected book
	expectedBook := &book.Book{
		ID:          "test-book-id",
		Title:       "Test Title",
		Author:      "Test Author",
		CoverURL:    "https://example.com/cover.jpg",
		PublishedAt: "2023-01-01",
	}

	// Expect the GetBookByID method to be called and return the expected book
	mockStorage.EXPECT().GetBookByID(gomock.Any(), "test-book-id").Return(expectedBook, nil)

	// Create the service with the mock storage
	bookService := service.NewBookService(mockStorage)

	// Call the Show method
	result, err := bookService.Show(context.Background(), req)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedBook.ID, result.ID)
	assert.Equal(t, expectedBook.Title, result.Title)
	assert.Equal(t, expectedBook.Author, result.Author)
	assert.Equal(t, expectedBook.CoverURL, result.CoverURL)
	assert.Equal(t, expectedBook.PublishedAt, result.PublishedAt)
}

func TestBookService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock for the BookStorageInterface
	mockStorage := mocks.NewMockStorage(ctrl)

	// Define the input payload
	req := &book.UpdatePayload{
		ID: "test-book-id",
		Book: &book.BookPayload{
			Title:       "Updated Title",
			Author:      "Updated Author",
			CoverURL:    "https://example.com/updated_cover.jpg",
			PublishedAt: "2024-01-01",
		},
	}

	// Define the expected book to be updated
	expectedBook := &book.Book{
		ID:          "test-book-id",
		Title:       "Updated Title",
		Author:      "Updated Author",
		CoverURL:    "https://example.com/updated_cover.jpg",
		PublishedAt: "2024-01-01",
	}

	// Expect the UpdateBook method to be called and return nil (no error)
	mockStorage.EXPECT().UpdateBook(gomock.Any(), gomock.Any()).Return(nil)

	// Create the service with the mock storage
	bookService := service.NewBookService(mockStorage)

	// Call the Update method
	result, err := bookService.Update(context.Background(), req)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedBook.ID, result.ID)
	assert.Equal(t, expectedBook.Title, result.Title)
	assert.Equal(t, expectedBook.Author, result.Author)
	assert.Equal(t, expectedBook.CoverURL, result.CoverURL)
	assert.Equal(t, expectedBook.PublishedAt, result.PublishedAt)
}

func TestBookService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock for the BookStorageInterface
	mockStorage := mocks.NewMockStorage(ctrl)

	// Define the input payload
	req := &book.DeletePayload{
		ID: "test-book-id",
	}

	// Expect the DeleteBook method to be called and return nil (no error)
	mockStorage.EXPECT().DeleteBook(gomock.Any(), "test-book-id").Return(nil)

	// Create the service with the mock storage
	bookService := service.NewBookService(mockStorage)

	// Call the Delete method
	err := bookService.Delete(context.Background(), req)

	// Assertions
	assert.NoError(t, err)
}
