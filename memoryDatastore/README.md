In order to test the endpoints using curl, install jq or remove the | jq from the end of the command:

curl -X GET http://localhost:8081/api/contacts | jq

curl -X GET http://localhost:8081/api/contacts/2 | jq 

curl -X POST http://localhost:8081/api/contacts -d '{"ID": "4","Name": "TestID4","Mail": "test@id4.com"}' -H "Content-Type: application/json" | jq 

curl -X PUT http://localhost:8081/api/contacts/4 -d '{"ID":"4","Name":"TestID4Mod","Mail":"test@id4.com"}' -H "Content-Type: application/json" | jq 

curl -X DELETE http://localhost:8081/api/contacts/4 | jq 

In order to test SQL querys, I did the below steps:

1. Install docker and configure a docker image:

- create a directory and a docker-compose.yml file

mkdir db-docker
cd db-docker
touch docker-compose.yml

- add the below in the docker-compose.yml file (image: to add the desired version of mysql -> https://hub.docker.com/_/mysql)

version: '3'

services:

  mysql-development:
    image: mysql:8.0.29
    environment:
      MYSQL_ROOT_PASSWORD: 123456
      MYSQL_DATABASE: testapp
    ports:
      - "3308:3306"

- after creatung the .yml file, run the below command. The command pulls the docker image and then it runs the containert

docker-compose up

- to check the status (if the container is running)

docker-compose ps
