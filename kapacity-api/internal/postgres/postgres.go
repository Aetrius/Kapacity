package postgres

import (
	"database/sql"
	"fmt"
	"kapacity-api/internal/kapacity"
	output "kapacity-api/internal/prettyconsole"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConnection PostgresConnection

func Connect() gorm.DB {

	var dbConnectionString = importVars()
	connection := DB(dbConnectionString)

	return connection
}

func importVars() PostgresConnection {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// fmt.Println("Loading environment variables...")

	dbConnection := PostgresConnection{
		User:     GetEnvVariable("DB_USER"),
		Port:     stringToInt(os.Getenv("DB_PORT")),
		Host:     GetEnvVariable("DB_HOST"),
		DBName:   GetEnvVariable("DB_NAME"),
		Password: GetEnvVariable("DB_PASSWORD"),
	}

	return dbConnection
}

func stringToInt(val string) int {
	intValue, err := strconv.Atoi(val)
	if err != nil {

		output.ERROR(fmt.Sprintf("Error converting %s to an integer: %v\n", val, err))
		return 0
	}
	return intValue
}

func DB(connectionInfo PostgresConnection) gorm.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", connectionInfo.Host, connectionInfo.Port, connectionInfo.User, connectionInfo.Password, connectionInfo.DBName)

	// Connect to the database (SQLite in this example)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		output.ERROR(fmt.Sprintf("Failed to connect to database"))
		panic("Failed to connect to database")

	}

	// Set connection pool settings on the database instance.
	sqlDB, err := db.DB()
	if err != nil {
		output.ERROR(fmt.Sprintf("%s", err))
	}

	// DB Connection Pool settings prevent unclosed connections from gorm, too many idles, etc.
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetConnMaxLifetime(2 * time.Minute)

	// Automatically create the table if it doesn't exist
	db.AutoMigrate(&kapacity.Cluster{}, &kapacity.Node{})

	return *db
}

func RefreshMaterializedViews(db gorm.DB, view string) error {

	// Execute the REFRESH MATERIALIZED VIEW command
	// Replace "my_materialized_view" with the name of your materialized view
	result := db.Exec(fmt.Sprintf("REFRESH MATERIALIZED VIEW %s", view))

	if result.Error != nil {
		output.ERROR(fmt.Sprintf("%s", result.Error))
		return result.Error
	}

	output.INFO(fmt.Sprintf("Materialized views refreshed"))

	return nil
}

func RunSQLPing(connectionInfo PostgresConnection) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", connectionInfo.Host, connectionInfo.Port, connectionInfo.User, connectionInfo.Password, connectionInfo.DBName)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping() ///run the command here
	CheckError(err)

	output.INFO(fmt.Sprintf("Connected to DB!"))

}

func CheckError(err error) {
	if err != nil {
		output.ERROR(fmt.Sprintf("%s", err))
		panic(err)
	}
}

func SaveClustersToDB(db gorm.DB, clusters []kapacity.Cluster) {
	for _, cluster := range clusters {
		// Check if the user already exists in the database based on Email
		existingCluster := kapacity.Cluster{}
		result := db.Where("incident_id = ?", cluster.ClusterID).Select("cluster_id").First(&existingCluster)
		if result.Error == gorm.ErrRecordNotFound {
			// Incident does not exist in the database, create a new record
			db.Create(&cluster)
		} else {
			// Incident already exists, update the existing record (using new record data)
			existingCluster = cluster
			// existingIncident.ResolvedBy = incident.ResolvedBy
			db.Save(&existingCluster) // save checks the record and updates it if anything is different
		}
	}
	output.INFO(fmt.Sprintf("Create or update completed for %d clusters", len(clusters)))
}
