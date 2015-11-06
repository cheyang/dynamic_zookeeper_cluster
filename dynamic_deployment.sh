#!/bin/bash
ZOO_DIR="/opt/zookeeper"
ZOO_CFG="${ZOO_DIR}/conf/zoo.cfg"

if [ -n $ZK ]; then 
     echo "`bin/zkCli.sh -server $ZK:$MYCLIENT_PORT get /zookeeper/config|grep ^server`" >> ${ZOO_CFG}.dynamic
     echo "server.$MYID=$MYSERVER_URL:observer;$MYCLIENT_PORT" >> ${ZOO_CFG}.dynamic
     cp ${ZOO_CFG}.dynamic  ${ZOO_CFG}.dynamic.org
     echo "${MYID}" > /tmp/zookeeper/myid
     ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' ${ZOO_DIR}/bin/zkServer.sh start
     ${ZOO_DIR}/bin/zkCli.sh -server $ZK:$MYCLIENT_PORT reconfig -add "server.$MYID=$MYSERVER_URL:participant;$MYCLIENT_PORT"
     ${ZOO_DIR}/bin/zkServer.sh stop
     ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' ${ZOO_DIR}/bin/zkServer.sh start-foreground  
else
     echo "server.$MYID=$MYSERVER_URL:participant;$MYCLIENT_PORT" >> ${ZOO_CFG}.dynamic
     echo "${MYID}" > /tmp/zookeeper/myid
     ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' ${ZOO_DIR}/bin/zkServer.sh start-foreground
fi
