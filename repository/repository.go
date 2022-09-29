package repository

import "grpc-unary/model"

// localstorage
var storage []model.Book = []model.Book{}

// Add Book
func AddBook(book model.Book) model.Book {
	storage = append(storage, book)
	return book
}

// Get Book
func GetBook(bookId string) (int, model.Book) {
	for i, v := range storage {
		if v.Id == bookId {
			return i, v
		}

	}
	return 0, model.Book{}
}

// UpdateBook
func UpdateBook(bookData model.Book, id string) model.Book {
	i, book := GetBook(id)

	book.Title = bookData.Title
	book.Author = bookData.Author
	book.IsRead = bookData.IsRead

	storage[i] = book

	return book
}

func DeleteBook(bookId string) bool {
	var deleted []model.Book = []model.Book{}
	for _, val := range storage {
		if bookId != val.Id {
			deleted = append(deleted, val)
		}

	}
	storage = deleted
	return true
}
