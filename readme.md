# Backend Coding Challenge

## Run the app on the docker
 - Created the docker and docker compose file to run the go code and database into the separate container
```
make run
```
![!\[Alt text\](image-3.png)](<doc/image 1.png>)


### Task 1 - Create an article
- Method: `POST`
- Path: `/articles`
```
  curl --location 'http://localhost:8080/articles' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Second Article",
    "content": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
    "author": "John"
}'
```
![!\[Alt text\](image.png)](<doc/image 2.png>)

### Error in Create an article
![Alt text](<doc/image 5.png>)

### Task 2 - Get article by id
- Method: `GET`
- Path: `/articles/<article_id>`
```
curl --location 'http://localhost:8080/articles/1'
```
![!\[Alt text\](image-1.png)](<doc/image 3.png>)

### Error in Get article by id
![Alt text](<doc/image 6.png>)

### Task 3 - Get all article
- Method: `GET`
- Path: `/articles`
```
curl --location 'http://localhost:8080/articles'
```
![Alt text](<doc/image 4.png>)

### Error in Get all article
![Alt text](<doc/image 7.png>)

## Clean code / Development practice
- Followed by using the separate business logic, db, utility, models, constants, db query, etc
```
tree
.
├── Dockerfile
├── Makefile
├── api
│   └── swagger.yaml
├── doc
│   ├── image 1.png
│   ├── image 10.png
│   ├── image 11.png
│   ├── image 2.png
│   ├── image 3.png
│   ├── image 4.png
│   ├── image 5.png
│   ├── image 6.png
│   ├── image 7.png
│   ├── image 8.png
│   └── image 9.png
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── controller
│   │   ├── controllers.go
│   │   └── controllers_test.go
│   ├── docs.go
│   └── routes
│       ├── routes.go
│       └── routes_test.go
├── main.go
├── mocks
│   ├── mock_handlers.go
│   ├── mock_routes.go
│   └── mock_service.go
├── pkg
│   ├── appconstant
│   │   ├── error.go
│   │   ├── log.go
│   │   ├── message.go
│   │   └── port.go
│   ├── db
│   │   ├── db.go
│   │   └── db_test.go
│   ├── models
│   │   ├── articles.go
│   │   ├── docs.go
│   │   └── response.go
│   ├── repository
│   │   └── dbrepo
│   │       ├── postgres_dbrepo.go
│   │       └── postgres_dbrepo_test.go
│   └── utility
│       ├── utils.go
│       └── utils_test.go
├── readme.md
└── services
    └── articles
        ├── articles_service.go
        └── articles_services_test.go

16 directories, 42 files, 13146 lines
```

## Test coverage
```
make test
```
- Achieved >75% test coverage

```
make cover
```
```
ok      backend/internal/routes (cached)        coverage: 100.0% of statements
ok      backend/pkg/db  (cached)                coverage: 91.7% of statements
ok      backend/pkg/repository/dbrepo           coverage: 88.6% of statements
ok      backend/internal/handlers       0.358s  coverage: 78.7% of statements
ok      backend/pkg/utility     (cached)        coverage: 86.7% of statements
ok      backend/services/articles               coverage: 100.0% of statements
```
## Error logging in containers
![Alt text](<doc/image 8.png>)

## Swagger run
```
make swagger
```
![Alt text](<doc/image 9.png>)

## Swagger UI
![Alt text](<doc/image 10.png>)

![Alt text](<doc/image 11.png>)