#!/bin/bash
#env:
#export CONFIG_MODE=STATIC or CONFIG_MODE=DYNAMIC
#export ADDITIONAL_ZOOKEEPER_1=server.1=192.168.94.135:2888:3888:observer;2181
#export MYID=6
#export PORT=2888:3888
#export CLIENT_PORT=2181
#export ZK=192.168.1.1
#export MYSERVER_URL=192.168.2.2:2888:3888

if [ x$CONFIG_MODE = x ] ; then
   echo "The environment variable CONFIG_MODE does not exist or is not correct."
   exit
fi
if [ $CONFIG_MODE = STATIC ] ; then
    static_config.sh
elif [ $CONFIG_MODE = DYNAMIC ]; then
    dynamic_config.sh
else
   echo "The environment variable CONFIG_MODE does not correct."
   exit
fi