package database

var INSERT_NEW_BOOK = `
	insert into 
		books (category_id, author_id, book_name) 
	values 
		($1, $2, $3)
	returning 
		book_id, book_name
`

var SELECT_BOOKS = `
	select 
		b.book_name,
		c.category_name,
		a.fullname,
		a.nickname
	from
		books as b
	join categories c using (category_id)
	join authors a using (author_id)
`

var SELECT_BOOK_BY_ID = `
	select 
		b.book_name,
		c.category_name,
		a.fullname,
		a.nickname
	from
		books as b
	join categories c using (category_id)
	join authors a using (author_id)
	where
		book_id = $1
`

var UPDATE_BOOK_BY_ID = `
	update 
		books
	set 
    	category_id = coalesce($2, category_id),
    	author_id = coalesce($3, author_id),
    	book_name = coalesce($4, book_name)
    where
    	book_id = $1	
	returning
		book_id, book_name
`

var DELETE_BOOK_BY_ID = `
	delete from 
		books
	where
		book_id = $1
	returning 
		*
`