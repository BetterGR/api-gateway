#!/bin/bash

# Set environment variables
export API_GATEWAY_PORT=1234
export CLIENT_SECRET=**********
export KEYCLOAK_URL=http://auth.BetterGR.org

# Run the API Gateway
go run src/main.go