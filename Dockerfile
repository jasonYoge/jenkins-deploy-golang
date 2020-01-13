FROM golang:alpine

COPY . /go/src

WORKDIR /go/src
RUN go build -o server .
RUN mv /go/src/server /go/bin

EXPOSE 8001

CMD ["server"]