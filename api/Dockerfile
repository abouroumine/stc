FROM golang:alpine as builder

RUN mkdir /api-build
ADD . /api-build/
WORKDIR /api-build
RUN go build -o main

FROM alpine
RUN mkdir /api
WORKDIR /api
COPY --from=builder /api-build/ /api/

CMD ["./main"]
