package usermigrate

import (
	"embed"
	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations/*.sql
var fsUserMigrations embed.FS

func Provide() migrate.EmbedFileSystemMigrationSource {
	return migrate.EmbedFileSystemMigrationSource{
		FileSystem: fsUserMigrations,
		Root:       "migrations",
	}
}
