FROM golang:alpine as builder

RUN mkdir /ss-build
ADD . /ss-build/
WORKDIR /ss-build
RUN go build -o main

FROM alpine
RUN mkdir /ss
WORKDIR /ss
COPY --from=builder /ss-build/ /ss/

CMD ["./main"]