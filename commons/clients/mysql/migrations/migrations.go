package migrations

import (
	"database/sql"
	"enceremony-be/commons/clients/mysql/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/manifoldco/promptui"

	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

type MysqlMigrator interface {
	Migrate(migrationsFolder string) error
}

type Impl struct {
	mysqlConf *config.MysqlConfig
}

func NewMysqlMigrator(mysqlConf *config.MysqlConfig) MysqlMigrator {
	return &Impl{mysqlConf: mysqlConf}
}

var short = "Proceed with caution. You are about to corrupt database if not careful"
var long = "When in doubt don't execute this task!!!'"

func (i *Impl) Migrate(migrationsFolder string) error {
	mysqlConf := i.mysqlConf

	// Confirm
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?multiStatements=true",
		mysqlConf.UserName, mysqlConf.Password, mysqlConf.Host, mysqlConf.Port, mysqlConf.DbName)
	fmt.Println("We are connecting to url => " + connString)
	prompt := promptui.Prompt{
		Label:     short,
		IsConfirm: true,
	}

	_, err := prompt.Run()
	if err != nil {
		fmt.Println("Cool, Good Day!")
		return err
	}

	db, _ := sql.Open("mysql", connString)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsFolder,
		"mysql",
		driver,
	)
	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	// Confirm up or down
	promptSelect := promptui.Select{
		Label: "Select the migration type",
		Items: []string{"up", "down"},
	}

	_, result, err := promptSelect.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return err
	}

	if result == "up" {
		err = m.Up()
		if err != nil {
			fmt.Printf("Up Failed %v\n", err)
			return err
		}
		fmt.Printf("Up Worked")
		return nil
	}

	if result == "down" {
		err = m.Down()
		if err != nil {
			fmt.Printf("Down Failed %v\n", err)
			return err
		}
		fmt.Printf("Down Worked")
		return nil
	}

	fmt.Printf("Something went wrong!!")
	return nil
}
