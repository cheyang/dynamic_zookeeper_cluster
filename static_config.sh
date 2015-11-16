#!/bin/bash
ZOO_DIR="/opt/zookeeper"
ZOO_CFG="${ZOO_DIR}/conf/zoo.cfg"
for i in `env |grep '^ADDITIONAL_ZOOKEEPER_'|awk -F'=' '{print $1}'|  sort -n -k 3 -t '_'`
do   
     ADDITIONAL_ZOOKEEPER=`eval echo '$'$i`   
     echo ${ADDITIONAL_ZOOKEEPER} >> ${ZOO_CFG}.dynamic    
done
echo "${MYID}" > /tmp/zookeeper/myid
${ZOO_DIR}/bin/zkServer.sh start-foreground
