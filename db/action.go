package db

const createTable string = `
	CREATE TABLE IF NOT EXISTS files (
		id INTEGER NOT NULL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
	    location VARCHAR(255) NOT NULL,
	    time DATETIME NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS logs (
	    id INTEGER NOT NULL PRIMARY KEY,
	    title VARCHAR(255) NOT NULL,
	    description TEXT,
	    time DATETIME NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS structures (
	    id INTEGER NOT NULL PRIMARY KEY,
	    name VARCHAR(255) NOT NULL,
	    filename_as VARCHAR(255) NOT NULL,
	    time DATETIME NOT NULL
	);

	CREATE TABLE IF NOT EXISTS file_columns (
	    id INTEGER NOT NULL PRIMARY KEY,
	    structure_id INTEGER NOT NULL,
	    key VARCHAR(255) NOT NULL,
	    is_unique boolean default false,
	    time DATETIME NOT NULL,
	    constraint fk_column_structure
		foreign key (structure_id) 
		REFERENCES structures (id)
	);

	CREATE TABLE IF NOT EXISTS file_structures (
	    id INTEGER NOT NULL PRIMARY KEY,
	    structure_id INTEGER NOT NULL,
	    location VARCHAR(512) NOT NULL,
	    completed boolean DEFAULT FALSE,
	    time DATETIME NOT NULL,
	    constraint fk_column_structure
		foreign key (structure_id) 
		REFERENCES structures (id)
	)
`
