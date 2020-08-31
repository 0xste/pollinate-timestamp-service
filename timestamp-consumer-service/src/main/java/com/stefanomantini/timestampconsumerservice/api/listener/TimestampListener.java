package com.stefanomantini.timestampconsumerservice.api.listener;

import com.stefanomantini.timestampconsumerservice.api.model.TimestampBO;
import com.stefanomantini.timestampconsumerservice.service.TimestampService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.kafka.listener.ConsumerSeekAware;
import org.springframework.kafka.support.KafkaHeaders;
import org.springframework.messaging.handler.annotation.Header;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class TimestampListener implements ConsumerSeekAware {

  @Autowired TimestampService timestampService;

  @KafkaListener(topics = "${kafka.topic}")
  public void listen(
      TimestampBO timestampBO,
      @Header(KafkaHeaders.RECEIVED_TOPIC) String topic,
      @Header(KafkaHeaders.RECEIVED_PARTITION_ID) Integer partition,
      @Header(KafkaHeaders.OFFSET) Long offset) {
    log.info(
        "received on topic: {} offset: {} partition: {} command_id: {} timestamp: {}",
        topic,
        offset,
        partition,
        timestampBO.getCommandId(),
        timestampBO.getEventTimestamp());
    timestampService.SubmitTimestamp(timestampBO);
  }
}
