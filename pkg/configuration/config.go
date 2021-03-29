package configuration

import (
	"os"
)

var (
	MysqlServerUri       string
	MysqlServerUriParams string
	MysqlDB              string
	MysqlTestDB          string
	MigratePath          string
)

func init() {
	MigratePath = os.Getenv("MIGRATE_PATH")
	MysqlServerUri = os.Getenv("MYSQL_URI")
	MysqlServerUriParams = os.Getenv("MYSQL_URI_PARAMS")
	MysqlDB = os.Getenv("MYSQL_DB")
	MysqlTestDB = os.Getenv("MYSQL_TEST_DB")
}
