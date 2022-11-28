## Keep yourself busy with BoredAPI on my /gin-server

#### Prerequisites:
* Golang

## Set up the project
```bash
# install Gin Framework
go get github.com/gin-gonic/gin
```

## Run App
```bash
go run main.go
# Or 
go run .
```


## Use cases
```bash

# --help
curl http://localhost:8080/

# Review MyBoredList
curl http://localhost:8080/mybored-list/

# Find Bored thing to do
curl http://localhost:8080/random-bored-thing/

# Add to MyBoredList last bored suggestion
curl -i -X POST http://localhost:8080/add-bored/

# Remove item from MyBoredList by id
curl -i -X POST http://localhost:8080/remove-bored/

```