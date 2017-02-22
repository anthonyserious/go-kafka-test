FROM golang:1.6

COPY consumer /go/src/consumer
COPY producer /go/src/producer

RUN for i in consumer producer; do cd /go/src/$i && go get -d -v; done
RUN for i in consumer producer; do cd /go/src/$i && go install -v; done


