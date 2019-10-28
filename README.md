# Geo location service

`geolocation` service provides data via gRPC and loads location data from CSV file  

## Required environment variables

| Variable                      | Description                                | Default    | Optional |
|-------------------------------|--------------------------------------------|:----------:|:--------:|
| `GEOLOCATION_DB_DSN`          | MYSQL DSN                                  |     -      |    -     |
| `GEOLOCATION_GRPC_PORT`       | TCP port on which to start gRPC server     |    5000    |   true   |
| `GEOLOCATION_MODE`            | Environment mode                           |    dev     |   true   |

## Run linting and testing

```
.scripts/check.sh
```

## Build and run locally

`geolocation` uses [dep](https://golang.github.io/dep/) to manage dependencies.

```
dep ensure -vendor-only -v
make
# export env variables described above ...
./geolocation
```

## Required environment variables to build Docker image

| Variable                   | Description                                             |
| ---------------------------|:-------------------------------------------------------:|
| `GEOLOCATION_GITLAB_TOKEN` | The token for download of repositories from gitlab.com  |
