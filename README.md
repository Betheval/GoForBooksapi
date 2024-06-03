# BooksAPI
MyApp is a RESTful API written in Go using the Gin framework and MongoDB as the database. This API allows managing a catalog of books, including operations to retrieve, add, and delete books.


## Prerequisites

- [Go](https://golang.org/doc/install) 1.18 or higher
- [MongoDB](https://docs.mongodb.com/manual/installation/) installed and running
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) for cloning the repository

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/Betheval/GoForBooksapi.git
    cd GoForBooksapi
    ```

2. Create a `.env` file in the project root with the following configuration:

    ```env
    MONGODB_URI=mongodb://localhost:27017
    ```

3. Install the project dependencies:

    ```sh
    go mod tidy
    ```

4. Start the server:

    ```sh
    go run main.go
    ```

    The server will run at `http://localhost:8080`.

## Endpoints

### `GET /ping`

Returns a health check message.

#### Response Example:

```json
{
    "message": "pong"
}

```

### `GET /books`

Retrieves the first 100 books.

#### Response Example:
```json
[
    {
        "id": "60d5ec49fc13ae1a7e000000",
        "title": "The Catcher in the Rye",
        "author": "J.D. Salinger",
        "isbn": "9780316769488",
        "publisher": "Little, Brown and Company",
        "year": 1951
    },
]
```

### `GET /book/id/:id`

Retrieves a book by its ID.
#### Parameters:
```json
{
    "id": "The book's ID (MongoDB ObjectID type)".
}
```
#### Response Example:
```json
{
    "id": "60d5ec49fc13ae1a7e000000",
    "title": "The Catcher in the Rye",
    "author": "J.D. Salinger",
    "isbn": "9780316769488",
    "publisher": "Little, Brown and Company",
    "year": 1951
}
```

### `GET /book/isbn/:isbn`

Retrieves a book by its ISBN.
#### Parameters:
```json
    "isbn": "The book's ISBN".
```
#### Response Example:
```json
{
    "id": "60d5ec49fc13ae1a7e000000",
    "title": "The Catcher in the Rye",
    "author": "J.D. Salinger",
    "isbn": "9780316769488",
    "publisher": "Little, Brown and Company",
    "year": 1951
}
```
### `POST /book`

Adds a new book.
#### Request Example:
```json
{
    "title": "To Kill a Mockingbird",
    "author": "Harper Lee",
    "isbn": "9780060935467",
    "publisher": "J.B. Lippincott & Co.",
    "year": 1960
}
```
#### Response Example:
```json
{
    "id": "60d5ec49fc13ae1a7e000001"
}
```
### `DELETE /book/deletion/:id`

Deletes a book by its ID.
#### Parameters:
```json
    "id": "The book's ID (MongoDB ObjectID type)".
```
#### Response Example:
```json
{
    "message": "Book deleted"
}
```
## Code Structure
### main.go

The main file that starts the application, connects to the database, and sets up the routes.
config/config.go

Handles the MongoDB connection configuration.
models/book.go

Defines the data model for books.
controllers/book_controller.go

Controllers to handle the logic for the endpoints related to books.
routes/routes.go

Defines the API routes.
