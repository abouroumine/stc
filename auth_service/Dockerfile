FROM golang:alpine as builder

RUN mkdir /auth-build
ADD . /auth-build/
WORKDIR /auth-build
RUN go build -o main

FROM alpine
RUN mkdir /auth
WORKDIR /auth
COPY --from=builder /auth-build/ /auth/

CMD ["./main"]