FROM ubuntu:14.04.2

ENV ES_VERSION 3.5.0

RUN DEBIAN_FRONTEND=noninteractive apt-get -y install wget \
&& cd /opt \
&& wget --no-check-certificate https://download.geteventstore.com/binaries/EventStore-OSS-Ubuntu-14.04-v$ES_VERSION.tar.gz \
&& tar xzf EventStore-OSS-Ubuntu-14.04-v$ES_VERSION.tar.gz \
&& chown -R eventstore:eventstore EventStore-OSS-Ubuntu-14.04-v$ES_VERSION \
&& rm EventStore-OSS-Ubuntu-14.04-v$ES_VERSION.tar.gz \
&& cd /

EXPOSE 2113
EXPOSE 1113

VOLUME /data/db
VOLUME /data/logs

ENV EVENTSTORE_MAX_MEM_TABLE_SIZE 100000
ENV EVENTSTORE_WORKER_THREADS 12

WORKDIR /opt/EventStore-OSS-Ubuntu-14.04-v$ES_VERSION

CMD ./run-node.sh --ext-http-prefixes=http://*:2113/ --ext-ip=0.0.0.0 \
    --db /data/db --log /data/logs --run-projections=all
