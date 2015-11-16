## dynamic_zookeeper_cluster
- Instructions:  
    A ZooKeeper docker image built on **ZooKeeper 3.5.1-alpha**  supports both static and [dynamic configuration](http://zookeeper.apache.org/doc/trunk/zookeeperReconfig.html)
    The docker image supports the following ZooKeeper modes:
       - **STATIC mode**
       - **DYNAMIC mode**
- Environment:  
    OS :Red Hat Enterprise Linux Server release 6.6 (Santiago)   
    Kernel: 2.6.32-573.1.1.el6.x86_64  
    docker:version 1.4.1  
    ZooKeeper: ZooKeeper 3.5.1-alpha    
- Install docker  
Follow the [installation guide] (https://docs.docker.com/installation/) to install docker in your system
- Install docker image
 ```bash
 docker build -t zookeeper .
  ```
####STATIC MODE
   - To start the image in static mode you need to specify a couple of environment variables for the container.
   
|Environment Variable|Description|example|
|:---------|:-------------|:-----------:|
|MYID| 	The id of the server|MYID=1|
|ADDITIONAL_ZOOKEEPER| 	Zookeeper service USES the IP and port|ADDITIONAL_ZOOKEEPER_1=server.1=  localhost:2888:3888:participant\;2181|
|CONFIG_MODE|Create a zookeeper mode|CONFIG_MODE=STATIC|
   - Start ZooKeeper Cluster with 3 docker containers  
    container zk1 is listening on localhost: 2181£¬2888£¬3888  
    container zk2 is listening on localhost: 2182£¬2889£¬3889  
    container zk3 is listening on localhost: 2183£¬2890£¬3890  
    ```txt
    docker run -d  \
     --name=zk1  \
    --net=host  \
    -v /data/zk1:/tmp/zookeeper \
    -e ADDITIONAL_ZOOKEEPER_1=server.1=localhost:2888:3888:participant\;2181 \
    -e ADDITIONAL_ZOOKEEPER_2=server.2=localhost:2889:3889:participant\;2182 \
    -e ADDITIONAL_ZOOKEEPER_3=server.3=localhost:2890:3890:participant\;2183 \
    -e CONFIG_MODE=STATIC   \
    -e MYID=1  \
    zookeeper

   docker run -d \
   --name=zk2 \
   --net=host \
   -v /data/zk2:/tmp/zookeeper \
   -e ADDITIONAL_ZOOKEEPER_1=server.1=localhost:2888:3888:participant\;2181 \
   -e ADDITIONAL_ZOOKEEPER_2=server.2=localhost:2889:3889:participant\;2182 \
   -e ADDITIONAL_ZOOKEEPER_3=server.3=localhost:2890:3890:participant\;2183 \
   -e CONFIG_MODE=STATIC  \
    -e MYID=2 \
    zookeeper 
    
    docker run -d \
    --name=zk3 \
    --net=host \
    -v /data/zk3:/tmp/zookeeper \
    -e ADDITIONAL_ZOOKEEPER_1=server.1=localhost:2888:3888:participant\;2181 \
    -e ADDITIONAL_ZOOKEEPER_2=server.2=localhost:2889:3889:participant\;2182 \
    -e ADDITIONAL_ZOOKEEPER_3=server.3=localhost:2890:3890:participant\;2183 \
    -e CONFIG_MODE=STATIC  \
    -e MYID=3 \
    zookeeper 
    ```
   - Check the ZooKeeper containers' status
   ```bash
   [root@SIJR34APMXP-001 ~]# docker ps
   ```
   ```txt
   CONTAINER ID IMAGE            COMMAND              CREATED  STATUS  PORTS NAMES
11a2c6c40720 zookeeper:latest "/usr/local/bin/run. 5 days ago Up 5 days      zk3
3cdd635e4cb3 zookeeper:latest "/usr/local/bin/run. 5 days ago Up 5 days      zk2
7ddf3a794eb7 zookeeper:latest "/usr/local/bin/run. 5 days ago Up 5 days      zk1
   ```

 
- Test ZooKeeper Cluster(zk1,zk2,zk3)
   - Create a new znode in cluster member zk1
      - To get the container ID
      ```bash
      docker inspect zk1|grep Id|awk -F '"' '{print $4}'
      ```
      - Log on to the container
      ```bash
      lxc-attach --name 7ddf3a794eb74e1ecf55a5a76c2bd0b3a2b1b06670504b6
      ```
      - inside container zk1,Create a new znode 
      ```bash
       /opt/zookeeper/bin/zkCli.sh -server 127.0.0.1:2181
      [zk: 127.0.0.1:2181(CONNECTED) 0]create /test1 abc
      Connecting to 127.0.0.1:2181
      Created /test
     ```  
   - Check the znode in cluster member zk2
      - To get the container ID
      ```bash
      docker inspect zk2|grep Id|awk -F '"' '{print $4}'
      ```
      - Log on to the container
      ```bash
      lxc-attach --name 3cdd635e4cb3708e167cc3c478c9322d6fddc4547a19d8de
      ```
      - inside container zk2 
    ```txt
    /opt/zookeeper/bin/zkCli.sh -server 127.0.0.1:2182
    [zk: 127.0.0.1:2182(CONNECTED) 0] get /test1  
    abc
    ```
   - Check the znode in cluster member zk3
      - To get the container ID
      ```bash
      docker inspect zk3|grep Id|awk -F '"' '{print $4}'
      ```
      - Log on to the container
      ```bash
      lxc-attach --name 11a2c6c40720842bcc9b3a011a75d05db920546fc95f65
      ```
      - inside container zk3 
      ```txt
      /opt/zookeeper/bin/zkCli.sh -server 127.0.0.1:2183
      [zk: 127.0.0.1:2183(CONNECTED) 0] get /test1  
       abc
      ```

####DYNAMIC MODE
- To start the image in dynamic mode you need to specify a couple of environment variables for the container.
   
|Environment Variable|Description|example|
|:-----------|:---------------|:------------------ |
|ZK_URL| 	The id of the server|MYID=1|
|MYSERVER_URL|Zookeeper service using IP and port|MYSERVER_URL=localhost:2891:3891|
|MYCLIENT_PORT|Clients use port|MYCLIENT_PORT=2184|
|CONFIG_MODE|Create a zookeeper mode|CONFIG_MODE=DYNAMIC|
|MYID|The id of the server|MYID=4|
   -  Start ZooKeeper Cluster with 2 docker containers  
    container zk4 is listening on localhost: 2184£¬2891£¬3891  
    container zk5 is listening on localhost: 2185£¬2892£¬3892  
   ```txt
   docker run -d \
   --name=zk4 \
   --net=host \
   -v /data/zk4:/tmp/zookeeper \
   -e ZK_URL=127.0.0.1:2181 \(Known cluster of any member of the IP and port)
   -e MYSERVER_URL=localhost:2891:3891 \
   -e MYCLIENT_PORT=2184 \
   -e CONFIG_MODE=DYNAMIC  \
   -e MYID=4 \
   zookeeper
    
   docker run -d \
   --name=zk5 \
   --net=host \
   -v /data/zk5:/tmp/zookeeper \
   -e MYSERVER_URL=localhost:2892:3892 \
   -e MYCLIENT_PORT=2185 \
   -e CONFIG_MODE=DYNAMIC  \
   -e MYID=5 \
   zookeeper
   ```
   - Check the ZooKeeper containers' status  
   ```bash
   [root@SIJR34APMXP-001 ~]# docker ps
   ```
   ```txt
   CONTAINER ID IMAGE            COMMAND              CREATED  STATUS  PORTS NAMES
018ae649dc89 zookeeper:latest "/usr/local/bin/run. 5 days ago Up 5 days      zk5
3a69528c688b zookeeper:latest "/usr/local/bin/run. 5 days ago Up 5 days      zk4
11a2c6c40720 zookeeper:latest "/usr/local/bin/run. 5 days ago Up 5 days      zk3
3cdd635e4cb3 zookeeper:latest "/usr/local/bin/run. 5 days ago Up 5 days      zk2
7ddf3a794eb7 zookeeper:latest "/usr/local/bin/run. 5 days ago Up 5 days      zk1
   ```


  


- Test zk4 and zk5
   - Check the znode in cluster member zk4
      - To get the container ID
      ```bash
      docker inspect zk4|grep Id|awk -F '"' '{print $4}'
      ```
      - Log on to the container
      ```bash
      lxc-attach --name 3a69528c688bb262631fc3e32d3e252e7018f43011d722a908
      ```
      - inside container zk4 
    ```txt
    /opt/zookeeper/bin/zkCli.sh -server 127.0.0.1:2184
    [zk: 127.0.0.1:2184(CONNECTED) 0] get /test1  
    abc
    ```
   - Check the  zk5
      - To get the container ID
      ```bash
      docker inspect zk5|grep Id|awk -F '"' '{print $4}'
      ```
      - Log on to the container
      ```bash
      lxc-attach --name 018ae649dc89ebebea17004ee7389d267c889ee9dbcfd
      ```
      - inside container zk5 
      ```txt
      /opt/zookeeper/bin/zkCli.sh -server 127.0.0.1:2183
      [zk: 127.0.0.1:2183(CONNECTED) 0] get /test1  
       Node does not exist: /test1
      ``` 
      *note*:By verifying zk5 is independent of the zookeeper cluster leader
    
- Summary  
  Started three zookeeper (zk1, zk2, zk3) form a zookeeper cluster, and then the fourth zookeeper£¨zk4) through any a zookeeper static IP and port, dynamic false join zookeeper cluster, and the fifth zookeeper(zk5) because don't know the static zookeeper information, dynamically created a separate cluster 
  




