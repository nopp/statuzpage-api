# StatuZpage API

Responsible for receive incidents from statuzpage-agent, and return informations about urls to statuzpage-ui.

Configurations:
===============
Default config dir: /etc/statuzpage-api/config.json
* mysql-host: ip/dns
* mysql-user: mysql user
* mysql-password: mysql password
* mysql-db: statuzpage(default)
* token: anyvalue "secret"
* hostport: ip:8000 to bind

Build:
======
$ go build

Start
=====
$ ./statuzpage-api
