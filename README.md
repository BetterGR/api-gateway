# api-gateway

## Running the api-gateway server

make sure to have a `.env` file. here is an example:

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

```bash
go run server.go
```

## Changing the schema

when changing the schema file (`schema.graphqls`) file, make sure to run `gqlgen generate` after.
