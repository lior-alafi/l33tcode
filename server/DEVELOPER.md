Install gomocks:
https://github.com/uber-go/mock

use ```go generate ./...``` to generate mocks
https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html
```
docker network create elastic

docker run --name es01 --net elastic -p 9200:9200 -p 9300:9300  -e "discovery.type=single-node" -it -m 1GB docker.elastic.co/elasticsearch/elasticsearch:8.16.1
or 
docker run -d --name elasticsearch \
  -p 127.0.0.1:9200:9200 \
  -p 127.0.0.1:9300:9300 \
  -e "discovery.type=single-node" \
  docker.elastic.co/elasticsearch/elasticsearch:8.10.2


docker run -p 127.0.0.1:9200:9200 -p 127.0.0.1:9300:9300 -e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.17.25

docker run -p 9000:9000 lmenezes/cerebro
```
