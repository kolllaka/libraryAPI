CREATE TABLE books (
	id BINARY(36) PRIMARY KEY,
	name_author VARCHAR(60) NOT NULL,
	name_book VARCHAR(255) NOT NULL,
	genre VARCHAR(255) NOT NULL,
	publication_year INT NOT NULL,
	num_pages INT NOT NULL,
	bbk VARCHAR(255) NOT NULL,
	description_book VARCHAR(255),
	add_date TIMESTAMP,
	is_here BOOLEAN
);