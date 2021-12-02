package categories

import (
	"log"
	d "github.com/JasurbekUz/Library/utils/database"
)

var categories []Category
var categoy Category

func SelectCategories () []Category {

	db := d.DB()

	rows, err := db.Query(d.SELECT_CATEGORIES)

	if err != nil {
		log.Fatalf("db query error: %v", err)
	}

	for rows.Next() {

		rows.Scan(
			&categoy.CategoryId,
			&categoy.Name,
		)

		categories = append(categories, categoy)
	}

	return categories
}