## About The Project

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
This command allows you to list books in book list.

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

### get command
```
go run main.go -delete <ID>
```

### buy command
```
go run main.go -buy <ID> <quantity>
```
