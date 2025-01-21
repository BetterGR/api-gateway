@echo off

:: Set environment variables
set API_GATEWAY_PORT=1234
set CLIENT_SECRET=**********
set GRADES_ADDRESS=localhost:50051
set STUDENTS_ADDRESS=localhost:50052
set KEYCLOAK_URL=http://auth.betterGR.org
set REDIRECT_URI=http://localhost:3000/callback

:: Run the API Gateway
go run cmd\api-gateway\main.go
