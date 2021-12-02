package database

var SELECT_AUTHORS = `
	select
		author_id,
		fullname,
		nickname,
		extract(year from now()) - extract(year from birth)
	from
		authors
`

var INSERT_NEW_AUTHOR = `
	insert into 
		authors (fullname, nickname, birth) 
	values 
		($1, $2, $3)
	returning 
		author_id,
		fullname,
		nickname,
		extract(year from now()) - extract(year from birth)
`

var UPDATE_AUTHOR_BY_ID = `
	update 
		authors
	set 
    	fullname = coalesce($2, fullname),
    	nickname = coalesce($3, nickname),
    	birth = coalesce($4, birth)
    where
    	author_id = $1	
	returning
		author_id, 
		fullname,
		nickname,
		extract(year from now()) - extract(year from birth)
`