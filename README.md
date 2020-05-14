# Gomen

Gomen is simple framework for building Microservice architecture with Go Language

## Installation

#### Clone the project

```bash
git clone https://github.com/fachryansyah/simple-golang-boilerplate.git
```

#### Install required library

Mysql database
```bash
$ go get github.com/go-sql-driver/mysql
```
Object Relation Mapping
```bash
$ go get -u github.com/jinzhu/gorm
```
Proto Gen
```
$ go get -u github.com/golang/protobuf/protoc-gen-go
```
GRPC
```
$ go get google.golang.org/grpc
```
Generate Proto file
```
$ cd app/services
$ PATH=$PATH:$GOPATH/bin/ protoc --go_out=plugins=grpc:. *.proto
```

#### Running
Serving the project
```bash
go run main.go serve
```

## Avaible CLI Commands
|COMMAND|DESCRIPTION|
|-------|-----------|
|```go run main.go serve```|running server instance|
|```go run main.go migrate```|running auto migration|

## Info
Gomen comes like MVC pattern, you can create logic bussines on /app/http/controller.

## Todo
- [X] Create Folder & File Structure
- [X] Added ORM Features
- [X] CLI Commands
- [X] Auto Migration
- [ ] JWT Authentication
- [ ] Auth Service Provider
- [ ] File Storage Provider
- [ ] Auto Generated Code

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
