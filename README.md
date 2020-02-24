# Article Maker

Article maker is a REST API server written in go.

## Getting Started

## Install
Clone project from github



## Run
Build
```bash
make build
```
Run
```bash
make run
```
## Testing
```bash
make test
```

## DB
The sql script with the DB schema can be found in the sql directory

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
## License

[MIT](https://choosealicense.com/licenses/mit/)
