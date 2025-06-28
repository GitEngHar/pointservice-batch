docker buildx build --platform linux/amd64 \
--build-arg JOB_PATH=./cmd/serverHealthCheck \
-t point-service/batch \
--load \
-f Dockerfile \
.

docker build \
  -t 363471485358.dkr.ecr.ap-northeast-1.amazonaws.com/point-service/batch:latest \
  -f Dockerfile \
  --build-arg JOB_PATH=./cmd/serverHealthCheck \
  .