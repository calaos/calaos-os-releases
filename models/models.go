package models

import (
	"github.com/calaos/calaos-os-releases/config"
	logger "github.com/calaos/calaos-os-releases/log"

	"github.com/sirupsen/logrus"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var (
	logging *logrus.Entry

	db *gorm.DB
)

func init() {
	logging = logger.NewLogger("database")
}

// Init models
func Init() (err error) {
	logging.Infof("Opening database %s", config.Config.String("database.sqlite"))
	db, err = gorm.Open(sqlite.Open(config.Config.String("database.sqlite")), &gorm.Config{
		//Logger: logger.NewGorm(), //TODO: our logrus impl does not work
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	if err != nil {
		return
	}

	migrateDb()

	return
}

// Shutdown models
func Shutdown() {

}

func migrateDb() {
	//Migrate all tables
	db.AutoMigrate(
		&Image{},
	)

	logging.Infof("Migration did run successfully")
}
