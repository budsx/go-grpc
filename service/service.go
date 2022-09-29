package service

import (
	"grpc-unary/model"
	"grpc-unary/repository"
)

func AddBook(book model.Book) model.Book {
	return repository.AddBook(book)
}

func GetBook(id string) (int, model.Book){
	return repository.GetBook(id)
}

func UpdateBook(book model.Book, id string) model.Book{
	return repository.UpdateBook(book, id)
}

func DeleteBook(id string) bool {
	return repository.DeleteBook(id)
}

