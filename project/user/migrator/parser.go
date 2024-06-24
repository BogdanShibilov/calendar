package migrator

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseAllMigrations(dir string) error {
	fileNames, err := getSqlFileNames(dir)
	if err != nil {
		return err
	}

	for _, fileName := range fileNames {
		m, _ := parseMigration(dir, fileName)
		migrations = append(migrations, *m)
	}

	return nil
}

func getSqlFileNames(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	fileNames := make([]string, 0)
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasSuffix(name, ".sql") {
			fileNames = append(fileNames, name)
		}
	}

	return fileNames, nil
}

func parseMigration(dir, fileName string) (*migration, error) {
	file, err := baseFs.Open(dir + "/" + fileName)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = file.Close()
	}()

	v, _, found := strings.Cut(fileName, "_")
	if !found {
		return nil, ErrInvalidMigrationName
	}

	var upArea strings.Builder
	var downArea strings.Builder
	scanner := bufio.NewScanner(file)
	isInUpArea := false
	isInDownArea := false
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "-- +goose Up") {
			isInUpArea = true
			isInDownArea = false
			continue
		} else if strings.Contains(line, "-- +goose Down") {
			isInUpArea = false
			isInDownArea = true
			continue
		}
		if isInUpArea {
			upArea.WriteString(line)
		}
		if isInDownArea {
			downArea.WriteString(line)
		}
	}

	versionNumber, _ := strconv.Atoi(v)
	return newMigration(
		versionNumber,
		upArea.String(),
		downArea.String(),
	), nil
}
