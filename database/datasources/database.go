package datasources

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func init() {
	dsn := "host=" + os.Getenv("LMS_DB_HOST") +
		" user=" + os.Getenv("LMS_DB_USER") +
		" password=" + os.Getenv("LMS_DB_PASSWORD") +
		" dbname=" + os.Getenv("LMS_DB_NAME") +
		" port=" + os.Getenv("LMS_DB_PORT") +
		" sslmode=" + os.Getenv("LMS_DB_SSLMODE")

	log.Println("DSN:", dsn)

	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Database connection failed: %s", err.Error())
		log.Println("Make sure the database is running and the connection details are correct.")
		return // Database initialization fails but program continues to run
	}

	Database.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io.Writer for log output
		logger.Config{
			SlowThreshold:             time.Second, // Threshold for logging slow queries
			LogLevel:                  logger.Info, // Log level for database operations
			IgnoreRecordNotFoundError: true,        // Do not log error for record not found
			Colorful:                  false,       // Disable colorful logs
		},
	)

	log.Println("Database connection successfully established.")
}
