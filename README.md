# Location-Tracker-Backend

This application will serve the location information for a person.

## Overview

This project uses a combination of AWS SAM (Serverless Application Model), Go (backend language), and make (build tool).

When you deploy with the command `bash build.sh`, the script will:

- compile the Go application for Linux (Lambda uses a Linux OS)
- create an S3 bucket if it doesn't exist: https://console.aws.amazon.com/s3/buckets/locationtracker-989414517243-sam
- copy the latest version of the code to the S3 bucket with a random filename
- perform a CreateStack or UpdateStack operation on the CloudFormation stack for the application: https://console.aws.amazon.com/cloudformation
- output the URL where you can access the application: https://ozozzj4qhd.execute-api.us-east-1.amazonaws.com/Prod/api/location-history

Here is an overview of the files and folders:

- build.sh - performs all the steps up to deploy the code to AWS, the URL is public so anyone on the internet can view the API information.
- src/app/api/main.go - the main Go file
- src/app/api/Gopkg.toml - the file that is primarily hand-edited to keep track of the libraries you're using in Go (https://golang.github.io/dep/docs/Gopkg.toml.html)
- src/app/api/Gopkg.lock - the file that keeps track of the versions of the dependencies you're using, it should not be hand edited (https://golang.github.io/dep/docs/Gopkg.lock.html)
- src/app/api/Makefile - a very simple build automation file to save on typing (https://en.wikipedia.org/wiki/Makefile), all the commands below start with: `make`
- src/app/api/template.json - the serverless template used by SAM to determine how your code will be set up in AWS as well as the endpoints that are available.
- src/app/api/pkg/core.go - code that controls how the app handles specific routes (these must be keep in sync with template.json)

### Environment Setup

You need Go, the AWS SAM toolkit, and docker installed.

```bash
# Install Go.
brew install go
go --version

# Install Dep (dependency manager for Go).
brew install dep

# Install AWS CLI and AWS SAM.
brew install awscli
brew tap aws/tap
brew install aws-sam-cli
sam --version
```

You should install Docker for Mac: https://docs.docker.com/docker-for-mac/install/

You should also set your `GOPATH` to the location of the root of the repository:

```bash
# Example of how to set your GOPATH.
export GOPATH=/Users/gencavjj/repos/Location-Tracker-Backend
```

## Testing Locally

This uses AWS SAM and make to run the commands:

```bash
cd src/app/api

# Test the API locally (after doing a build).
make run
```

You can also use these commands:

```bash
# Run the test suite.
make test

# Build for linux (not actually used for anything).
make build

# Remove the linux build from above.
make clean
```

### Deploying to Production

You need to login to the AWS console and generate API keys for your user.

```bash
# Run this command to configure AWS credentials on your system.
aws configure

# Enter in these fields:
#AWS Access Key ID [None]: (your Access key ID)
#AWS Secret Access Key [None]: (your Secret access key)
#Default region name [None]: us-east-1
#Default output format [None]: json 

# Run this command to ensure it's set up properly.
aws sts get-caller-identity
```

You credentials will be saved to: ~/.aws/credentials and ~/.aws/config

```bash
# Run this command once your AWS credentials are set and this will perform an update on the API in AWS.
bash build.sh
```

### Testing Production

This is a sample command to hit the API using curl:

```bash
curl -X GET https://ozozzj4qhd.execute-api.us-east-1.amazonaws.com/Prod/api/location-history -H "Accept: application/json"
```

## Steps Used to Create this Repo

```bash
# Create the src dir.
mkdir src
cd src

# Create the initial same package.
sam init --runtime go1.x --name api --output-dir app

# Moved a few things around to get the structure to work.
```

## Reference

These are helpful articles:
- AWS SAM Quickstart: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-quick-start.html
- AWS SAM Install for MacOS: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install-mac.html