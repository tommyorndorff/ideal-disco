---
marp: true
theme: gaia
_class: invert
---

# Build Networking

---

<!-- paginate: true -->

## Create the stack

Create the stack...
```sh
aws cloudformation create-stack --stack-name "torndorff-vpc" --template-body file://src/01-base_vpc.yml --parameters ParameterKey=EnvironmentName,ParameterValue=torndorff
```

Describe the events...
```sh
aws cloudformation describe-stack-events --stack-name "torndorff-vpc"
```

Describe the stack...
```sh
aws cloudformation describe-stacks --stack-name "torndorff-vpc"
```

---

## Create an ec2 instance

Create an ssh key & set the permissions
```sh
aws ec2 create-key-pair --key-name `whoami` \
  --query "KeyMaterial" \
  --output text > `whoami`.pem

chmod 0400 `whoami`.pem
```

Prove you did it
```sh
aws describe-key-pairs --key-name `whoami`
```

---

## Launch an ec2 instance via Cloudformation
https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/example-templates-autoscaling.html
https://s3.amazonaws.com/cloudformation-templates-us-east-1/AutoScalingMultiAZWithNotifications.template

This template will make the following:
- 

Deploy the template, either through CLI or Console:
```sh
aws cloudformation create-stack --stack-name "torndorff-webserver-2" \
  --template-body file://src/02-deploy_webserver.yml \
  --parameters ParameterKey=InstanceType,ParameterValue=t2.small \
  ParameterKey=KeyName,ParameterValue=torndorff \
  ParameterKey=OperatorEMail,ParameterValue=torndorff@expediagroup.com \
  ParameterKey=SSHLocation,ParameterValue=`curl ifconfig.me`/32 \
  ParameterKey=Subnets,ParameterValue=subnet-03e51f5daef0e047d\\,subnet-09f4438f7111c4e56 \
  ParameterKey=VpcId,ParameterValue=vpc-0c8a04827d1eea914
```

After creation, we can get the URL as stored in the output:
```sh
aws cloudformation describe-stacks --stack-name torndorff-webserver --query "Stacks[0].Outputs[?OutputKey=='URL'].Outp
utValue" --output text
```

## Examine Launch Config

## Examine Autoscaling

## Exercise Autoscaling

We'll use a golang performance testing tool called `k6` to generate traffic
```sh
$ brew install k6
$ k6 run --vus 5 --stage 3m:10,5m:10,10m:35,1m30s:0 src/script.js
```

To use docker (if you dont have brew):
```sh
$ docker pull loadimpact/k6
$ docker run -i loadimpact/k6 run - <src/script.js
```

Or on Windows:
```sh
docker pull loadimpact/k6
cat src/script.js | docker run -i loadimpact/k6 run -
```


