Install gomocks:
https://github.com/uber-go/mock

use ```go generate ./...``` to generate mocks

```
docker network create elastic

docker run --name es01 --net elastic -p 9200:9200 -it -m 1GB docker.elastic.co/elasticsearch/elasticsearch:8.16.1

docker run -p 9000:9000 lmenezes/cerebro
```
