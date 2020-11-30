package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jessevdk/go-flags"
	"github.com/kazeburo/go-mysqlflags"
	"github.com/mackerelio/checkers"
)

// Version by Makefile
var version string

type Opts struct {
	mysqlflags.MyOpts
	Timeout     time.Duration `long:"timeout" default:"5s" description:"Timeout to connect mysql"`
	AccountName string        `short:"a" long:"account-name" arg:"String" required:"true" description:"account user name"`
	AccountHost string        `short:"n" long:"account-host" arg:"String" required:"true" description:"account user host"`
	Version     bool          `short:"v" long:"version" description:"Show version"`
}

func main() {
	ckr := checkUser()
	ckr.Name = "MySQL User"
	ckr.Exit()
}

func checkUser() *checkers.Checker {
	opts := Opts{}
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

	db, err := mysqlflags.OpenDB(opts.MyOpts, opts.Timeout, false)
	if err != nil {
		return checkers.Critical(fmt.Sprintf("couldn't connect DB: %v", err))
	}
	defer db.Close()

	var num int64

	ctx, cancel := context.WithTimeout(context.Background(), opts.Timeout)
	defer cancel()
	ch := make(chan error, 1)

	go func() {
		ch <- db.QueryRow(
			"SELECT COUNT(*) FROM mysql.user WHERE Host=? AND User=?",
			opts.AccountHost,
			opts.AccountName,
		).Scan(&num)
	}()

	select {
	case err = <-ch:
		// nothing
	case <-ctx.Done():
		err = fmt.Errorf("connection or query timeout")
	}

	if err != nil {
		return checkers.Critical(fmt.Sprintf("Couldn't fetch mysql.user: %v", err))
	}

	if num < 1 {
		return checkers.Critical(fmt.Sprintf("user '%s'@'%s' not found", opts.AccountName, opts.AccountHost))
	}
	return checkers.Ok(fmt.Sprintf(fmt.Sprintf("user '%s'@'%s' exists", opts.AccountName, opts.AccountHost)))
}
