FROM golang:alpine as builder

RUN mkdir /db-build
ADD . /db-build/
WORKDIR /db-build
RUN go build -o main

FROM alpine
RUN mkdir /db
WORKDIR /db
COPY --from=builder /db-build/ /db/

CMD ["./main"]