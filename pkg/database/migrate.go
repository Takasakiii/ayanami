package database

import "fmt"

func (d *GormDatabase) Migrate(tables ...interface{}) error {
	if d.connection == nil {
		return fmt.Errorf("database migration connection is nil")
	}

	err := d.connection.AutoMigrate(tables...)
	if err != nil {
		return fmt.Errorf("database migration error: %v", err)
	}
	return nil
}
