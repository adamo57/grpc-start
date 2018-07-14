# gRPC-start

## Steps For Usage

1. Start the gRPC Server
    ```bash
    go run server/main.go
    ```
2. Start the HTTP server
    ```bash
    go run main.go
    ```
3. cURL or use a REST client to hit endpoint
    ```bash
    curl -X POST -k localhost:8080/v1/user -d '{"name": "Adam Ouellette", "age": 24, "email": "adamouellette57@gmail.com"}'
    ```