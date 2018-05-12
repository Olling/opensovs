CREATE TABLE recipes (
	id bigserial PRIMARY KEY,
	title VARCHAR(25),
	added VARCHAR(25),
	blog_id int,
	instructions_id int
);

CREATE TABLE blog (
	id bigserial PRIMARY KEY
);

CREATE TABLE instructions (
	id bigserial PRIMARY KEY
);