# Kapacity API

## Powered by Golang + GORM

Using the API we use GORM to apply IAC to the database structure. We apply credentials from the database to administrate from the code side.
Once this is in place we can manage our database using the application. 

# Kapacity DB

## Powered by Golang + GORM
Runs an initial task and periodically runs based on a timer value. This timer is essentially just doing a sanity check to ensure the database is meeting the standards for the code. It also will handle any necessary upgrades to the database using AutoMigrate function of GORM.