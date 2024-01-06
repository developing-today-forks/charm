//go:build sqlite
// +build sqlite

package migration

// Migration0003_DT_Passwords_And_Keys is a db migration script.
// This migration adds the following tables:
// - connection
// - hash_public_key
// - text_public_key
// - private_key
// The tables are used to store passwords and keys.
// One can get the public key for a given hash or text.
// Certain public keys can have their private key stored and shared.
// Connections have their details stored in the connection table.
var Migration0003_DT_Passwords_And_Keys = Migration{
	Version: 3,
	Name:    "DT: passwords & keys",
	SQL: `
CREATE TABLE IF NOT EXISTS connection(
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
	password_length INTEGER NOT NULL,
	password_hash varchar NOT NULL,
	password_hash_type varchar NOT NULL,
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

CREATE TABLE IF NOT EXISTS hash_public_key(
  id INTEGER NOT NULL PRIMARY KEY,
	charm_id varchar NOT NULL,
	connection_id varchar NOT NULL,
  hash varchar NOT NULL,
	hash_type varchar NOT NULL,
  public_key varchar NOT NULL,
	created_at timestamp default current_timestamp,
	update_at timestamp default current_timestamp,
	delete_at timestamp
);

CREATE TABLE IF NOT EXISTS text_public_key(
	id INTEGER NOT NULL PRIMARY KEY,
	charm_id varchar NOT NULL,
	connection_id varchar NOT NULL,
	text varchar NOT NULL,
	text_type varchar NOT NULL,
	public_key varchar NOT NULL,
	created_at timestamp default current_timestamp,
	update_at timestamp default current_timestamp,
	delete_at timestamp
);

CREATE TABLE IF NOT EXISTS private_key(
	id INTEGER NOT NULL PRIMARY KEY,
	charm_id varchar NOT NULL,
	connection_id varchar NOT NULL,
	type varchar NOT NULL,
	private_key varchar NOT NULL,
	public_key varchar NOT NULL,
	created_at timestamp default current_timestamp,
	update_at timestamp default current_timestamp,
	delete_at timestamp
);
`,
}
