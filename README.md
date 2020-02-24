# Article Maker

Article maker is a REST API server written in go.

This api is used for managing article data with support for operations such as.

* Creating articles
* Updating articles
* Getting all aritcles
* Get single article
* Get filted articles depending on publisher, category, publish date and created date

## Getting Started

## Install
Clone project from github

## DB
The server uses a Mysql database for persisting data during operation.
The sql script with the DB schema can be found in the sql directory.

### Db Config
Replace the const values in config.go with the respective database details
```go
const (
	DATABASE_USER     = ""
	DATABASE_PASSWORD = ""
	DATABASE_NAME     = ""
)

```

## Build
This generates a compiled executable that can be directly executed from the terminal.
```bash
make build
```
##Run
This starts the Server and listens on port :9000, the server can be accessed on http://localhost:9000
```bash
make run
```
## Testing
```bash
make test
```



## Endpoints

| Method        | Path           | Comment  |
| ------------- |:-------------| :-----|
| GET     | /article | Gets all articles
| GET     | /article/1      |   Gets single article with id of 1 |
| POST | /article      |    Creates a new article |
| PUT | /article      |    Updates article |
| DELETE | /article/1      |    Deletes article with id of 1 |


### Filtered results
| Method        | Path           | Comment  |
| ------------- |:-------------| :-----|
| GET     | /article?category=science | Gets all articles in the category of science
| GET     | /article?publisher=John Doe | Gets all articles by John Doe
| GET     | /article?published_at=2020-02-23 10:10 | Gets all articles according to published on 2020-02-23 10:10
| GET     | /article?created_at=2020-02-23 10:10 | Gets all articles that where create on 2020-02-23 10:10

## Article example
```json
{

  "title": "Lorem ipsum dolor sit amet",
  "body": "Lorem ipsum dolor sit amet, ut la quis nostrud exercitation ullamco laboris nisi ut al"
  "category": "Lorem ipsum",
  "publisher": "John Doe",
  "created_at": "2020-01-19 11:50:39",
  "published_at": "2020-01-20 10:15:50"

}
```

## Creating article
To create article submit a post request to /article
#### post body
```json
{
  "title": "Lorem ipsum dolor sit amet",
  "body": "Lorem ipsum dolor sit amet, ut la quis nostrud exercitation ullamco laboris nisi ut al"
  "category": "Lorem ipsum",
  "publisher": "John Doe",
  "published_at": "2020-01-20 10:15:50"
}
```

## Updating article
To update article submit a put request to /article
#### put body
```json
{
  "id":1,
  "title": "Lorem ipsum dolor sit amet",
  "body": "Lorem ipsum dolor sit amet, ut la quis nostrud exercitation ullamco laboris nisi ut al"
  "category": "Lorem ipsum",
  "publisher": "John Doe",
  "published_at": "2020-01-20 10:15:50"
}
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
