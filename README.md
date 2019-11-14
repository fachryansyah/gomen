# Gomen

Gomen is simple boilerplate for building Microservice with golang.

## Installation

#### Clone the project

```bash
git clone https://github.com/fachryansyah/simple-golang-boilerplate.git
```

#### Install required library

Mysql database
```bash
go get github.com/go-sql-driver/mysql
```
Router Web service
```bash
go get github.com/gorilla/mux
```
Socket Connection
```bash
go get github.com/googollee/go-socket.io
```
Object Relation Mapping
```bash
go get -u github.com/jinzhu/gorm
```

#### Running
Serving the project
```bash
go run main.go serve
```

## Info
Gomen comes like MVC pattern, you can create logic bussines on /app/http/controller.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
