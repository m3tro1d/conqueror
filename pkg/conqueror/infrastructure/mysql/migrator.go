package mysql

import (
	"context"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

type Migrator interface {
	MigrateUp() error
}

func NewMigrator(ctx context.Context, migrationsDir string, client *sqlx.Conn) Migrator {
	return &migrator{
		ctx:           ctx,
		migrationsDir: migrationsDir,
		client:        client,
	}
}

type migrator struct {
	ctx           context.Context
	migrationsDir string
	client        *sqlx.Conn
}

func (m *migrator) MigrateUp() error {
	migrations, err := m.listMigrations()
	if err != nil {
		return err
	}

	executedMigrations, err := m.getExecutedMigrations()
	if err != nil {
		return err
	}

	for _, migration := range migrations {
		if !executedMigrations[migration.Version] {
			err = m.executeMigration(migration)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *migrator) listMigrations() ([]migrationInfo, error) {
	fileInfos, err := ioutil.ReadDir(m.migrationsDir)
	if err != nil {
		return nil, err
	}

	result := make([]migrationInfo, 0, len(fileInfos))
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			version, err := getMigrationVersion(fileInfo.Name())
			if err != nil {
				return nil, err
			}

			result = append(result, migrationInfo{
				Version:  version,
				FilePath: path.Join(m.migrationsDir, fileInfo.Name()),
			})
		}
	}

	return result, nil
}

func (m *migrator) getExecutedMigrations() (map[int]bool, error) {
	const sqlQuery = `SELECT version
		              FROM migration_versions`

	var versions []int
	err := m.client.SelectContext(m.ctx, &versions, sqlQuery)
	if err != nil {
		return nil, err
	}

	result := make(map[int]bool)
	for _, version := range versions {
		result[version] = true
	}

	return result, nil
}

func (m *migrator) executeMigration(migration migrationInfo) error {
	content, err := getFileContent(migration.FilePath)
	if err != nil {
		return err
	}

	_, err = m.client.ExecContext(m.ctx, content)
	if err != nil {
		return err
	}

	return m.saveMigration(migration.Version)
}

func (m *migrator) saveMigration(version int) error {
	const sqlQuery = `INSERT INTO migration_versions (version)
		              VALUES (?)`

	_, err := m.client.ExecContext(m.ctx, sqlQuery, version)
	return err
}

func getFileContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func getMigrationVersion(filename string) (int, error) {
	version := strings.Split(filename, "_")[0]
	return strconv.Atoi(version)
}
