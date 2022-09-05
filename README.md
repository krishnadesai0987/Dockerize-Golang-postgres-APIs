#   Golang, Postgres and Rest APIs(using gorila/mux package).Covered CURD operation on Postgres DB using REST APIs.
##  Install Postgres on your machine/server. you can find the tutorials on YouTube/google.
##  Install Postman to make GET/POST/PUT/DELETE and many more requests to your application.**

you need to install following packages to make it run:
- go get -u github.com/lib/pq
- go get -u github.com/gorilla/mux

## Usage
- Get the 2 packages from above
- Run go run main.go
- Use Postman to test the requests
 
## Requests
- Get all books - [GET] http://localhost:8000/books
- Get a book by id - [GET] http://localhost:8000/books/id 
- Create a book - [POST] http://localhost:8000/books - (body: 'x-www-form-urlencoded', pass in id)
- Update a book by id - [PUT] http://localhost:8000/books/id - (body: 'x-www-form-urlencoded', pass updated details)
- Delete a book by id - [DELETE] http://localhost:8000/books/id
- Delete all books - [DELETE] http://localhost:8000/books/

```diff
## Book table

CREATE TABLE books (
    id PRIMARY KEY int,
    title text,
    author text,
    description character(100))  ); 
```
## Install Docker on your machine/server
Follow the following commnad to deploy your application on docker
- docker-compose up -d --build -> automate the configuration require to build images and container
- docker logs -f go_image_name-> check the logs of .go file
- docker ps ->verify the running containers
- curl http://localhost:8000/books -> you will be able to hit API.

    


