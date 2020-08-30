package com.stefanomantini.timestampconsumerservice.entity;

import java.time.Instant;
import java.util.UUID;
import javax.persistence.*;
import lombok.*;
import org.hibernate.annotations.CreationTimestamp;

@Entity
@NoArgsConstructor
@AllArgsConstructor
@Data
@Builder
@Table(name = "TIMESTAMP_RECORDS", schema = "public")
public class TimestampEntity {

  @Id
  @NonNull
  @Column(name = "ID")
  private UUID id;

  @NonNull
  @Column(name = "EVENT_TIMESTAMP")
  private Instant eventTimestamp;

  @CreationTimestamp
  @Column(name = "CREATED_AT")
  private Instant createdAt;
}
