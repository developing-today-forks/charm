package migration

// Migration0004_AlterReduceNull allows null for more columns.
var Migration0004_AlterReduceNull = Migration{
	Version: 4,
	Name:    "DT: Alter: reduce null",
	SQL: `
CREATE TABLE IF NOT EXISTS new_connection(
	id INTEGER NOT NULL PRIMARY KEY,
	connection_id VARCHAR NOT NULL,
	status varchar NOT NULL,
	charm_id varchar,
	target_id varchar,
	app varchar NOT NULL,
	type varchar NOT NULL,
	name varchar NOT NULL,
	description varchar NOT NULL,
	username varchar NOT NULL,
	password_length INTEGER,
	password_hash varchar,
	password_hash_type varchar,
	public_key varchar NOT NULL,
	interactive varchar NOT NULL,
	pty varchar NOT NULL,
	protocol varchar NOT NULL,
	server_version varchar NOT NULL,
	client_version varchar NOT NULL,
	session_hash varchar NOT NULL,
	permissions_critical_options varchar NOT NULL,
	permissions_extensions varchar NOT NULL,
	admin varchar NOT NULL,
	query varchar NOT NULL,
	host varchar NOT NULL,
	port INTEGER NOT NULL,
	commands varchar NOT NULL,
	comments varchar NOT NULL,
	history varchar NOT NULL,
	remote_addr varchar NOT NULL,
	remote_addr_network varchar NOT NULL,
	opened_at timestamp default current_timestamp,
	closed_at timestamp default current_timestamp,
	created_at timestamp default current_timestamp,
	updated_at timestamp default current_timestamp,
	deleted_at timestamp
);
INSERT INTO new_connection
SELECT * FROM connection;
DROP TABLE connection;
ALTER TABLE new_connection RENAME TO connection;
`,
}
