# Location-Tracker-Backend

This application will serve the location information for a person.

## Overview

This project uses a combination of AWS SAM (Serverless Application Model), Go (backend language), and make (build tool).

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
make test-api
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
# Run this command once your AWS credentials are set and this will perform and update on the API in AWS.
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