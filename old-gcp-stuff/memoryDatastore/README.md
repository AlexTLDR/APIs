For MySQL

1. Install docker and configure a Dockerfile:

- create a directory called mysql-docker in the project's directory
- copy the sql-scripts directory in the mysql-docker folder
- copy the Dockerfile in the mysql-docker folder

Or use the shared image that I already created and skip to step 3. In the terminal, type -> docker pull alextldr/aiggato-sql:latest 

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

For GCP

1. Create Google Cloud Storage Bucket

- https://cloud.google.com/storage/docs/creating-buckets

2. Generate the service account

- https://cloud.google.com/storage/docs/reference/libraries#client-libraries-install-go for details, and choose “Create Service Account Key”

3. Conduct testing to upload a file using the API

- I used Postman for testing. Configure the POST command with http://localhost:8082/upload
- Under the Body tab, select form-data and under KEY add file-input (the name of the c.FormFile) and under the VALUE, browse for the file to be uploaded

4. Check the uploaded file in the Google bucket

https://console.cloud.google.com/storage 

5. Run the program

- run the program by adding the os variable in the command so the same program can be used on multiple machines/environments 
as os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/alex/Keys/GoogleCloud/aiggato/aiggato-upload-18942db9665f.json") should be removed
from the main.go file.

-> GOOGLE_APPLICATION_CREDENTIALS="/home/alex/Keys/GoogleCloud/aiggato/aiggato-upload-18942db9665f.json" go run . 


