FROM debian:jessie
MAINTAINER Che Yang <cheyang@163.com>

RUN apt-get update \
 && apt-get -y install git ant openjdk-8-jdk \
 && apt-get clean

RUN wget -q -O - http://apache.arvixe.com/zookeeper/zookeeper-3.5.1-alpha/zookeeper-3.5.1-alpha.tar.gz | tar -xzf - -C /opt \
    && mv /opt/zookeeper-3.5.1-alpha /opt/zookeeper \
    && cp /opt/zookeeper/conf/zoo_sample.cfg /opt/zookeeper/conf/zoo.cfg \
    && mkdir -p /tmp/zookeeper

RUN cp /opt/zookeeper/conf/zoo_sample.cfg /opt/zookeeper/conf/zoo.cfg
RUN echo "standaloneEnabled=false" >> /opt/zookeeper/conf/zoo.cfg
RUN echo "dynamicConfigFile=/opt/zookeeper/conf/zoo.cfg.dynamic" >> /opt/zookeeper/conf/zoo.cfg
ADD zk-init.sh /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/zk-init.sh"]