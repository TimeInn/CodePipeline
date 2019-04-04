#!/usr/bin/bash
# 
# TimeInn CodePipeline
# 
# @version		1.0.0
# @date			  2019-01-02
# @generator  2019-04-04
# 
# PLEASE DO NOT MODIFY THIS FILE !!
#
PIPE_HOST='https://opendev.uncrash.net/pipe/'
TOKEN=''
PROJECT_NAME='backend'
PROJECT_VERSION='0.2.9'

curl --request POST \
  --url ${PIPE_HOST} \
  --header 'Content-Type: application/x-www-form-urlencoded' \
  --data "token=$TOKEN&project=$PROJECT_NAME&ver=$PROJECT_VERSION"