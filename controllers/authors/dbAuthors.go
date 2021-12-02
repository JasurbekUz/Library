package authors

import (
	"log"
	d "github.com/JasurbekUz/Library/utils/database"
)

var authors []Author
var author Author


func InsertNewAuthor (nAuthor PostAndUpdateAuthor) Author {

	db := d.DB()

	err := db.QueryRow(
			d.INSERT_NEW_AUTHOR,
			nAuthor.FullName,
			nAuthor.NickName,
			nAuthor.Birth,
		).Scan(
			&author.AuthorId,
			&author.FullName,
			&author.NickName,
			&author.Age,
		)

	if err != nil {
			log.Fatalf("db query error: %v", err)
	}

	return author
}


func SelectAuthors () []Author {

	db := d.DB()

	rows, err := db.Query(d.SELECT_AUTHORS)

	if err != nil {
		log.Fatalf("db query error: %v", err)
	}

	for rows.Next() {

		rows.Scan(
			&author.AuthorId,
			&author.FullName,
			&author.NickName,
			&author.Age,
		)

		authors = append(authors, author)
	}

	return authors
}

func UpdateAuthorById (id int, nAuthor PostAndUpdateAuthor) (Author, error) {

	db := d.DB()

	err := db.QueryRow(
			d.UPDATE_AUTHOR_BY_ID,
			id,
			nAuthor.FullName,
			nAuthor.NickName,
			nAuthor.Birth,
		).Scan(
			&author.AuthorId,
			&author.FullName,
			&author.NickName,
			&author.Age,
		)

	return author, err
}
