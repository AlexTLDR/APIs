A minimal function that responds “Hello, World!” to any incoming HTTP request for GCP.

1. Go to https://console.cloud.google.com/cloud-resource-manager and create a new project

2. Go to https://console.cloud.google.com/functions and enable the Cloud Functions API

3. Install the gcloud CLI tool -> https://cloud.google.com/sdk/docs/ in order to write the functions locally

4. Run the below commands to update and install the Cloud Functions beta (needed in order to use Go):

gcloud components update
gcloud components install beta

5. Point gcloud client to the project:

gcloud config set project <PROJECT_ID>

6. Test that gcloud is logged correctly:

gcloud functions list

7. deploy the helloworld.go function:

gcloud functions deploy HelloWorld --runtime go116 --trigger-http

when asked "Allow unauthenticated invocations of new function [HelloWorld]? (y/N)?" type y so the function can be tested from the https trigger.

8. When successful, there should be an URL under httpsTrigger:

httpsTrigger:
  securityLevel: SECURE_ALWAYS
  url: https://us-central1-<PROJECT_ID>.cloudfunctions.net/HelloWorld
ingressSettings: ALLOW_ALL

9. When opening the URL, a hello world message will be returned.