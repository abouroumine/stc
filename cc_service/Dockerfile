FROM golang:alpine as builder

RUN mkdir /cc-build
ADD . /cc-build/
WORKDIR /cc-build
RUN go build -o main

FROM alpine
RUN mkdir /cc
WORKDIR /cc
COPY --from=builder /cc-build/ /cc/

CMD ["./main"]