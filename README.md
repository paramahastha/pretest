# BE Pre-Test

A project that runs a Go server (Gin-Gonic) and a PostgreSQL via two separate containers, using Docker Compose.

## Required

```
Docker : https://docs.docker.com/install/

Docker Compose : https://docs.docker.com/compose/install/

```

## How to run

```
export APP_ENV=dev

docker-compose up --build
```

For development, the `api/` and `database/` directories have their own docker containers, which are configured via the `docker-compose.yml` file.

The server is up at `localhost:5000/v1/` and it proxies internally to the server using the linked name as `localhost:8080`. for example `localhost:5000/v1/products`
