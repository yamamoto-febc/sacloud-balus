FROM golang:1.7.3-alpine
LABEL maintainer="Kazumichi Yamamoto <yamamoto.febc@gmail.com>"

RUN set -x && apk add --no-cache --virtual .build_deps bash git make zip 
RUN go get -u github.com/kardianos/govendor

ADD . /go/src/github.com/yamamoto-febc/sacloud-balus
WORKDIR /go/src/github.com/yamamoto-febc/sacloud-balus
RUN make build

ENTRYPOINT ["bin/sacloud-balus"]
