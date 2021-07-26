package main

import (
	"Task2/cmd"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
)

var (
	dbtype       = "mysql"
	dbconnection = "root:mypass@tcp(127.0.0.1:3306)/temp"
	db           *gorm.DB
	err          error
	rootCmd      = &cobra.Command{
		Use:   "Creating Table",
		Short: "Create table using this command",
		Long:  `Creates a table if this command is used`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("No arguments provied. Please provide create or drop arguments")
				return
			}

			if args[0] != "create" && args[0] != "drop" {
				fmt.Println("please provide either \"create\" or \"drop\" args")
				return
			}
			if args[0] == "create" {
				fmt.Println("creating table")
				db.AutoMigrate(&User{})
				return
			}

			if args[0] == "drop" {
				fmt.Println("dropping table")
				db.DropTable(&User{})
				return
			}
		},
	}
)

func main() {
	cmd.Execute(rootCmd)
}
func init() {
	db, err = gorm.Open(dbtype, dbconnection)
	if err != nil {
		log.Fatal("Error in opening database connection: ", err)
	}
}

type User struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Age       int    `db:"age"`
}
