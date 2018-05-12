#!/bin/bash
echo "POST"
curl -X POST -H "Content-Type: application/json" --data '{"ID":1337,"Title":"HelloWorld","Added": "Somedate","Blog":"Blah blah blah Summer\r\n blah blah blah","Instructions": "1. Crack eggs\r\n2. ???\r\n3. Profit"}' http://localhost:8080/api/recipes && echo

echo "GET"
curl -X GET -H "Content-Type: application/json" --data '{"ID":1337}' http://localhost:8080/api/recipes && echo
