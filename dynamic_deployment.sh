#!/bin/bash
ZOO_DIR="/opt/zookeeper"
ZOO_CFG="${ZOO_DIR}/conf/zoo.cfg"

if [ -n $ZK_LEADER_URL ]; then 
     echo "`bin/zkCli.sh -server $ZK_LEADER_URL get /zookeeper/config|grep ^server`" >> ${ZOO_CFG}.dynamic
     echo "server.$MY_ID=$MY_SERVER_URL:observer;$MY_CLIENT_PORT" >> ${ZOO_CFG}.dynamic
     cp ${ZOO_CFG}.dynamic  ${ZOO_CFG}.dynamic.org
     echo "${MY_ID}" > /tmp/zookeeper/MY_ID
     ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' ${ZOO_DIR}/bin/zkServer.sh start
     ${ZOO_DIR}/bin/zkCli.sh -server $ZK:$MYCLIENT_PORT reconfig -add "server.$MY_ID=$MY_SERVER_URL:participant;$MY_CLIENT_PORT"
     ${ZOO_DIR}/bin/zkServer.sh stop
     ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' ${ZOO_DIR}/bin/zkServer.sh start-foreground  
else
     echo "server.$MY_ID=$MY_SERVER_URL:participant;$MY_CLIENT_PORT" >> ${ZOO_CFG}.dynamic
     echo "${MY_ID}" > /tmp/zookeeper/MY_ID
     ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' ${ZOO_DIR}/bin/zkServer.sh start-foreground
fi
