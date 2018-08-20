# go-config-example
A example code of manage configuration and secure information in Go

## Getting started

```bash
$ docker-compose build

$ docker-compose up -d

$ docker-compose ps
    Name                 Command             State           Ports         
---------------------------------------------------------------------------
app            ./server                      Up      0.0.0.0:8080->8080/tcp
master_mysql   docker-entrypoint.sh mysqld   Up      0.0.0.0:3306->3306/tcp

$ curl -i http://localhost:8080/.healthcheck
HTTP/1.1 200 OK
Date: Mon, 20 Aug 2018 02:52:33 GMT
Content-Length: 0
```
