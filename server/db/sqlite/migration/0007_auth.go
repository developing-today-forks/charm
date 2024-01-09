package migration

// Migration0007_Auth
var Migration0007_Auth = Migration{
	Version: 7,
	Name:    "DT: add auth method to connections",
	SQL: `
ALTER TABLE connection ADD COLUMN auth_method varchar;
`,
}
