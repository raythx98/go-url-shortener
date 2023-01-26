docker stop url-shortener-go
docker rm url-shortener-go
docker rmi url-shortener-go
docker build -t url-shortener-go .
docker run \
  --name url-shortener-go \
  --rm \
  -p 9281:9281 \
  -p 9282:9282 \
  -p 9283:9283 \
  -p 9284:9284 \
  url-shortener-go