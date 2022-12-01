
# GO Movies REST API

This is a little REST API builded with GO, and using JSON files to storage data.
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