package cmd

import (
	"fmt"
	"ggclass_go/src/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func setupM() *migrate.Migrate {
	db, err := config.Cfg.GetDB().DB()
	if err != nil {
		log.Fatalln("err connect to db")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	dir, _ := os.Getwd()
	path := "file://" + dir + "/src/database/postgres"

	m, err := migrate.NewWithDatabaseInstance(path, "postgres", driver)

	if err != nil {
		fmt.Println("err migrate up 1 ", err)
	}

	return m
}

func migrateUp() *cobra.Command {
	return &cobra.Command{

		Use: "migrate-up",
		Run: func(cmd *cobra.Command, args []string) {

			m := setupM()

			err := m.Up()

			if err != nil {
				fmt.Println("err migrate up 2", err)
			}
		},
	}
}

func migrateDown() *cobra.Command {
	return &cobra.Command{
		Use: "migrate-down",
		Run: func(cmd *cobra.Command, args []string) {

			m := setupM()

			err := m.Steps(-1)

			if err != nil {
				fmt.Println("err migrate up 2", err)
			}
		},
	}
}

func migrateRefresh() *cobra.Command {
	return &cobra.Command{
		Use: "migrate-refresh",
		Run: func(cmd *cobra.Command, args []string) {

			m := setupM()

			err := m.Down()
			err = m.Up()

			if err != nil {
				fmt.Println("err migrate up 2", err)
			}
		},
	}
}
