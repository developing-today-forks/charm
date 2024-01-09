//go:build libsql
// +build libsql

package migration

// Migration0006_Files
var Migration0006_Files = Migration{
	Version: 6,
	Name:    "DT: create files table",
	SQL: `
CREATE TABLE IF NOT EXISTS files(
	hash varchar NOT NULL,
	hash_type varchar NOT NULL,
	namespace varchar NOT NULL,
	status varchar NOT NULL,
	name varchar NOT NULL,
	description varchar NOT NULL,
	size int NOT NULL,
	path varchar NOT NULL,
	content_type varchar NOT NULL,
	charm_id varchar,
	connection_id varchar,
	public_key varchar,
	text varchar,
	text_type varchar,
	visibility varchar NOT NULL,
	created_at timestamp default current_timestamp,
	update_at timestamp default current_timestamp,
	delete_at timestamp
) RANDOM ROWID;

CREATE TABLE IF NOT EXISTS tags(
	tag varchar NOT NULL,
	tag_type varchar NOT NULL,
	namespace varchar NOT NULL,
	tag_hash varchar NOT NULL,
	tag_hash_type varchar NOT NULL,
	status varchar NOT NULL,
	hash varchar NOT NULL,
	hash_type varchar NOT NULL,
	name varchar NOT NULL,
	description varchar NOT NULL,
	path varchar NOT NULL,
	content_type varchar NOT NULL,
	charm_id varchar,
	connection_id varchar,
	public_key varchar NOT NULL,
	sponsor varchar NOT NULL,
	visibility varchar NOT NULL,
	created_at timestamp default current_timestamp,
	update_at timestamp default current_timestamp,
	delete_at timestamp
) RANDOM ROWID;

CREATE TABLE IF NOT EXISTS hash_activity(
	hash varchar NOT NULL,
	hash_type varchar NOT NULL,
	namespace varchar NOT NULL,
	private_key_conn int NOT NULL,
	password_conn int NOT NULL,
	interactive_conn int NOT NULL,
	room_conn int NOT NULL,
	mutable_conn int NOT NULL,
	get_key int NOT NULL,
	get_value int NOT NULL,
	get_tag int NOT NULL,
	get_meta int NOT NULL,
	get_error int NOT NULL,
	put_key int NOT NULL,
	put_value int NOT NULL,
	put_tag int NOT NULL,
	put_meta int NOT NULL,
	put_error int NOT NULL,
	put_delete int NOT NULL,
	delete_error int NOT NULL,
	admin_conn int NOT NULL,
	name varchar NOT NULL,
	description varchar NOT NULL,
	type varchar NOT NULL,
	path varchar NOT NULL,
	started_at timestamp default current_timestamp,
	ended_at timestamp default current_timestamp,
	created_at timestamp default current_timestamp,
	update_at timestamp default current_timestamp,
	delete_at timestamp
) RANDOM ROWID;

CREATE TABLE IF NOT EXISTS public_key_activity(
	public_key varchar NOT NULL,
	public_key_type varchar NOT NULL,
	namespace varchar NOT NULL,
	private_key_conn int NOT NULL,
	password_conn int NOT NULL,
	interactive_conn int NOT NULL,
	room_conn int NOT NULL,
	mutable_conn int NOT NULL,
	get_key int NOT NULL,
	get_value int NOT NULL,
	get_tag int NOT NULL,
	get_meta int NOT NULL,
	get_error int NOT NULL,
	put_key int NOT NULL,
	put_value int NOT NULL,
	put_tag int NOT NULL,
	put_meta int NOT NULL,
	put_error int NOT NULL,
	put_delete int NOT NULL,
	delete_error int NOT NULL,
	admin_conn int NOT NULL,
	name varchar NOT NULL,
	description varchar NOT NULL,
	type varchar NOT NULL,
	path varchar NOT NULL,
	started_at timestamp default current_timestamp,
	ended_at timestamp default current_timestamp,
	created_at timestamp default current_timestamp,
	update_at timestamp default current_timestamp,
	delete_at timestamp
) RANDOM ROWID;
`,
}
