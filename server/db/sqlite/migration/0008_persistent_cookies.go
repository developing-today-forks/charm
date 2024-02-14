package migration

// Migration0008_Persistent_Cookies is a db migration script.
// Adds columns to the connection table to store cookies.
// column: cookie_id varchar
// column: cookie_domain varchar
// column: cookie_url varchar
// column: cookie_type varchar
// column: cookie_name varchar
// column: cookie_status varchar
// column: cookie_value varchar
// column: cookie_secret_hash varchar
// column: cookie_secret_hash_type varchar
// column: cookie_secret_type varchar
// column: cookie_secret_payload varchar
// column: cookie_secret_expires_at timestamp
// column: cookie_created_at timestamp
// column: cookie_expires_at timestamp
// column: cookie_context varchar
var Migration0008_Persistent_Cookies = Migration{
	Version: 8,
	Name:    "DT: persistent cookies",
	SQL: `
	ALTER TABLE connection ADD COLUMN cookie_id varchar;
	ALTER TABLE connection ADD COLUMN cookie_domain varchar;
	ALTER TABLE connection ADD COLUMN cookie_url varchar;
	ALTER TABLE connection ADD COLUMN cookie_type varchar;
	ALTER TABLE connection ADD COLUMN cookie_name varchar;
	ALTER TABLE connection ADD COLUMN cookie_status varchar;
	ALTER TABLE connection ADD COLUMN cookie_value varchar;
	ALTER TABLE connection ADD COLUMN cookie_secret_hash varchar;
	ALTER TABLE connection ADD COLUMN cookie_secret_hash_type varchar;
	ALTER TABLE connection ADD COLUMN cookie_secret_type varchar;
	ALTER TABLE connection ADD COLUMN cookie_secret_payload varchar;
	ALTER TABLE connection ADD COLUMN cookie_secret_expires_at timestamp;
	ALTER TABLE connection ADD COLUMN cookie_created_at timestamp;
	ALTER TABLE connection ADD COLUMN cookie_expires_at timestamp;
	ALTER TABLE connection ADD COLUMN cookie_context varchar;
	`,
}
