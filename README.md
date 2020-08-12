# Prometheus Histogram example

This is a example of how prometheus histogram works.

## Run
* `docker-compose up -d` to start prometheus server (it will run at :9090)
* `go run main.go` to start the application (it will expose metrics on `localhost:9091`)