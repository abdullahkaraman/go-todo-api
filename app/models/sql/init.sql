CREATE TABLE todo (
    id serial       primary key,
    title           VARCHAR NOT NULL,
	description     TEXT NOT NULL,
	category        VARCHAR NOT NULL,
	progress        VARCHAR NOT NULL,
	status          VARCHAR NOT NULL,
	remainingDay    VARCHAR NOT NULL,
	deadline        VARCHAR NOT NULL,
	createdAt       TIME NOT NULL,
	updatedAt       TIME NOT NULL,
);