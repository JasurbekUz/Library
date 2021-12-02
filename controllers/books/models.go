package books

type Category struct {
	CategoryName string `json:"categoryName"`
}

type Author struct {
	FullName string 	`json:"fullName"`
	NickName string 	`json:"nickName"`
}

type Book struct{
	BookName string		`json:"bookName"`
	Category Category	`json:"category"`
	Author Author		`json:"author"`
}

type UpdateAndPostBook struct {
	CategoryId int 		`json:"categoryId"`
	AuthorId int 		`json:"authorId"`
	BookName string 	`json:"bookName"`
}

type NewBook struct {
	BookId int 			`json:"bookId"`
	BookName string 	`json:"bookName"`
}

type DeletedBook struct {
	BookId int 			`json:"bookId"`
	CategoryId int 		`json:"categoryId"`
	AuthorId int 		`json:"authorId"`
	BookName string 	`json:"bookName"`
}