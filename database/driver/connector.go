package driver

import (
	"fmt"
	"go-starterkit-project/config"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/**
this variable is a global variable to reuse database connection
*/
var DB *gorm.DB

/**
This function as a factory to connect to MySQL or PostgreSQL
*/
func ConnectDB() *gorm.DB {
	var err error
	var dbType string

	// check whether in the configuration using mysql or postgresql driver
	if config.Config("DB_DRIVER") == "mysql" {
		// Open MySQL connection
		DB, err = gorm.Open(mysql.Open(DsnMySqlDB()))
	} else if config.Config("DB_DRIVER") == "pgsql" {
		// Open PostgreSQL connection
		DB, err = gorm.Open(postgres.Open(DsnPostgreSqlDB()))
	} else {
		// Stop the application if the driver does not match
		panic("Database driver not available")
	}

	// Display an error message if an error occurs in the database connection
	dbType = config.Config("DB_DRIVER")
	if err != nil {
		errMessage := fmt.Sprintf("Failed to connect database %s", dbType)
		panic(errMessage)
	}

	// Call db pooling function
	dbPooling(DB)

	return DB
}

/**
This function is for database pooling
*/
func dbPooling(sqlDb *gorm.DB) error {
	// Get generic database object sql.DB to use its functions
	sqlDB, err := DB.DB()

	if err != nil {
		panic("failed to connect database")
	}

	// Get param config into var
	maxIdleConsConf := config.Config("DB_MAX_IDLE_CONNS")
	maxOpenConsConf := config.Config("DB_MAX_OPEN_CONNS")

	// Convert string to integer
	maxIdleCons, _ := strconv.Atoi(maxIdleConsConf)
	maxOpenCons, _ := strconv.Atoi(maxOpenConsConf)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(maxIdleCons)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(maxOpenCons)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
