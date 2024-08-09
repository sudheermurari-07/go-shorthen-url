# multistage docker build. This redices the size of the final docker image.
# stage 1 to build the app
FROM golang:latest as builder

RUN mkdir /build 

ADD . /build/

WORKDIR /build 

RUN go build -o main .

# stage 2 deploys the app built in stage 1
FROM alpine:latest

RUN adduser -S -D -H -h /api appuser

USER appuser

COPY . /api

COPY --from=builder /build/main /api/

WORKDIR /api

EXPOSE 3000

CMD ["./main"]