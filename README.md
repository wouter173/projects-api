# projects-api

API for my projects display on https://wouterdb.nl/ from mongodb.<br>
The API is written in Go using fiber and the official mongodb library.<br>
Currently live on https://projects.wouter.cloud.<br>

### env
```sh
DBURI > mongodb connection string 
DB > database name
COLL > collection name
PORT > the port the server goes live on
```

### run
```sh
$ go run *.go
```

### Notes
This project is build to deploy on a dokku stack.
But it can be deploy to an heroku app.
