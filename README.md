# cdk-golang-lambda-deployment

This repo is an example on how to deploy a serverless API on AWS Lambda using aws CDK.

## Lambda
The GO Lambda function is a simple API using `go-chi/chi` to handle http requests.  
The code is located in the [`/lambda/api`](/lambda/api) folder.

## Deployment with CDK
The CDK stack is also simple, it build the binary of the lambda and deploy it with an API Gateway in front of this lambda to access it through HTTP.  
The code is located in the [`/infra-cdk`](/infra-cdk) folder.
