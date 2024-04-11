//go:build libsql
// +build libsql

package migration

// Migration0005_AlterReduceNull allows null for more columns.
var Migration0005_AlterReduceNull = Migration{
	Version: 5,
	Name:    "DT: Alter: reduce null on keys",
	SQL: `
CREATE TABLE IF NOT EXISTS new_hash_public_key(
	id INTEGER NOT NULL PRIMARY KEY,
charm_id varchar,
connection_id varchar,
	hash varchar NOT NULL,
hash_type varchar NOT NULL,
	public_key varchar NOT NULL,
created_at timestamp default current_timestamp,
update_at timestamp default current_timestamp,
delete_at timestamp
) RANDOM ROWID;

INSERT INTO new_hash_public_key
SELECT * FROM hash_public_key;
DROP TABLE hash_public_key;
ALTER TABLE new_hash_public_key RENAME TO hash_public_key;

CREATE TABLE IF NOT EXISTS new_text_public_key(
	id INTEGER NOT NULL PRIMARY KEY,
	charm_id varchar,
	connection_id varchar,
	text varchar NOT NULL,
	text_type varchar NOT NULL,
	public_key varchar NOT NULL,
	created_at timestamp default current_timestamp,
	update_at timestamp default current_timestamp,
	delete_at timestamp
) RANDOM ROWID;

INSERT INTO new_text_public_key
SELECT * FROM text_public_key;
DROP TABLE text_public_key;
ALTER TABLE new_text_public_key RENAME TO text_public_key;

CREATE TABLE IF NOT EXISTS new_private_key(
	id INTEGER NOT NULL PRIMARY KEY,
	charm_id varchar,
	connection_id varchar,
	type varchar NOT NULL,
	private_key varchar NOT NULL,
	public_key varchar NOT NULL,
	created_at timestamp default current_timestamp,
	update_at timestamp default current_timestamp,
	delete_at timestamp
) RANDOM ROWID;

INSERT INTO new_private_key
SELECT * FROM private_key;
DROP TABLE private_key;
ALTER TABLE new_private_key RENAME TO private_key;
	`,
}
