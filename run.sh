#!/bin/bash
#env:
#export DEPLOY_MODE=STATIC or DEPLOY_MODE=DYNAMIC
#export ADDITIONAL_ZOOKEEPER_1=server.1=192.168.94.135:2888:3888:observer;2181
#export MYID=6
#export PORT=2888:3888
#export CLIENT_PORT=2181
#export ZK=192.168.1.1
#export MYSERVER_URL=192.168.2.2:2888:3888

if [ $DEPLOY_MODE = STATIC ] ; then  
   sh static_deployment.sh
elif [ $STATUS = DYNAMIC ]; then  
   sh dynamic_deployment.sh 
else
   echo "The environment variable DEPLOY_MODE does not exist or is not correct."
   exit
fi