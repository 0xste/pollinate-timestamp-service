package com.stefanomantini.timestampqueryservice.repository;

import com.stefanomantini.timestampqueryservice.entity.TimestampRecord;
import java.util.UUID;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface TimestampRepository extends JpaRepository<TimestampRecord, UUID> {}
