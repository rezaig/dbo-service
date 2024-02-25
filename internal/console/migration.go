package console

import (
	"database/sql"
	"fmt"
	"strconv"

	sq "github.com/Masterminds/squirrel"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rezaig/dbo-service/internal/helper"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Long:  `This subcommand is used to migrate database`,
	Run:   processMigration,
}

func init() {
	migrateCmd.PersistentFlags().Int("step", 0, "maximum migration steps")
	migrateCmd.PersistentFlags().String("direction", "up", "migration direction")
	RootCmd.AddCommand(migrateCmd)
}

func processMigration(cmd *cobra.Command, args []string) {
	direction := cmd.Flag("direction").Value.String()
	stepStr := cmd.Flag("step").Value.String()
	step, err := strconv.Atoi(stepStr)
	if err != nil {
		log.WithField("stepStr", stepStr).Fatal("failed to parse step to int: ", err)
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "./db/migration",
	}

	migrate.SetTable("schema_migrations")

	var (
		host     = viper.GetString("mysql.host")
		username = viper.GetString("mysql.username")
		password = viper.GetString("mysql.password")
		database = viper.GetString("mysql.database")
	)

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, host, database)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.WithField("DSN", connStr).Fatal("failed to connect database: ", err)
	}

	var n int
	if direction == "down" {
		n, err = migrate.ExecMax(db, "mysql", migrations, migrate.Down, step)
	} else {
		n, err = migrate.ExecMax(db, "mysql", migrations, migrate.Up, step)
	}
	if err != nil {
		log.WithFields(log.Fields{
			"db":         db,
			"migrations": helper.Dump(migrations),
			"direction":  direction}).
			Fatal("failed to migrate database: ", err)
	}

	log.Infof("applied %d migrations!\n", n)

	// Seed product
	_, _ = sq.Insert("product").
		Columns("id", "name", "description").
		Values(1, "Rucika", "Rucika lorem ipsum").
		RunWith(db).Exec()
	// Seed customer
	_, _ = sq.Insert("customer").
		Columns("id", "name", "email").
		Values(1, "Reza Indra", "mail.rezaindra@gmail.com").
		RunWith(db).Exec()
	// Seed account
	hashedPassword, _ := helper.HashPassword("password123")
	_, _ = sq.Insert("account").
		Columns("id", "username", "password").
		Values(1, "reza", hashedPassword).
		RunWith(db).Exec()
	// Seed order
	_, _ = sq.Insert("`order`").
		Columns("customer_id", "product_id", "quantity", "shipping_address").
		Values(1, 1, 2, "Jl. Merdeka 17").
		RunWith(db).Exec()
}
