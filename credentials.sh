#!/bin/bash

profile=$1
role_arn=$2
role_session_name=${USER} # お好みで
duration_seconds=900 # 900-3600 必要に応じて増減

result=`aws --profile ${profile} sts assume-role --role-arn ${role_arn} --role-session-name ${role_session_name} --duration-seconds ${duration_seconds}`
access_key_id=`echo ${result} | jq -r '.Credentials.AccessKeyId'` 
secret_access_key=`echo ${result} | jq -r '.Credentials.SecretAccessKey'` 
session_token=`echo ${result} | jq -r '.Credentials.SessionToken'` 

echo export AWS_ACCESS_KEY_ID=${access_key_id}
echo export AWS_SECRET_ACCESS_KEY=${secret_access_key}
echo export AWS_SESSION_TOKEN=${session_token}
