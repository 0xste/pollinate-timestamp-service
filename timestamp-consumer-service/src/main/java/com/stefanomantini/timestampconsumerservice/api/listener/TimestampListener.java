package com.stefanomantini.timestampconsumerservice.api.listener;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.google.gson.Gson;
import com.stefanomantini.timestampconsumerservice.api.model.TimestampRecord;
import lombok.extern.slf4j.Slf4j;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.messaging.handler.annotation.Payload;
import org.springframework.stereotype.Component;

import java.sql.Time;
import java.util.Map;

@Slf4j
@Component
public class TimestampListener {

    @KafkaListener(topics = "timestamp.command", groupId = "timestamp-consumer-service")
    public void listen(TimestampRecord timestampRecord) {
        log.info("recieved message on topic msg: {}", timestampRecord.toString());
    }

}
