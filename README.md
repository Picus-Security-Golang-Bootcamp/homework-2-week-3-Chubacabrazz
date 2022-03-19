## About The Project

My article about concurrency in golang ; https://medium.com/@denzszr/concurrency-in-golang-a7843e2d3073


This app has a book list and it has multiple functions;

•list --> lists all books

•search --> searchs book by name

•get --> prints book details by ID

•delete --> delete books by ID

•buy --> buys books by ID and quantity


### Build With
Go

## How to use

### list command
```
go run main.go -list all
```

### search command 
```
go run main.go search <bookName>
go run main.go search Lord of the Rings 
```
It works case insensitive.

### get command
```
go run main.go -get <ID>
```

### delete command
```
go run main.go -delete <ID>
```

### buy command
```
go run main.go -buy <ID> <quantity>
```
