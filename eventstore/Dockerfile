FROM ubuntu:14.04.2

ENV ES_VERSION 3.5.0

# EventStore archive needs to have been downloaded to 'data' directory
ADD data/EventStore-OSS-Ubuntu-14.04-v$ES_VERSION.tar.gz /opt/

EXPOSE 2113
EXPOSE 1113

VOLUME /data/db
VOLUME /data/logs

ENV EVENTSTORE_MAX_MEM_TABLE_SIZE 100000
ENV EVENTSTORE_WORKER_THREADS 12

WORKDIR /opt/EventStore-OSS-Ubuntu-14.04-v$ES_VERSION

CMD ./run-node.sh --ext-http-prefixes=http://*:2113/ --ext-ip=0.0.0.0 \
    --db /data/db --log /data/logs --run-projections=all
