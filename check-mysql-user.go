package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	"os"
)

type options struct {
	Host string `short:"H" long:"host" default:"localhost" description:"Hostname"`
	Port string `short:"p" long:"port" default:"3306" description:"Port"`
	User string `short:"u" long:"user" default:"root" description:"Username"`
	Pass string `short:"P" long:"password" default:"" description:"Password"`
	AccountName string `short:"a" long:"account-name" arg:"String" required:"true" description:"account user name"`
	AccountHost string `short:"n" long:"account-host" arg:"String" required:"true" description:"account user host"`
}

func main() {
	ckr := checkUser()
	ckr.Name = "MySQL User"
	ckr.Exit()
}

func checkUser() *checkers.Checker {
	opts := options{}
	psr := flags.NewParser(&opts, flags.Default)
	_, err := psr.Parse()
	if err != nil {
		os.Exit(1)
	}

	db := mysql.New("tcp", "", fmt.Sprintf("%s:%s", opts.Host, opts.Port), opts.User, opts.Pass, "")
	err = db.Connect()
	if err != nil {
		return checkers.Critical("couldn't connect DB")
	}
	defer db.Close()
	
	stmt, err := db.Prepare("SELECT COUNT(*) FROM mysql.user WHERE Host=? AND User=?")
	if err != nil {
		return checkers.Critical(fmt.Sprintf("db.Prepare:%s",err))
	}

	stmt.Bind(opts.AccountHost, opts.AccountName)
	rows, _, err := stmt.Exec()
	num := rows[0].Int64(0)

	if num < 1 {
		return checkers.Critical(fmt.Sprintf("user '%s'@'%s' not found", opts.AccountName, opts.AccountHost))
	}
	return checkers.Ok(fmt.Sprintf(fmt.Sprintf("user '%s'@'%s' exists", opts.AccountName, opts.AccountHost)))
}

