# builder
FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o timestamp-command-service .
# service
FROM alpine:3.12
ARG APP_USER=user
RUN addgroup -S ${APP_USER} && adduser -S ${APP_USER} -G ${APP_USER}
COPY --from=builder /build/timestamp-command-service /timestamp-command-service
COPY --from=builder /build/config/default.env /config/default.env
USER ${APP_USER}:${APP_USER}
ENTRYPOINT ["./timestamp-command-service"]
