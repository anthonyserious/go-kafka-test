FROM golang:1.6

COPY consumer /go/src/consumer
COPY producer /go/src/producer
COPY citiapi /go/src/citiapi

RUN for i in consumer producer citiapi; do cd /go/src/$i && go get -d -v; done
RUN for i in consumer producer citiapi; do cd /go/src/$i && go install -v; done


