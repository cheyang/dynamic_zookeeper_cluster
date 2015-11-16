FROM vixns/java8
MAINTAINER Che Yang <cheyang@163.com>


RUN wget -q -O - http://apache.arvixe.com/zookeeper/zookeeper-3.5.1-alpha/zookeeper-3.5.1-alpha.tar.gz | tar -xzf - -C /opt \
    && mv /opt/zookeeper-3.5.1-alpha /opt/zookeeper \
    && cp /opt/zookeeper/conf/zoo_sample.cfg /opt/zookeeper/conf/zoo.cfg \
    && mkdir -p /tmp/zookeeper



RUN cp /opt/zookeeper/conf/zoo_sample.cfg /opt/zookeeper/conf/zoo.cfg
RUN echo "standaloneEnabled=false" >> /opt/zookeeper/conf/zoo.cfg
RUN echo "dynamicConfigFile=/opt/zookeeper/conf/zoo.cfg.dynamic" >> /opt/zookeeper/conf/zoo.cfg

ADD dynamic_config.sh /usr/local/bin/
ADD static_config.sh /usr/local/bin/
ADD run.sh /usr/local/bin/
ENTRYPOINT ["/usr/local/bin/run.sh"]