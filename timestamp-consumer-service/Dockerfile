# builder
FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o timestamp-consumer-service .
# service
FROM alpine:3.12
ARG APP_USER=user
RUN addgroup -S ${APP_USER} && adduser -S ${APP_USER} -G ${APP_USER}
COPY --from=builder /build/timestamp-consumer-service /timestamp-consumer-service
COPY --from=builder /build/config/default.env /config/default.env
USER ${APP_USER}:${APP_USER}
ENTRYPOINT ["./timestamp-consumer-service"]
