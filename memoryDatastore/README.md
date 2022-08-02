
1. Install docker and configure a Dockerfile:

- create a directory called mysql-docker in the project's directory
- copy the sql-scripts directory in the mysql-docker folder
- copy the Dockerfile in the mysql-docker folder

2. Create an image from the docker file:

- open a terminal from the mysql-docker folder
- run this command to build the image -> docker build -t aiggato-sql . 

3. Start the container from the image:

- in the same terminal from the mysql-docker folder, run -> docker run -d -p 3306:3306 --name aiggato-sql \-e MYSQL_ROOT_PASSWORD=123456 aiggato-sql

4. In order to restart the container, just run -> docker start aiggato-sql

In order to test the endpoints using curl, install jq or remove the | jq from the end of the command:

curl -X GET http://localhost:8081/api/contacts | jq

curl -X GET http://localhost:8081/api/contacts/2 | jq 

curl -X POST http://localhost:8081/api/contacts -d '{"id": "2","user_name": "TestID2","mail": "test@id2.com"}' 

curl -X PUT http://localhost:8081/api/contacts/4 -d '{"id":"2","user_name":"TestID4Mod","mail":"test@id4.com"}' 

curl -X DELETE http://localhost:8081/api/contacts/4 | jq 

