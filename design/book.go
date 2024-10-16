package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("books", func() {
	Title("Books API")
	Description("API for managing books")
	Server("books", func() {
		Host("localhost", func() {
			URI("http://localhost:8080")
		})
	})
})

// Define Book Type
var Book = Type("Book", func() {
	Attribute("id", String, "ID of the book")
	Attribute("title", String, "Title of the book", func() {
		Example("The Go Programming Language")
	})
	Attribute("author", String, "Author of the book", func() {
		Example("Alan A. A. Donovan")
	})
	Attribute("cover_url", String, "URL of the book cover", func() {
		Example("https://example.com/cover.jpg")
	})
	Attribute("published_at", String, "Publish date of the book", func() {
		Example("2023-01-01")
	})
	Required("id", "title", "author", "cover_url", "published_at")
})

// Define BookPayload Type
var BookPayload = Type("BookPayload", func() {
	Attribute("title", String, "Title of the book")
	Attribute("author", String, "Author of the book")
	Attribute("cover_url", String, "URL of the book cover")
	Attribute("published_at", String, "Publish date of the book", func() {
		Example("2023-01-01")
	})
	Required("title", "author", "cover_url", "published_at")
})

// Define the Book Service
var _ = Service("book", func() {
	Description("Service for managing books")

	// Define service-level errors
	Error("NotFound", func() {
		Description("The book with the given ID was not found.")
	})
	Error("ValidationError", func() {
		Description("The input provided for the book is invalid.")
	})
	Error("InternalError", func() {
		Description("An internal error occurred while processing the request.")
	})

	// Define Create Method
	Method("create", func() {
		Payload(BookPayload)
		Result(Book)

		Error("ValidationError") // Method-specific error reusing the service-level error

		HTTP(func() {
			POST("/books")
			Response(StatusCreated)
			Response("ValidationError", StatusUnprocessableEntity)
			Response("InternalError", StatusInternalServerError)
		})
	})

	// Define Show Method
	Method("show", func() {
		Payload(func() {
			Attribute("id", String, "ID of the book")
			Required("id")
		})
		Result(Book)

		Error("NotFound") // Book not found
		Error("InternalError")

		HTTP(func() {
			GET("/books/{id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("InternalError", StatusInternalServerError)
		})
	})

	// Define Update Method
	Method("update", func() {
		Payload(func() {
			Attribute("id", String, "ID of the book")
			Attribute("book", BookPayload)
			Required("id", "book")
		})
		Result(Book)

		Error("NotFound")
		Error("ValidationError")
		Error("InternalError")

		HTTP(func() {
			PUT("/books/{id}")
			Response(StatusOK)
			Response("NotFound", StatusNotFound)
			Response("ValidationError", StatusUnprocessableEntity)
			Response("InternalError", StatusInternalServerError)
		})
	})

	// Define Delete Method
	Method("delete", func() {
		Payload(func() {
			Attribute("id", String, "ID of the book")
			Required("id")
		})

		Error("NotFound")
		Error("InternalError")

		HTTP(func() {
			DELETE("/books/{id}")
			Response(StatusNoContent)
			Response("NotFound", StatusNotFound)
			Response("InternalError", StatusInternalServerError)
		})
	})
})
