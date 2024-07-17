# Kapacity API

## Powered by Golang + GORM

Using the API we use GORM to apply IAC to the database structure. We apply credentials from the database to administrate from the code side.
Once this is in place we can manage our database using the application. 

# Kapacity DB

## Powered by Golang + GORM
Runs an initial task and periodically runs based on a timer value. This timer is essentially just doing a sanity check to ensure the database is meeting the standards for the code. It also will handle any necessary upgrades to the database using AutoMigrate function of GORM.


### Running Development Mode - Sanity Check (GORM Builds the Database Structure from code)
    `
        go run cmd/db/main.go
    `

Example Output
    ```
        [INFO]:[2024-07-17T03:54:36Z][ Starting DB Sanity Check ]
        [INFO]:[2024-07-17T03:54:36Z][ Completed sanity check, sleeping for: 30s ]
        [INFO]:[2024-07-17T03:55:06Z][ Starting DB Sanity Check ]
        [INFO]:[2024-07-17T03:55:06Z][ Completed sanity check, sleeping for: 30s ]
    ```

### ENV File
    `
        MY_VARIABLE=
        DB_PORT=5432
        DB_HOST=
        DB_NAME=
        DB_USER=
        DB_PASSWORD=
    `