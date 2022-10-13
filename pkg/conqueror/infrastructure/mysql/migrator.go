package mysql

import (
	"context"
	"io/ioutil"
	"os"
	"path"

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
	migrationFilePaths, err := m.listMigrationFilePaths()
	if err != nil {
		return err
	}

	for _, migrationFilePath := range migrationFilePaths {
		err = m.executeMigration(migrationFilePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *migrator) listMigrationFilePaths() ([]string, error) {
	fileInfos, err := ioutil.ReadDir(m.migrationsDir)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0, len(fileInfos))
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			result = append(result, path.Join(m.migrationsDir, fileInfo.Name()))
		}
	}

	return result, nil
}

func (m *migrator) executeMigration(migrationFilePath string) error {
	content, err := getFileContent(migrationFilePath)
	if err != nil {
		return err
	}

	_, err = m.client.ExecContext(m.ctx, content)
	if err != nil {
		return err
	}

	// TODO: save version

	return nil
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
