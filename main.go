package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

type Account_value struct {
	Alias_name    string "json:alias_name"
	Account_name  string "json:account_name"
	Account_email string "json:account_email"
}

func main() {

	app := cli.NewApp()
	app.Name = "Change git account"
	app.Usage = "chgitacc"
	app.Version = "1.0.0"
	app.Action = func(c *cli.Context) error {
		if c.Bool("private") {
			change_account("private")
		} else {
			var account_asset string
			fmt.Print("Select Account : ")
			fmt.Scanln(&account_asset)
			change_account(account_asset)
		}
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "private, p",
			Usage: "Change to  private git account",
		},
	}
	app.Run(os.Args)
}

func change_account(account_alias string) {

	var account_name, account_email string

	account_json, err := ioutil.ReadFile("./account.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var account_data []Account_value
	if err := json.Unmarshal(account_json, &account_data); err != nil {
		log.Fatal(err)
	}

	for _, data := range account_data {
		if account_alias == data.Alias_name {
			account_email = data.Account_email
			account_name = data.Account_name
			break
		}
	}

	if len(account_email) == 0 || len(account_name) == 0 {
		log.Fatal("account not exist")
	}
	fmt.Println("Set account name", account_name, " , Set account email", account_email)
	_, err = exec.Command("git", "config", "--global", "user.name", account_name).Output()
	if err != nil {
		log.Fatal("change user name", err)
	}
	_, err = exec.Command("git", "config", "--global", "user.email", account_email).Output()
	if err != nil {
		log.Fatal("change user email ", err)
	}
}
