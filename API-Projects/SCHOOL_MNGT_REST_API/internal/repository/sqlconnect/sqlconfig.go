package sqlconnect

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // Importing the MySQL driver
)

// ConnectDb establishes a connection to the MariaDB database with the given name.
func ConnectDb(dbname string) (*sql.DB, error) {

	// Load environment variables from .env file
	// if err := godotenv.Load(); err != nil {
	// 	return nil, err
	// }

	// Fetch database connection parameters from environment variables
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("HOST")
	dbport := os.Getenv("DB_PORT")

	// Data Source Name (DSN) format: username:password@protocol(address)/dbname
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, dbport, dbname)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MariaDB")
	return db, nil
}
