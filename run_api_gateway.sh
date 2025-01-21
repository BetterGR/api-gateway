#!/bin/bash

# Set environment variables
export API_GATEWAY_PORT=1234
export CLIENT_SECRET=**********
export KEYCLOAK_URL=http://auth.betterGR.org
export GRADES_ADDRESS=localhost:50051
export STUDENTS_ADDRESS=localhost:50052
export HOMEWORK_ADDRESS=localhost:50053
export COURSES_ADDRESS=localhost:50054
export STAFF_ADDRESS=localhost:50055


# Run the API Gateway
go run cmd/api-gateway/main.go