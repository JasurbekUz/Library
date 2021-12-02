package books

import (
	"log"
	d "github.com/JasurbekUz/Library/utils/database"
)

var book Book
var books []Book
var newBook NewBook
var dltdbook DeletedBook

func InsertNewBook (nBook UpdateAndPostBook) NewBook {

	db := d.DB()

	err := db.QueryRow(
			d.INSERT_NEW_BOOK,
			nBook.CategoryId,
			nBook.AuthorId,
			nBook.BookName,
		).Scan(
			&newBook.BookId,			
			&newBook.BookName,
		)

	if err != nil {
			log.Fatalf("db query error: %v", err)
	}

	return newBook
}

func SelectAllBooks () []Book {

	db := d.DB()

	rows, err := db.Query(d.SELECT_BOOKS)

	if err != nil {
			log.Fatalf("db query error: %v", err)
	}

	for rows.Next() {

		rows.Scan(
			&book.BookName,			
			&book.Category.CategoryName,
			&book.Author.FullName,
			&book.Author.NickName,
		)

		books = append(books, book)
	}

	return books
}

func SelectBookById (id int) (Book, error) {

	db := d.DB()

	err := db.QueryRow(
			d.SELECT_BOOK_BY_ID,
			id,
		).Scan(
			&book.BookName,			
			&book.Category.CategoryName,
			&book.Author.FullName,
			&book.Author.NickName,
		)

	return book, err
}

func UpdateBookById (id int, nBook UpdateAndPostBook) (NewBook, error) {

	db := d.DB()

	err := db.QueryRow(
			d.UPDATE_BOOK_BY_ID,
			id,
			nBook.CategoryId,
			nBook.AuthorId,
			nBook.BookName,
		).Scan(
			&newBook.BookId,			
			&newBook.BookName,
		)

	return newBook, err
}

func DbDeleteBookbyId (id int) (DeletedBook, error) {

	db := d.DB()

	err := db.QueryRow(
			d.DELETE_BOOK_BY_ID,
			id,
		).Scan(
			&dltdbook.BookId,			
			&dltdbook.CategoryId,
			&dltdbook.AuthorId,
			&dltdbook.BookName,
		)

	return dltdbook, err
}

