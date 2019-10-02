VERSION=0.0.2
LDFLAGS=-ldflags "-X main.Version=${VERSION}"
GO111MODULE=on

all: check-mysql-user

.PHONY: check-mysql-user

check-mysql-user: check-mysql-user.go
	go build $(LDFLAGS) -o check-mysql-user

linux: check-mysql-user.go
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o check-mysql-user

clean:
	rm -rf check-mysql-user

tag:
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin master
	goreleaser --rm-dist
