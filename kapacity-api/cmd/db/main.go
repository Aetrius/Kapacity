package main

import (
	Postgres "kapacity-api/internal/postgres"
)

func main() {

	Postgres.DBSanityCheck()

}
