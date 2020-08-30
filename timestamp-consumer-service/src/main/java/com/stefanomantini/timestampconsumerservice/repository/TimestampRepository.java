package com.stefanomantini.timestampconsumerservice.repository;

import com.stefanomantini.timestampconsumerservice.entity.TimestampEntity;
import java.util.UUID;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface TimestampRepository extends JpaRepository<TimestampEntity, UUID> {}
