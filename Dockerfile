FROM golang

ADD . /go/src/github.com/optonaut/cron-server
WORKDIR /go/src/github.com/optonaut/cron-server

ENV GOPATH /go/src/github.com/optonaut/cron-server/Godeps/_workspace:$GOPATH
RUN CGO_ENABLED=0 go build -o main -v -a -installsuffix nocgo .

EXPOSE 3000
CMD ["./main"]
