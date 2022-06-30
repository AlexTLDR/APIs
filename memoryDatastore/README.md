In order to test the endpoints using curl, install jq or remove the | jq from the end of the command:

curl -X GET http://localhost:8081/api/contacts | jq
curl -X GET http://localhost:8081/api/contacts/2 | jq . #to retrieve a specific article by id.
curl -X POST http://localhost:8081/api/contacts -d '{"ID": "4","Name": "TestID4","Mail": "test@id4.com"}' -H "Content-Type: application/json" | jq 
curl -X PUT http://localhost:8081/api/contacts/4 -d '{"ID":"4","Name":"TestID4Mod","Mail":"test@id4.com"}' -H "Content-Type: application/json" | jq 
curl -X DELETE http://localhost:8081/api/contacts/4 | jq 