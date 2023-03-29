package schema_test

const parallelCreate = 8

const DatabaseSchema = `
	CREATE TABLE accounts (
		user_id serial PRIMARY KEY
	);
`
