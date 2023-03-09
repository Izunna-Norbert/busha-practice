# busha-practice
Golang Project API, getting Movies and Characters
Technical Documentation for Movie Database System

This document provides a technical overview of a RESTful API service that can be used to interact with a movie database. The API is written in Golang, and uses the gin and gorm libraries to handle HTTP requests and database interactions respectively.

Requirements
To use the RESTful API service, you will need the following:
Golang installed on your machine.
An API client like Postman to send HTTP requests to the endpoints.
A database management system (DBMS) PostgreSQL.

```
git clone https://github.com/Izunna-Norbert/busha-practice

```


Navigate to the project directory

```
cd busha-practice
```


Install the required dependencies

```
go mod download
```


Build and Run the application

```
go build && go ./busha-practice
```


The API endpoints can now be accessed at http://localhost:8000/ .

API Endpoints
List movies
Returns a list of movies containing the name, opening crawl and comment count.

Endpoint: /api/v1/movies

HTTP Method: GET

Request Parameters: None




Example Request:

```

GET http://localhost:8000/movies

```

Example Response:

```
[
    {
        "name": "A New Hope",
        "opening_crawl": "It is a period of civil war...",
        "comment_count": 2
    },
    {
        "name": "The Empire Strikes Back",
        "opening_crawl": "It is a dark time for the Rebellion...",
        "comment_count": 1
    }
]

```


Add a new comment for a movie
Adds a new comment for a movie.

Endpoint: /movies/{id}/comments

HTTP Method: POST

Request Parameters:
```

Parameter	Type	      Description
id	      string	The ID of the movie for which to add a comment

```

Example Request:

```

POST http://localhost:8000/movies/1/comments
Content-Type: application/json

{
    "comment": "I love this movie!",
    "name": "Agu Norbert",
}


Example Response:
{
    "id": 1,
    "movie_id": 1,
    "comment": "I love this movie!"
}

```

List comments for a movie
Returns a list of comments for a movie.

Endpoint: /movies/{id}/comments

HTTP Method: GET

Request Parameters:

```
Parameter	Type	      Description
id	      string	The ID of the movie for which to list comments
```

Example Request:

GET http://localhost:8000/movies/1/comments
Example Response:

```
[
    {
        "id": 1,
        "movie_id": 1,
        "comment": "I love this movie!"
    },
    {
        "id": 2,
        "movie_id": 1,
        "comment": "This movie is awesome!"
    }
]
```


Get list of characters for a movie
Returns a list of characters for a movie.

Endpoint: /movies/{id}/characters

HTTP Method: GET

Request Parameters:

Parameter	Type	      Description
id	      string	The ID of the movie for which to list characters

Example Request:





