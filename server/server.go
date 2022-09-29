package server

import (
	"context"
	"errors"
	pb "grpc-unary/book"
	"grpc-unary/model"
	"grpc-unary/service"
)

// server gRPC
type Server struct{
	pb.UnimplementedBookServiceServer
}

// mapping  book model to pb
func mapToBookPb(book model.Book) *pb.Book {
	return &pb.Book{
		Id:     book.Id,
		Title:  book.Title,
		Author: book.Author,
		IsRead: book.IsRead,
	}
}

// implement all method BookServiceServer
func (*Server) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	var bookID string = req.GetId()
	_, res := service.GetBook(bookID)
	if res.Id != bookID {
		return &pb.GetBookResponse{Status: false, Data: nil}, errors.New("Not Found")
	}

	var bookData *pb.Book = &pb.Book{
		Id:     res.Id,
		Title:  res.Title,
		Author: res.Author,
		IsRead: res.IsRead,
	}

	return &pb.GetBookResponse{
		Status: true,
		Data:   bookData,
	}, nil
}

func AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	var bookRequest *pb.Book = req.GetData()

	bookData := model.Book{
		Id:     bookRequest.GetId(),
		Title:  bookRequest.GetTitle(),
		Author: bookRequest.GetAuthor(),
		IsRead: bookRequest.GetIsRead(),
	}

	var insertedBook model.Book = service.AddBook(bookData)

	return &pb.AddBookResponse{
		Status: true,
		Data:   mapToBookPb(insertedBook),
	}, nil
}

func UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	var bookRequest *pb.Book = req.GetBook()

	var bookData model.Book = model.Book{
		Id:     bookRequest.GetId(),
		Title:  bookRequest.GetTitle(),
		Author: bookRequest.GetAuthor(),
		IsRead: bookRequest.GetIsRead(),
	}

	var updatedBook model.Book = service.UpdateBook(bookData, bookData.Id)

	return &pb.UpdateBookResponse{
		Status: true,
		Data:   mapToBookPb(updatedBook),
	}, nil
}

func DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error){
	var bookID string = req.GetBookId()

	var res bool = service.DeleteBook(bookID)

	return &pb.DeleteBookResponse{
		Status: res,
	}, nil
}


