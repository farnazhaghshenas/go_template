package config

import (
	"fmt"
	"os"
)

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return fallback
}

//App Holds Application config
var App = struct {
	Port                     string
	GracefullShutdownTimeSec int
}{
	Port:                     getEnv("PORT", "8080"),
	GracefullShutdownTimeSec: 1,
}

type dbConfig struct {
	Name            string
	Host            string
	Port            int
	Username        string
	Password        string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifeTime int
}
func (db *dbConfig) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", db.Username, db.Password, db.Host, db.Port, db.Name)
}

var Database = func() *dbConfig {
	return &dbConfig{
		Name:            "",
		Host:            "mysql",
		Port:            3306,
		Username:        "",
		Password:        "",
		MaxIdleConns:    10,
		MaxOpenConns:    100,
		ConnMaxLifeTime: 10,
	}
}
