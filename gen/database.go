package gen

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gopkg.in/gormigrate.v1"
)

// DB ...
type DB struct {
	db *gorm.DB
}

// NewDBFromEnvVars Create database client using DATABASE_URL environment variable
func NewDBFromEnvVars() *DB {
	urlString := os.Getenv("DATABASE_URL")
	if urlString == "" {
		panic(fmt.Errorf("missing DATABASE_URL environment variable"))
	}
	return NewDBWithString(urlString)
}

func TableName(name string) string {
	prefix := os.Getenv("TABLE_NAME_PREFIX")
	if prefix != "" {
		return prefix + "_" + name
	}
	return name
}

// NewDB ...
func NewDB(db *gorm.DB) *DB {
	prefix := os.Getenv("TABLE_NAME_PREFIX")
	if prefix != "" {
		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return prefix + "_" + defaultTableName
		}
	}
	v := DB{db}
	return &v
}

// NewDBWithString creates database instance with database URL string
func NewDBWithString(urlString string) *DB {
	u, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	urlString = getConnectionString(u)

	db, err := gorm.Open(u.Scheme, urlString)
	if err != nil {
		panic(err)
	}

	if urlString == "sqlite3://:memory:" {
		db.DB().SetMaxIdleConns(1)
		db.DB().SetConnMaxLifetime(time.Second * 300)
		db.DB().SetMaxOpenConns(1)
	} else {
		db.DB().SetMaxIdleConns(5)
		db.DB().SetConnMaxLifetime(time.Second * 60)
		db.DB().SetMaxOpenConns(10)
	}
	db.LogMode(os.Getenv("DEBUG") == "true")

	return NewDB(db)
}

func getConnectionString(u *url.URL) string {
	if u.Scheme == "postgres" {
		password, _ := u.User.Password()
		params := u.Query()
		params.Set("host", strings.Split(u.Host, ":")[0])
		params.Set("port", u.Port())
		params.Set("user", u.User.Username())
		params.Set("password", password)
		params.Set("dbname", strings.TrimPrefix(u.Path, "/"))
		return strings.Replace(params.Encode(), "&", " ", -1)
		// return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, u.Port(), u.User.Username(), password, strings.TrimPrefix(u.Path, "/"))
	}
	if u.Scheme != "sqlite3" {
		u.Host = "tcp(" + u.Host + ")"
	}
	if u.Scheme == "mysql" {
		q := u.Query()
		q.Set("parseTime", "true")
		u.RawQuery = q.Encode()
	}
	return strings.Replace(u.String(), u.Scheme+"://", "", 1)
}

// Query ...
func (db *DB) Query() *gorm.DB {
	return db.db
}

// AutoMigrate run basic gorm automigration
func (db *DB) AutoMigrate() error {
	return AutoMigrate(db.db)
}

// Migrate run migrations using automigrate
func (db *DB) Migrate(migrations []*gormigrate.Migration) error {
	options := gormigrate.DefaultOptions
	options.TableName = TableName("migrations")
	return Migrate(db.db, options, migrations)
}

// Close ...
func (db *DB) Close() error {
	return db.db.Close()
}

func (db *DB) Ping() error {
	return db.db.DB().Ping()
}
