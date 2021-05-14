Main Project [StatuZpage](https://github.com/nopp/statuzpage)

# StatuZpage API

Responsible for receive incidents from statuzpage-agent, and return informations about urls to statuzpage-ui.

## Configurations:

Default config dir: /etc/statuzpage-api/config.json
* mysql-host: ip/dns
* mysql-user: mysql user
* mysql-password: mysql password
* mysql-db: statuzpage(default)
* token: anyvalue "secret"

## Build:
$ go build

## Start
$ ./statuzpage-api

## Kubernetes Image
noppp/statuzpage-api:tagname
