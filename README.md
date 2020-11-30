# go-check-mysql-user

check mysql user exists

## usage

```
Usage:
  check-mysql-user [OPTIONS]

Application Options:
      --defaults-extra-file= path to defaults-extra-file
      --mysql-socket=        path to mysql listen sock
  -H, --host=                Hostname (default: localhost)
  -p, --port=                Port (default: 3306)
  -u, --user=                Username (default: root)
  -P, --password=            Password
      --database=            database name connect to
      --timeout=             Timeout to connect mysql (default: 5s)
  -a, --account-name=        account user name
  -n, --account-host=        account user host
  -v, --version              Show version

Help Options:
  -h, --help                 Show this help message
```

```
$ check-mysql-user --account-name=readuser --account-host=localhost
MySQL User OK: user 'readuser'@'localhost' exists
```

  ## Install

Please download release page or `mkr plugin install kazeburo/go-check-mysql-user`.
