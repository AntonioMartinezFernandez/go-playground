# Private Golang Playground for PoC's

## Resources

- [An In-Depth Guide for Using Go sync.Map with Code Sample](https://reliasoftware.com/blog/go-sync-map)

---

- [ElasticSearch docker compose + curl requests examples](https://geshan.com.np/blog/2023/06/elasticsearch-docker/)
- [Elastic Stack -primeros pasos-](https://www.elastic.co/es/blog/getting-started-with-the-elastic-stack-and-docker-compose)
- [Full-text Search with ElasticSearch with Golang (EXAMPLE PROJECT)](https://github.com/code-heim/go_46_elasticsearch)

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
