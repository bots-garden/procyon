#!/bin/bash
curl -X POST -d '{"FirstName": "Jane", "LastName": "Doe"}' http://localhost:8082
curl -X POST -d '{"FirstName": "John", "LastName": "Doe"}' http://localhost:8083
curl -X POST -d 'Bob Morane' http://localhost:8084
curl -X POST -d '{"FirstName": "John", "LastName": "Doe"}' http://localhost:8085
