# Private Golang Playground for PoC's

## Resources

- [Golang Benchmarking](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)
- [Comparable in Golang](https://levelup.gitconnected.com/comparable-in-golang-d7b080e49443)

---

- [Official docs about sync.Map](https://pkg.go.dev/sync#Map)
- [How does sync.Map work internally](https://sreramk.medium.com/go-inside-sync-map-how-does-sync-map-work-internally-97e87b8e6bf)
- [An In-Depth Guide for Using Go sync.Map with Code Sample](https://reliasoftware.com/blog/go-sync-map)
- [Comparing Go performance of Mutex, sync.Map and Channels for concurrent map access for increments](https://medium.com/@nikolay.dimitrov/comparing-go-performance-of-mutex-sync-map-and-channels-for-syncing-a-map-of-counters-61810c201e17)

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
