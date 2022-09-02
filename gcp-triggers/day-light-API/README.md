An implementation of the https://sunrise-sunset.org/api

1. deploy the function

gcloud functions deploy DayLight --runtime go116 --trigger-http

2. the deployed function’s URL should be of the form:

https://us-central1-aiggato-upload.cloudfunctions.net/DayLight  ->  https://us-central1-<PROJECT_ID>.cloudfunctions.net/DayLight

3. query URL needs to be of the form:

curl "https://us-central1-aiggato-upload.cloudfunctions.net/DayLight?lat=48.8&lon=9.18&date=2023-01-01" 

-> curl "https://us-central1--<PROJECT_ID>.cloudfunctions.net/DayLight?lat=48.8&lon=9.18&date=2023-01-01"