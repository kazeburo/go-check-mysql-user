VERSION=0.0.1

all: check-mysql-user

.PHONY: check-mysql-user

gom:
	go get -u github.com/mattn/gom

bundle:
	gom install

check-mysql-user: check-mysql-user.go
	gom build -o check-mysql-user

linux: check-mysql-user.go
	GOOS=linux GOARCH=amd64 gom build -o check-mysql-user

fmt:
	go fmt ./...

dist:
	git archive --format tgz HEAD -o check-mysql-user-$(VERSION).tar.gz --prefix check-mysql-user-$(VERSION)/

clean:
	rm -rf check-mysql-user check-mysql-user-*.tar.gz

