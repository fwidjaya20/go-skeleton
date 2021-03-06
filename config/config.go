package config

import (
	"time"

	"github.com/payfazz/go-apt/pkg/fazzdb"
)

const (
	// ENV_DEVELOPMENT development
	ENV_DEVELOPMENT = "development"
	// ENV_STAGING staging
	ENV_STAGING = "staging"
	// ENV_PRODUCTION production
	ENV_PRODUCTION = "production"
)

const (
	// FORCE_MIGRATE_OFF off
	FORCE_MIGRATE_OFF = "off"
	// FORCE_MIGRATE_ON on
	FORCE_MIGRATE_ON = "on"
)

var base = map[string]string{
	// ENV ENV_PRODUCTION | ENV_STAGING | ENV_DEVELOPMENT
	ENV: ENV_DEVELOPMENT,
	// PORT default port for server
	PORT: ":8080",
	// DB_HOST host for DB
	DB_HOST: "localhost",
	// DB_PORT port for DB
	DB_PORT: "5432",
	// DB_USER user for DB
	DB_USER: "postgres",
	// DB_PASS pass for DB
	DB_PASS: "postgres",
	// DB_NAME name for DB
	DB_NAME: "example",
	// REDIS_HOST host for redis, format: {host}:{port}
	REDIS_HOST: "",
	// REDIS_PASSWORD pass for redis
	REDIS_PASSWORD: "",
	// CLIENT_ID client id
	CLIENT_ID: "",
	// AUTHFAZZ_URL url for authfazz
	AUTHFAZZ_URL: "",
	// AUTHFAZZ_SECRET secret
	AUTHFAZZ_SECRET: "",
	// FORCE_MIGRATE default: on
	FORCE_MIGRATE: "on",
}

var baseInterface = map[string]interface{}{
	// I_QUERY_CONFIG conf for DB limit, offset and lock
	I_QUERY_CONFIG: fazzdb.Config{
		Limit:           MAX_PER_PAGE,
		Offset:          0,
		Lock:            fazzdb.LO_NONE,
		DevelopmentMode: Get(ENV) != ENV_PRODUCTION,
	},
	// I_READ_TIMEOUT read timeout
	I_READ_TIMEOUT: 5 * time.Minute,
	// I_WRITE_TIMEOUT write timeout
	I_WRITE_TIMEOUT: 5 * time.Minute,
	// I_IDLE_TIMEOUT idle timeout
	I_IDLE_TIMEOUT: 30 * time.Second,
	// I_WAIT_SHUTDOWN graceful shutdown delay
	I_WAIT_SHUTDOWN: 5 * time.Second,
	// I_CREDENTIAL_EXPIRE credential expired time
	I_CREDENTIAL_EXPIRE: 24 * 25 * time.Hour,
	// I_REDIS_SESSION_EXPIRE const for redis expired session
	I_REDIS_SESSION_EXPIRE: 24 * time.Hour,
	// I_THROTTLE_LIMIT limit per hit for throttling
	I_THROTTLE_LIMIT: 40,
	// I_THROTTLE_DURATION duration for throttle TTL
	I_THROTTLE_DURATION: time.Minute,
	// I_THROTTLE_TYPE type checker for throttle
	I_THROTTLE_TYPE: 0,
}

const (
	// ENV ENV
	ENV = "ENV"
	// PORT PORT
	PORT = "PORT"
	// DB_HOST DB_HOST
	DB_HOST = "DB_HOST"
	// DB_PORT DB_PORT
	DB_PORT = "DB_PORT"
	// DB_USER DB_USER
	DB_USER = "DB_USER"
	// DB_PASS DB_PASS
	DB_PASS = "DB_PASS"
	// DB_NAME DB_NAME
	DB_NAME = "DB_NAME"
	// REDIS_HOST REDIS_HOST
	REDIS_HOST = "REDIS_HOST"
	// REDIS_PASSWORD REDIS_PASSWORD
	REDIS_PASSWORD = "REDIS_PASSWORD"
	// AUTHFAZZ_URL AUTHFAZZ_URL
	AUTHFAZZ_URL = "AUTHFAZZ_URL"
	// AUTHFAZZ_SECRET AUTHFAZZ_SECRET
	AUTHFAZZ_SECRET = "AUTHFAZZ_SECRET"
	// CLIENT_ID CLIENT_ID
	CLIENT_ID = "CLIENT_ID"
	// FORCE_MIGRATE FORCE_MIGRATE
	FORCE_MIGRATE = "FORCE_MIGRATE"
)

const (
	// I_QUERY_CONFIG QUERY_CONFIG
	I_QUERY_CONFIG = "QUERY_CONFIG"
	// I_READ_TIMEOUT READ_TIMEOUT
	I_READ_TIMEOUT = "READ_TIMEOUT"
	// I_WRITE_TIMEOUT WRITE_TIMEOUT
	I_WRITE_TIMEOUT = "WRITE_TIMEOUT"
	// I_IDLE_TIMEOUT IDLE_TIMEOUT
	I_IDLE_TIMEOUT = "IDLE_TIMEOUT"
	// I_WAIT_SHUTDOWN WAIT_SHUTDOWN
	I_WAIT_SHUTDOWN = "WAIT_SHUTDOWN"
	// I_CREDENTIAL_EXPIRE CREDENTIAL_EXPIRE
	I_CREDENTIAL_EXPIRE = "CREDENTIAL_EXPIRE"
	// I_REDIS_SESSION_EXPIRE REDIS_SESSION_EXPIRE
	I_REDIS_SESSION_EXPIRE = "REDIS_SESSION_EXPIRE"
	// I_THROTTLE_LIMIT THROTTLE_LIMIT
	I_THROTTLE_LIMIT = "THROTTLE_LIMIT"
	// I_THROTTLE_DURATION THROTTLE_DURATION
	I_THROTTLE_DURATION = "THROTTLE_DURATION"
	// I_THROTTLE_TYPE THROTTLE_TYPE
	I_THROTTLE_TYPE = "THROTTLE_TYPE"
)

const (
	// MAX_PER_PAGE 20
	MAX_PER_PAGE = 20
	// DEFAULT_LIMIT_QUERY 1
	DEFAULT_LIMIT_QUERY = 1
)
