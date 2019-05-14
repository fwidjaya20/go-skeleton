package config

import (
	"fmt"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/payfazz/go-apt/pkg/fazzdb"
)

var db *sqlx.DB
var once sync.Once

// Set is a function to append value to base config
func Set(key string, value string) {
	base[key] = value
}

// Get config by key
func Get(key string) string {
	r := os.Getenv(key)
	if r == "" {
		if _, ok := base[key]; !ok {
			return ""
		}
		r = base[key]
	}
	return r
}

// GetDb is a function to get DB instance
func GetDb() *sqlx.DB {
	once.Do(func() {
		var err error
		conn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			Get(DB_HOST),
			Get(DB_PORT),
			Get(DB_USER),
			Get(DB_PASS),
			Get(DB_NAME),
		)

		db, err = sqlx.Connect("postgres", conn)
		if nil != err {
			panic(err)
		}
	})
	return db
}

// GetIfString get config as string
func GetIfString(key string) string {
	var t string
	return getIf(key, t).(string)
}

// GetIfInteger get config as int
func GetIfInteger(key string) int {
	var t int
	return getIf(key, t).(int)
}

// GetIfDuration get config as duration
func GetIfDuration(key string) time.Duration {
	var t time.Duration
	return getIf(key, t).(time.Duration)
}

// GetIfQueryConfig get config as fazzdb.Config
func GetIfQueryConfig(key string) fazzdb.Config {
	var t fazzdb.Config
	return getIf(key, t).(fazzdb.Config)
}

// ForceMigrate get force migrate status depending on FORCE_MIGRATE and current ENV
func ForceMigrate() bool {
	return Get(ENV) != ENV_PRODUCTION && Get(FORCE_MIGRATE) == FORCE_MIGRATE_ON
}

func getIf(key string, p interface{}) interface{} {
	t := reflect.TypeOf(p)

	if _, ok := baseInterface[key]; !ok {
		panic(fmt.Sprintf(`config with key "%s" not found`, key))
	}

	if keyType := reflect.TypeOf(baseInterface[key]); t != keyType {
		panic(fmt.Sprintf(`different type of config with key "%s" got %s expected %s`, key, keyType, t))
	}

	return baseInterface[key]
}
