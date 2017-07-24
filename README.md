# GRUD <a href="https://travis-ci.org/lohanbodevan/grud"><img alt="Travis Status" src="https://travis-ci.org/lohanbodevan/grud.svg?branch=master"></a>

GRUD is **G**olang C**RUD** REST API using MongoDB and JWT security standard example.


## Development
### Requirements
* Docker 17.6.x
* Docker Compose 1.11.x

### Build Containers
```
make docker-build
```

### Run API
```
make docker-run
```

Access: http://localhost:8080/tvseries

### Run Tests
```
make docker-test
```

## Endpoints
### Create
POST /tvseries

To create TV Series record, do a `post` HTTP request in `/tvseries` endpoint

**Token**

This endpoint is protected and you should pass `Authorization` header with token.
To get token, see `Login` endpoint

Header:
```
Authorization: Bearer <token>
```
Payload:
```
{
    "title": "Braking Bad",
    "description": "Branking Bad é uma premiada série de televisão estadunidense criada e produzida por Vince Gilligan que retrata a vida do químico Walter White, um homem brilhante frustrado em dar aulas para adolescentes do ensino médio enquanto lida com um filho sofrendo de paralisia cerebral, uma esposa grávida e dívidas intermináveis. White, então, é diagnosticado com um cancro no pulmão - o que o leva a sofrer um colapso emocional e abraçar uma vida de crimes para pagar suas dívidas hospitalares e dar uma boa vida aos seus filhos. Walter resolve produzir metanfetamina com seu ex-aluno, Jesse Pinkman.",
    "casting": [
      {
        "name": "Bryan Cranston"
      },
      {
        "name": "Anna Gunn"
      },
      {
        "name": "Aaron Paul"
      }
    ],
    "stars": 5
}
```

### Read
GET /tvseries

To read all TV Series available in our catalog, do a `get` HTTP request in `/tvseries` endpoint.

This endpoint is open

Reponse:
```
[
  {
    "code": "c0eb63ef-8570-4aad-82bd-79c36502e755",
    "title": "Braking Bad",
    "description": "Branking Bad é uma premiada série de televisão estadunidense criada e produzida por Vince Gilligan que retrata a vida do químico Walter White, um homem brilhante frustrado em dar aulas para adolescentes do ensino médio enquanto lida com um filho sofrendo de paralisia cerebral, uma esposa grávida e dívidas intermináveis. White, então, é diagnosticado com um cancro no pulmão - o que o leva a sofrer um colapso emocional e abraçar uma vida de crimes para pagar suas dívidas hospitalares e dar uma boa vida aos seus filhos. Walter resolve produzir metanfetamina com seu ex-aluno, Jesse Pinkman.",
    "casting": [
      {
        "name": "Bryan Cranston"
      },
      {
        "name": "Anna Gunn"
      },
      {
        "name": "Aaron Paul"
      }
    ],
    "stars": 5
  }
]
```

### Update
PUT /tvseries/code

To update a TV Series record, do a `put` HTTP request in `/tvseries/code` endpoint

**Token**

This endpoint is protected and you should pass `Authorization` header with token.
To get token, see `Login` endpoint

**Code**

The `code` is returned in `GET /tvseries` endpoint

Header:
```
Authorization: Bearer <token>
```

Payload:
```
{
    "title": "Braking Bad",
    "description": "Branking Bad é uma premiada série de televisão estadunidense criada e produzida por Vince Gilligan que retrata a vida do químico Walter White, um homem brilhante frustrado em dar aulas para adolescentes do ensino médio enquanto lida com um filho sofrendo de paralisia cerebral, uma esposa grávida e dívidas intermináveis. White, então, é diagnosticado com um cancro no pulmão - o que o leva a sofrer um colapso emocional e abraçar uma vida de crimes para pagar suas dívidas hospitalares e dar uma boa vida aos seus filhos. Walter resolve produzir metanfetamina com seu ex-aluno, Jesse Pinkman.",
    "casting": [
      {
        "name": "Bryan Cranston"
      },
      {
        "name": "Anna Gunn"
      },
      {
        "name": "Aaron Paul"
      },
      {
        "name": "Lohan Bodevan"
      }
    ],
    "stars": 5
}
```

### Delete
DELETE /tvseries/code

To delete a TV Series record, do a `delete` HTTP request in `/tvseries/code` endpoint

**Token**

This endpoint is protected and you should pass `Authorization` header with token.
To get token, see `Login` endpoint

**Code**

The `code` is returned in `GET /tvseries` endpoint

Header:
```
Authorization: Bearer <token>
```

### Login
POST /login
To autheticate yourself on API, do a `post` HTTP request in `/login` endpoint

Payload:
```
{
    "email": "admin@example.com",
    "password": "somecoolpassword"
}
```

PS.: This repository did not has DB seeds.
You need to create `user` document in your MongoDB instance with `email` and sha256 encrypted `password`

## Docs
There are environment and collection for [Postman](https://www.getpostman.com/) in `docs` folder to facilitate play with api.
