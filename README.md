# projects-api

API for my projects display on https://wouterdb.nl/.<br>
The API is written in Go using fiber.<br>

### env
```sh
PROJECTS_PATH > the path the projects are on
PORT > the port the server goes live on
```

### run
```sh
$ go run .
```

### TODO

- [ ] feature: set the CORS header to only allow requests from wouterdb.nl wouter173.nl and localhost for dev