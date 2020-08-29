package com.stefanomantini.timestampqueryservice.api;

import com.stefanomantini.timestampqueryservice.entity.TimestampRecord;
import com.stefanomantini.timestampqueryservice.repository.TimestampRepository;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.Optional;
import java.util.UUID;

@Slf4j
@RestController
public class TimestampResource {

    // omitted service layer for brevity
    @Autowired TimestampRepository timestampRepository;

    @GetMapping(value = "/api/v1/timestamp/{id}")
    public ResponseEntity<TimestampRecord> getTimestampById(@PathVariable UUID id) {
        Optional<TimestampRecord> timestampRecord = timestampRepository.findById(id);
        if (timestampRecord.isPresent()){
            return ResponseEntity.ok(timestampRecord.get());
        }else{
            return ResponseEntity.notFound().build();
        }
    }

}
