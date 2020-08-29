package com.stefanomantini.timestampconsumerservice.api.model;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.*;

@Getter
@Setter
@ToString
@NoArgsConstructor
@AllArgsConstructor
public class TimestampRecord {

    @JsonProperty("event_timestamp")
    private String eventTimestamp;

    @JsonProperty("command_id")
    private String commandId;

}
