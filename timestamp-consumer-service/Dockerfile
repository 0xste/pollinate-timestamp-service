# builder
FROM gradle:6.3.0-jdk13 as builder
RUN mkdir /build
ADD timestamp-consumer-service /build/
WORKDIR /build
RUN ./gradlew build
# service
FROM openjdk:13-jdk-alpine3.10
ARG APP_USER=user
RUN addgroup -S ${APP_USER} && adduser -S ${APP_USER} -G ${APP_USER}
COPY --from=builder build ./app.jar
USER ${APP_USER}:${APP_USER}
ENTRYPOINT java -jar ./app.jar