This is a simple API for Go to write files in Google Cloud buckets. In order to test it, follow the instructions.

1. Create Google Cloud Storage Bucket

- https://cloud.google.com/storage/docs/creating-buckets

2. Generate the service account

- https://cloud.google.com/storage/docs/reference/libraries#client-libraries-install-go for details, and choose “Create Service Account Key”

3. Conduct testing to upload a file using the API

- I used Postman for testing. Configure the POST command with http://localhost:8082/upload (check in the code for r.Run(":80**")) or in the terminal for
something like [GIN-debug] Listening and serving HTTP on :8082 to get the correct port.
- Under the Body tab, select form-data and under KEY add file-input (the name of the c.FormFile) and under the VALUE, browse for the file to be uploaded

4. Check the uploaded file in the Google bucket

https://console.cloud.google.com/storage 

5. Run the program

- run the program by adding the os variable in the command so the same program can be used on multiple machines/environments 
-> GOOGLE_APPLICATION_CREDENTIALS="/home/alex/Keys/GoogleCloud/aiggato/aiggato-upload-18942db9665f.json" go run main.go