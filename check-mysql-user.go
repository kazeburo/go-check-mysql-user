package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
)

// Version by Makefile
var version string

type options struct {
	Host        string `short:"H" long:"host" default:"localhost" description:"Hostname"`
	Port        string `short:"p" long:"port" default:"3306" description:"Port"`
	User        string `short:"u" long:"user" default:"root" description:"Username"`
	Pass        string `short:"P" long:"password" default:"" description:"Password"`
	AccountName string `short:"a" long:"account-name" arg:"String" required:"true" description:"account user name"`
	AccountHost string `short:"n" long:"account-host" arg:"String" required:"true" description:"account user host"`
	Version     bool   `short:"v" long:"version" description:"Show version"`
}

func main() {
	ckr := checkUser()
	ckr.Name = "MySQL User"
	ckr.Exit()
}

func checkUser() *checkers.Checker {
	opts := options{}
	psr := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash)
	_, err := psr.Parse()
	if opts.Version {
		fmt.Fprintf(os.Stderr, "Version: %s\nCompiler: %s %s\n",
			version,
			runtime.Compiler,
			runtime.Version())
		os.Exit(0)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
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
		return checkers.Critical(fmt.Sprintf("db.Prepare:%s", err))
	}

	stmt.Bind(opts.AccountHost, opts.AccountName)
	rows, _, err := stmt.Exec()
	num := rows[0].Int64(0)

	if num < 1 {
		return checkers.Critical(fmt.Sprintf("user '%s'@'%s' not found", opts.AccountName, opts.AccountHost))
	}
	return checkers.Ok(fmt.Sprintf(fmt.Sprintf("user '%s'@'%s' exists", opts.AccountName, opts.AccountHost)))
}
