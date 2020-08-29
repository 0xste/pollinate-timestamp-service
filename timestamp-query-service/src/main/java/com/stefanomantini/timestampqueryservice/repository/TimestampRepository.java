package com.stefanomantini.timestampqueryservice.repository;

import com.stefanomantini.timestampqueryservice.entity.TimestampRecord;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.UUID;

@Repository
public interface TimestampRepository extends JpaRepository<TimestampRecord, UUID> {}
