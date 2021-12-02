package authors

type Author struct {

	AuthorId int 	`json:"authorId"`
	FullName string `json:"fullName"`
	NickName string `json:"nickName"`
	Age string 		`json:"age"`
}

type PostAndUpdateAuthor struct {

	FullName string		`json:"fullName"`
	NickName string 	`json:"nickName"`
	Birth string 		`json:"birth"`
}