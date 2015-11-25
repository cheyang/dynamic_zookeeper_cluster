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
    -e ADDITIONAL_ZOOKEEPER_1=server.1=0.0.0.0:2888:3888:participant\;2181 \
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
   -e ADDITIONAL_ZOOKEEPER_2=server.2=0.0.0.0:2889:3889:participant\;2182 \
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
    -e ADDITIONAL_ZOOKEEPER_3=server.3=0.0.0.0:2890:3890:participant\;2183 \
    -e CONFIG_MODE=STATIC  \
    -e MYID=3 \
    zookeeper 
    ```
   


####DYNAMIC MODE
- To start the image in dynamic mode you need to specify a couple of environment variables for the container.
   
|Environment Variable|Description|example|
|:-----------|:---------------|:------------------ |
|ZK_URL| 	The url of one of  the zookeeper cluster|localhost:2181|
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
   -e MYSERVER_URL=0.0.0.0:2891:3891 \
   -e MYCLIENT_PORT=2184 \
   -e CONFIG_MODE=DYNAMIC  \
   -e MYID=4 \
   zookeeper
    
   docker run -d \
   --name=zk5 \
   --net=host \
   -v /data/zk5:/tmp/zookeeper \
   -e MYSERVER_URL=0.0.0.0:2892:3892 \
   -e MYCLIENT_PORT=2185 \
   -e CONFIG_MODE=DYNAMIC  \
   -e MYID=5 \
   zookeeper
   ```
   
  





