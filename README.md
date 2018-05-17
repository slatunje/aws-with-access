#   aws-with-access

##  with configuration use cases

The aim of this project is to simplify configuration of aws and allow you to get on with your task either via cli or sdk.
It relies on the following [aws environment variables](https://docs.aws.amazon.com/cli/latest/userguide/cli-environment.html).

### with default environment...

Although not required, the minimum expected environment variables to set are:

```bash
export AWS_ACCESS_KEY_ID=___ACCESSKEY___
export AWS_SECRET_ACCESS_KEY=___SECRETACCESSKEY___
```

Setting the above values are equivalent to:

```bash
export AWS_ACCESS_KEY_ID=___ACCESSKEY___
export AWS_SECRET_ACCESS_KEY=___SECRETACCESSKEY___

export AWS_SESSION_TOKEN=""
export AWS_DEFAULT_REGION=eu-west-1
export AWS_DEFAULT_OUTPUT=json
export AWS_CA_BUNDLE=""
export AWS_PROFILE=default
export AWS_SHARED_CREDENTIALS_FILE=~/.aws/credentials
export AWS_CONFIG_FILE=~/.aws/config
```
Default `with` will always use the above environment values as default values.

### with default credentials...

Default `with` expects to find the default path `~/.aws/credentials`

##  Usage Direct

    with aws s3 ls 

    with aws s3 mb s3://vit-prod-lambda
    
    with aws s3 ls s3://vit-prod-lambda
    
##  Usage Session

- [visit awssudo](https://github.com/JSainsburyPLC/awssudo#usage)

    with \
    --interactive \
    --profile vit-prod 
   
    aws s3 ls
    
    terraform
    
    with \ 
    --interactive \
    --profile vit-dev \
    --switch 
    
    aws s3 ls
    
    exit
    
