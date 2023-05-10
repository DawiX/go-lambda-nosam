# AWS Lambda in Go without SAM

AWS Lambda written in Go, that is locally executable without the need to use SAM or Serverless Framework.

## Building and Deploying

- Run `./build.sh` script to build `main.zip`.
- Deal with the lambda payload zip file as you see fit (upload to S3, use for deployment directly...)
- Deploy lambda using your favourite IaC tooling (Terraform/Terragrunt, CFN, ...)

## Running locally

- Tweak `main_test.go` to provide desired JSON payload for the Lambda
- Tweak `run_local.sh` to provide AWS profile and region
- Run `./run_local.sh` to invoke lambda locally
