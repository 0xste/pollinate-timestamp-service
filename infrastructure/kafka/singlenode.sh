docker run -d --network=host --env ADVERTISED_HOST=`docker-machine ip \`docker-machine active\`` --env ADVERTISED_PORT=9092 --env TOPICS=timestamp.command --name shared-kafka spotify/kafka
