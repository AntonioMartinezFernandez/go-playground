test:
  go test -race ./...

clean-cache:
  go clean -testcache

start-es:
  docker compose --file build/elasticsearch/compose.yml up -d

stop-es:
  docker compose --file build/elasticsearch/compose.yml down

start-kafka:
  docker compose --file build/kafka/compose.yml up -d

stop-kafka:
  docker compose --file build/kafka/compose.yml down