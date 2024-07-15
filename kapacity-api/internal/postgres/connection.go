package postgres

// Incidents structure for output
type PostgresConnection struct {
	Host          string
	Port          int
	User          string
	Password      string
	DBName        string
	EnvironmentID string
}
