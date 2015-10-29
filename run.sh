#!/bin/sh

# Output server ID
echo "server id (myid): ${SERVER_ID}"
echo "${SERVER_ID}" > /tmp/zookeeper/myid

HOSTNAME=`hostname`
IPADDRESS=`ip -4 addr show scope global dev eth0 | grep inet | awk '{print \$2}' | cut -d / -f 1`
cd /opt/zookeeper
 
if [ -n "$ZK_URL" ]
then
  set -o pipefail;/opt/zookeeper/bin/zkCli.sh -server $ZK_URL get /zookeeper/config|grep ^server`
  echo "`set -o pipefail;/opt/zookeeper/bin/zkCli.sh -server $ZK_URL get /zookeeper/config|grep ^server`" >> /opt/zookeeper/conf/zoo.cfg.dynamic
  echo "server.$SERVER_ID=$ZK_PEER_URL:observer;$CLIENT_PORT" >> /opt/zookeeper/conf/zoo.cfg.dynamic
  cp /opt/zookeeper/conf/zoo.cfg.dynamic /opt/zookeeper/conf/zoo.cfg.dynamic.org
  ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' /opt/zookeeper/bin/zkServer.sh start
  /opt/zookeeper/bin/zkCli.sh -server $ZK_URL reconfig -add "server.$SERVER_ID=$ZK_PEER_URL:participant;$CLIENT_PORT"
  /opt/zookeeper/bin/zkServer.sh stop
  ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' /opt/zookeeper/bin/zkServer.sh start-foreground  
else
  echo "server.$MYID=$ZK_PEER_URL:participant;$CLIENT_PORT" >> /opt/zookeeper/conf/zoo.cfg.dynamic
  ZOO_LOG_DIR=/var/log ZOO_LOG4J_PROP='INFO,CONSOLE,ROLLINGFILE' /opt/zookeeper/bin/zkServer.sh start-foreground
fi