package variables

import "os"

var Mode string

// Mysql environment
var Mysql struct {
	Host     string
	User     string
	Pass     string
	Database string
}

//Api Keys
var ApiKeys struct {
	OMDB string
}

func Initialization() {
	Mode = os.Getenv("MODE")

	Mysql.Host = os.Getenv("MYSQL_HOST")
	Mysql.User = os.Getenv("MYSQL_USER")
	Mysql.Pass = os.Getenv("MYSQL_PASSWORD")
	Mysql.Database = os.Getenv("MYSQL_DATABASE")

	ApiKeys.OMDB = os.Getenv("OMDB_API_KEY")
}
