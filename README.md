# API Gateway

This repository contains the GraphQL API Gateway for the BetterGR system. The API gateway acts as an interface between the frontend application and the various microservices.

## Setup

### Prerequisites

- Go (latest stable version recommended)
- Make (for running the Makefile commands)

### Environment Configuration

Make sure to have a `.env` file in the root directory. Here is an example:

```env
# API Gateway Configuration
API_GATEWAY_PORT=1234

# Authentication Settings
CLIENT_SECRET=**********
KEYCLOAK_URL=http://auth.betterGR.org
REDIRECT_URI=http://localhost:3000/callback

# Microservice Addresses
GRADES_PORT=localhost:50051
STUDENTS_PORT=localhost:50052
HOMEWORK_PORT=localhost:50053
COURSES_PORT=localhost:50054
STAFF_PORT=localhost:50055
```

### Running the API Gateway

To run the API Gateway server:

```bash
go run server.go
```

Alternatively, you can use the provided PowerShell script:

```bash
./run-api-gateway.ps1
```

## Development

### Modifying the GraphQL Schema

When changing the schema file (`schema.graphqls`), make sure to regenerate the Go code by running:

```bash
gqlgen generate
```

This will update all generated files based on your schema changes.

## Tools API for AI Models

The API Gateway includes a Tools API that enables AI models to interact with the BetterGR system programmatically.

### Endpoints

- `GET /tools`: Returns a list of all available tools and their parameters
- `POST /tools/execute`: Executes a specified tool with provided parameters

### Authentication

Tools API requests require authentication. Include an `Authorization` header with a JWT token:

```
Authorization: Bearer your-jwt-token
```

### Example Usage

To see examples of how to use the Tools API from an AI model, refer to:
- JavaScript: `examples/ai_client_example.js`
- Python: `examples/ai_client_example.py`

For detailed documentation on the Tools API and available tools, see [API Tools Documentation](docs/ai_tool_usage.md).

## License

This project is licensed under the Apache 2.0 License. See the LICENSE file for more details.
