# projects-api

API for my projects display on https://wouterdb.nl/.<br>
The API is written in Go using fiber.<br>

### endpoints
```
/
```
Main endpoint listing all projects w/ metadata.

```
/:id
```
Send back the markdown body from the defined project with the :id paramater.

```
/:id/meta
```
Get all the metadata from the defined project formatted in JSON.

### env
```sh
PROJECTS_PATH > the path the projects are on
PORT > the port the server goes live on
```

### run

Run from source:
```sh
$ go run .
```

Or build a docker image:
```
$ make build
```

### TODO

- [ ] feature: set the CORS header to only allow requests from wouterdb.nl wouter173.nl and localhost for dev
- [ ] fix: fix in GetMeta that if there is no metadata just serve back and empty metadata struct
