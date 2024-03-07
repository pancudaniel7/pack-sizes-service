# Pack Sizes Algorithm
The pack sizes algorithm optimally calculates the required quantities of different pack sizes to fulfill an order, minimizing the total number of packs needed.

### Building and Running Locally
To build the project, navigate to the project directory and run:
```
go build -o pack-sizes-service cmd/pack-sizes-service/main.go
```
To run the project locally:
```
./pack-sizes-service
```

### Running with Docker Compose
Navigate project directory and run:
```
docker-compose -f deployments/docker-compose.yml up --build
```
This command builds the service and starts it, making the UI accessible at `localhost:8080`.

### Accessing the UI
Open your web browser and navigate to `http://localhost:8080` to interact with the UI for the pack sizes service.

### Testing the REST API with Postman
Import the provided Postman collection and environment files located in `docs/postman/` into Postman to test the REST API endpoints.

### Running Unit Tests
To run the unit tests for the internal package with verbose output, navigate to the project directory and execute:
```
go test -v ./internal/...
```
This command provides detailed output for each test, including which tests are running and their results.
