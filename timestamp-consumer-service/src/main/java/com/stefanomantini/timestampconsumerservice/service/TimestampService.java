package com.stefanomantini.timestampconsumerservice.service;

import com.stefanomantini.timestampconsumerservice.api.model.TimestampBO;
import com.stefanomantini.timestampconsumerservice.entity.TimestampEntity;
import com.stefanomantini.timestampconsumerservice.repository.TimestampRepository;
import com.stefanomantini.timestampconsumerservice.service.mapper.TimestampMapper;
import java.util.UUID;
import javax.transaction.Transactional;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Slf4j
@Service
@Transactional
public class TimestampService {

  @Autowired TimestampRepository timestampRepository;

  @Autowired TimestampMapper timestampMapper;

  public void SubmitTimestamp(TimestampBO tsrb) {
    // idempotency check
    if (timestampRepository.existsById(UUID.fromString(tsrb.getCommandId()))) {
      log.info("record {} exists, skipping", tsrb.getCommandId());
      return;
    }
    // save the record
    TimestampEntity saved =
        timestampRepository.saveAndFlush(timestampMapper.timestampBoToEntity(tsrb));

    log.info(tsrb.getCommandId());
    // retrieve the saved record
    TimestampEntity retrieved = timestampRepository.getOne(saved.getId());

    log.info(
        "record id {} saved successfully at {} for timestamp {}",
        retrieved.getId(),
        retrieved.getCreatedAt(),
        retrieved.getEventTimestamp());
  }
}
