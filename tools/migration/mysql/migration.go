package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	mysqlUtils "github.com/nguyenhoai890/wager_service/pkg/mysql"
)
type DBConfig struct {
	DBName string
	Username string
	Password string


}
type Migration struct {
	*migrate.Migrate
}

func Init(uri string, filePath string) (m *Migration, err error) {
	if err != nil {
		return
	}
	driverName := "mysql"
	db, err := sql.Open(driverName, uri)
	if err != nil {
		return
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(err)
	}
	migrateDB, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s", filePath),
		driverName,
		driver,
	)
	if migrateDB != nil {
		v := Migration{
			Migrate: migrateDB,
		}
		m = &v
	}
	return
}

func InitDB(uri mysqlUtils.Uri) error {
	driverName := "mysql"
	db, err := sql.Open(driverName, uri.GetServerUri())
	if err != nil {
		return err
	}
	defer func() {
		db.Close()
	}()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", uri.DatabaseName))//dbName))
	return err
}

func Up(uri mysqlUtils.Uri, filePath string) error {
	if err := InitDB(uri); err != nil {
		return err
	}
	m, err := Init(uri.GetFullUri(), filePath)
	if err != nil {
		return nil
	}
	return m.Up()
}
func Drop(uri mysqlUtils.Uri, filePath string) error {
	if err := InitDB(uri); err != nil {
		return err
	}
	m, err := Init(uri.GetFullUri(), filePath)
	if err != nil {
		return nil
	}
	return m.Drop()
}