# Private Golang Playground for PoC's

## Resources

-[ElasticSearch docker compose + curl requests examples](https://geshan.com.np/blog/2023/06/elasticsearch-docker/)

## PoC's

### Worker Pool

### SyncMap

### ElasticSearch

Create index example:

```bash
curl -X PUT http://localhost:9200/test
```

Indexing example:

```bash
curl -X POST -H 'Content-Type: application/json' -d '{ "title": "Awesome Test!", "description": "This is an awesome example", "price": 3.50, "category": "Example Category", "brand": "Example Brand" }' http://localhost:9200/test/_doc
```

Search example:

```bash
curl -X GET "localhost:9200/test/_search?pretty" -H 'Content-Type: application/json' -d' { "query": { "match": { "title": "test" } } }'
```
