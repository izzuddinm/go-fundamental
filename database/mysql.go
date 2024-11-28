package database

var database string

func init() {
	database = "MySql"
}

func GetDatabase() string {
	return database
}
