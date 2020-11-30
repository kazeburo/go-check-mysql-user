# go-check-mysql-user

check mysql user exists

## usage

```
Usage:
  check-mysql-user [OPTIONS]

Application Options:
  -H, --host=         Hostname (localhost)
  -p, --port=         Port (3306)
  -u, --user=         Username (root)
  -P, --password=     Password
  -a, --account-name= account user name
  -n, --account-host= account user host

Help Options:
  -h, --help          Show this help message
```

```
$ check-mysql-user --account-name=readuser --account-host=localhost
MySQL User OK: user 'readuser'@'localhost' exists
```

  ## Install

Please download release page or `mkr plugin install kazeburo/go-check-mysql-user`.
