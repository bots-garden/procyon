#!/bin/bash
curl -X POST -d 'Jane' http://localhost:8080/functions/hello; echo ""
curl -X POST -d 'John' http://localhost:8080/functions/hey; echo ""
curl -X POST -d 'Bob Morane' http://localhost:8080/functions/hi; echo ""
curl -X POST -d 'Sam' http://localhost:8080/functions/yo; echo ""

