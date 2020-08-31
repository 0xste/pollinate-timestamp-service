package com.stefanomantini.timestampqueryservice.entity;

import com.fasterxml.jackson.annotation.JsonProperty;
import java.time.LocalDateTime;
import java.util.UUID;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import lombok.*;

@Entity
@NoArgsConstructor
@AllArgsConstructor
@Data
@Table(name = "TIMESTAMP_RECORDS", schema = "public")
public class TimestampRecord {

  @Id
  @NonNull
  @Column(name = "ID")
  @JsonProperty("id")
  private UUID id;

  @NonNull
  @Column(name = "EVENT_TIMESTAMP")
  @JsonProperty("event_timestamp")
  private LocalDateTime eventTimestamp;

  @NonNull
  @Column(name = "CREATED_AT")
  @JsonProperty("created_at")
  private LocalDateTime createdAt;
}
