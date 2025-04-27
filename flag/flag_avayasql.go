package flag

import (
	"server/global"
	"server/model/database"
)

// SQL table migration
func avayaSQL() error {
	return global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&database.Employee{},
		&database.SnowItems{},
		&database.SiebelItems{},
	)
}
