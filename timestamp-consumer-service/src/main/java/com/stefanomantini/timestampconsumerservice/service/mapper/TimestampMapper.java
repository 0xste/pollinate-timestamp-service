package com.stefanomantini.timestampconsumerservice.service.mapper;

import com.stefanomantini.timestampconsumerservice.api.model.TimestampBO;
import com.stefanomantini.timestampconsumerservice.entity.TimestampEntity;
import java.time.Instant;
import java.util.UUID;
import org.springframework.stereotype.Component;

@Component
public class TimestampMapper {

  public TimestampEntity timestampBoToEntity(TimestampBO tsrb) {
    return TimestampEntity.builder()
        .id(UUID.fromString(tsrb.getCommandId()))
        .eventTimestamp(Instant.parse(tsrb.getEventTimestamp()))
        .build();
  }
}
