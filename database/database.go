package database

import (
	"errors"
	"oriontask-v2/dharmas"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(appName string) (*gorm.DB, error) {
	databasePath, err := databaseFilePath(appName)
	if err != nil {
		return nil, err
	}

	if err := ensureDatabaseLocation(databasePath); err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&dharmas.Dharmas{}); err != nil {
		return nil, err
	}

	return db, nil
}

func databaseFilePath(appName string) (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, appName, "database.db"), nil
}

func ensureDatabaseLocation(databasePath string) error {
	if err := os.MkdirAll(filepath.Dir(databasePath), 0o755); err != nil {
		return err
	}

	legacyDatabasePath := "database.db"
	if samePath(databasePath, legacyDatabasePath) {
		return nil
	}

	if _, err := os.Stat(databasePath); err == nil {
		return nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if _, err := os.Stat(legacyDatabasePath); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	return copyFile(legacyDatabasePath, databasePath)
}

func samePath(left string, right string) bool {
	leftAbs, leftErr := filepath.Abs(left)
	rightAbs, rightErr := filepath.Abs(right)
	if leftErr != nil || rightErr != nil {
		return left == right
	}

	return leftAbs == rightAbs
}

func copyFile(source string, destination string) error {
	content, err := os.ReadFile(source)
	if err != nil {
		return err
	}

	return os.WriteFile(destination, content, 0o600)
}
