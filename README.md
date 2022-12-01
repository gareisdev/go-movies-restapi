
# GO Movies REST API

[![Go Report Card](https://goreportcard.com/badge/github.com/gareisdev/go-movies-restapi)](https://goreportcard.com/report/github.com/gareisdev/go-movies-restapi)

This is a little REST API built with GO, and uses JSON files to store data.
## Endpoints

### Get all movies

```http
  GET /api/movies
```

### Get movie

```http
  GET /api/movies/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Id of movie to fetch |


### Create movie

```http
  POST /api/movies
```

#### Body

```json
{
    "isbn": "string",
    "title": "string",
    "director": {
        "firstname": "string",
        "lastname": "string"
    }
}
```

### Update movie

```http
  PUT /api/movies/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Id of movie to fetch |

#### Body

```json
{
    "isbn": "string",
    "title": "string",
    "director": {
        "firstname": "string",
        "lastname": "string"
    }
}
```

### Delete movie

```http
  DELETE /api/movies/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `number` | **Required**. Id of movie to fetch |


## Run Locally

Clone the project

```bash
  git clone https://github.com/gareisdev/go-movies-restapi
```

Go to the project directory

```bash
  cd go-movies-restapi
```

Up containers

```bash
  docker-compose up
```

## License

[MIT](https://choosealicense.com/licenses/mit/)