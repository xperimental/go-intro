FROM golang:1.10
LABEL maintainer="Robert Jacob <xperimental@solidproject.de>"

RUN go get golang.org/x/tools/cmd/present \
 && mkdir /data

WORKDIR /data
COPY . /data

EXPOSE 3999
CMD [ "/go/bin/present", "-http", ":3999", "-orighost", "localhost" ]
