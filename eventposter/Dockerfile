FROM golang:1.6-alpine

RUN apk --update add --no-cache git nodejs \
 && npm install --global gulp-cli

USER nobody

ENV GOPATH=/go GOBIN=$GOPATH/bin PATH=$PATH:/go/bin

VOLUME ["/github.com/mramshaw/go-docker-eventstore/poster", "/go/pkg", "/go/bin"]

WORKDIR  /github.com/mramshaw/go-docker-eventstore/poster

CMD ["gulp"]
