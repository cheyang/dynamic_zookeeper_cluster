#!/bin/sh

# Output server ID
echo "server id (myid): ${SERVER_ID}"
echo "${SERVER_ID}" > /tmp/zookeeper/myid

HOSTNAME=`hostname`
IPADDRESS=`ip -4 addr show scope global dev eth0 | grep inet | awk '{print \$2}' | cut -d / -f 1`
cd /tmp/zookeeper
 
if [ -n "$ZK_URL" ]
then
  echo "`bin/zkCli.sh -server $ZK_URL get /zookeeper/config|grep ^server`" >> /tmp/zookeeper/conf/zoo.cfg.dynamic
  echo "server.$SERVER_ID=$ZOOKEEPER_ADVERTISE_URL:observer;$CLIENT_PORT" >> /tmp/zookeeper/conf/zoo.cfg.dynamic
  cp /tmp/zookeeper/conf/zoo.cfg.dynamic /tmp/zookeeper/conf/zoo.cfg.dynamic.org
  /tmp/zookeeper/bin/zkServer-initialize.sh --force --myid=$MYID
  ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' /tmp/zookeeper/bin/zkServer.sh start
  /tmp/zookeeper/bin/zkCli.sh -server $ZK_URL reconfig -add "server.$SERVER_ID=$ZOOKEEPER_ADVERTISE_URL:participant;$CLIENT_PORT"
  /tmp/zookeeper/bin/zkServer.sh stop
  ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' /tmp/zookeeper/bin/zkServer.sh start-foreground  
else
  echo "server.$MYID=$ZOOKEEPER_ADVERTISE_URL:participant;$CLIENT_PORT" >> /tmp/zookeeper/conf/zoo.cfg.dynamic
  /tmp/zookeeper/bin/zkServer-initialize.sh --force --myid=$SERVER_ID
  ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' /tmp/zookeeper/bin/zkServer.sh start-foreground
fi