test:
  go test -race ./...

clean-cache:
  go clean -testcache

start-es:
  docker compose --file deployments/elasticsearch/compose.yml up -d

stop-es:
  docker compose --file deployments/elasticsearch/compose.yml down

start-kafka:
  docker compose --file deployments/kafka/compose.yml up -d

stop-kafka:
  docker compose --file deployments/kafka/compose.yml down