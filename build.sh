#!/bin/sh

# Ensure the script returns an exit code on failure.
set -e

# Build the application and include the API in the binary.
cd src/app/api
GOOS=linux go build
zip $GOPATH/handler.zip ./api

# Make an S3 bucket.
aws s3 mb s3://locationtracker-$(aws sts get-caller-identity \
--query 'Account' --output text)-sam

# Copy the code to S3.
aws cloudformation package \
    --template-file $GOPATH/src/app/api/template.json \
    --s3-bucket locationtracker-$(aws sts get-caller-identity \
    --query 'Account' --output text)-sam \
    --output-template-file $GOPATH/packaged-template.yml

# Deploy the API gateway and lambda.
aws cloudformation deploy \
    --template-file $GOPATH/packaged-template.yml \
    --stack-name api \
    --capabilities CAPABILITY_IAM

# Output the outputs.
aws cloudformation describe-stacks --stack-name api --query 'Stacks[0].Outputs'