package postgres

import (
	"fmt"
	output "kapacity-api/internal/prettyconsole"
	"time"
)

var token string

func DBSanityCheck() {

	for {
		output.INFO(fmt.Sprintf("Starting DB Sanity Check"))

		pgConnection := importVars()

		// Forces the auto migration process for the DB
		DB(pgConnection)

		output.INFO(fmt.Sprintf("Completed sanity check, sleeping for: %s", 30*time.Second))

		// Keep goroutine loop healthy + pause
		time.Sleep(30 * time.Second)
	}
}
