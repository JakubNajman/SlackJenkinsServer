FROM golang:1.16-alpine

ADD . /go/src/gms
WORKDIR /go/src/gms

RUN go mod download

RUN go build -o main .

CMD [ "/go/src/gms/main" ]