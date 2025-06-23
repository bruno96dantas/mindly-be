package settings

import (
	"os"

	"github.com/go-sql-driver/mysql"
)

var MySQLConfig = mysql.NewConfig()

func LoadMySQLConfigs() {
	MySQLConfig.Net = "tcp"
	MySQLConfig.User = os.Getenv("DB_MYSQL_USER")
	MySQLConfig.Passwd = os.Getenv("DB_MYSQL_PASSWORD")
	MySQLConfig.Addr = os.Getenv("DB_MYSQL_HOSTS")
	MySQLConfig.DBName = os.Getenv("DATABASE_NAME")
}
