# Gomen

Gomen is GoLang simple boilerplate for building Microservice architecture with multiple transport layer such as GRPC & Rest Api

## Installation

#### Clone the project

```bash
git clone https://github.com/fachryansyah/gomen.git
```

#### Install required library
```bash
$ cd gomen && go mod tidy
```

Generate Proto file
```bash
$ make gen-proto
```

#### Running
Serving the project
```bash
$ make serve
```

## Avaible CLI Commands
|COMMAND|DESCRIPTION|
|-------|-----------|
|```go run main.go serve```|running server instance|
|```go run main.go migrate```|running auto migration|

## Todo
- [X] Create Folder & File Structure
- [ ] Added ORM Features
- [X] CLI Commands
- [ ] Auto Migration
- [X] Rest API
- [X] GRPC
- [ ] Auto Generated Code && CLI Tools

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
