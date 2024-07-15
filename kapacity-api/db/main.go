package main

import (
	Postgres "kapacity-api/internal/postgres"
)

//var dbConnection postgres.PostgresConnection

func main() {

	Postgres.DBSanityCheck()

}
