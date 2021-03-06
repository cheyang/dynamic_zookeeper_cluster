FROM vixns/java8

RUN wget -q -O - http://apache.arvixe.com/zookeeper/zookeeper-3.5.1-alpha/zookeeper-3.5.1-alpha.tar.gz | tar -xzf - -C /opt \
    && mv /opt/zookeeper-3.5.1-alpha /opt/zookeeper \
    && cp /opt/zookeeper/conf/zoo_sample.cfg /opt/zookeeper/conf/zoo.cfg \
    && mkdir -p /tmp/zookeeper

RUN cp /opt/zookeeper/conf/zoo_sample.cfg /opt/zookeeper/conf/zoo.cfg \
    && echo "standaloneEnabled=false" >> /opt/zookeeper/conf/zoo.cfg  \
    && echo "dynamicConfigFile=/opt/zookeeper/conf/zoo.cfg.dynamic" >> /opt/zookeeper/conf/zoo.cfg \
    && sed -ri 's/clientPort=2181/#clientPort=2181/g' /opt/zookeeper/conf/zoo.cfg

ADD . /usr/local/bin/

RUN chmod 0755 /usr/local/bin/*.sh

ENTRYPOINT ["/usr/local/bin/run.sh"]