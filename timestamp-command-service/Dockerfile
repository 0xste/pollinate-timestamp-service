FROM alpine:3.12
ARG APP_USER=gonad

RUN addgroup -S ${APP_USER} && adduser -S ${APP_USER} -G ${APP_USER}

WORKDIR /timestamp-command-service

COPY timestamp-command-service timestamp-command-service
COPY config/.default.env config/.default.env

USER ${APP_USER}:${APP_USER}
ENTRYPOINT ["./timestamp-command-service"]
