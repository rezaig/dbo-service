package config

import (
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.AddConfigPath("./../../..")
	viper.SetConfigName("config")
	viper.SetEnvPrefix("svc")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Warningf("config not found, error: %v", err)
	}
}

func Port() string {
	return viper.GetString("port")
}

func MySQLHost() string {
	return viper.GetString("mysql.host")
}

func MySQLDBName() string {
	return viper.GetString("mysql.database")
}

func MySQLUsername() string {
	return viper.GetString("mysql.username")
}

func MySQLPassword() string {
	return viper.GetString("mysql.password")
}

func MySQLMaxIdleConns() int64 {
	if viper.GetInt64("mysql.max_idle_conns") <= 0 {
		return 10
	}
	return viper.GetInt64("mysql.max_idle_conns")
}

func MySQLMaxOpenConns() int64 {
	if viper.GetInt64("mysql.max_open_conns") <= 0 {
		return 10
	}
	return viper.GetInt64("mysql.max_open_conns")
}

func MySQLConnMaxLifetime() time.Duration {
	return parseDuration(viper.GetString("mysql.conn_max_lifetime"), 3*time.Second)
}

func MySQLConnMaxIdletime() time.Duration {
	return parseDuration(viper.GetString("mysql.conn_max_idletime"), 3*time.Second)
}

func JWTSigningKey() string {
	return viper.GetString("jwt.signing_key")
}

func JWTExp() time.Duration {
	return parseDuration(viper.GetString("jwt.exp"), 24*time.Hour)
}

func parseDuration(in string, defaultDuration time.Duration) time.Duration {
	dur, err := time.ParseDuration(in)
	if err != nil {
		return defaultDuration
	}
	return dur
}
